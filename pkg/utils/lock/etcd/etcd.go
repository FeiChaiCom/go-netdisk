package locketcd

import (
	"context"
	"crypto/tls"
	"sync"
	"time"

	"go-netdisk/pkg/utils/lock"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

type EtcdLock struct {
	client  *clientv3.Client
	session *concurrency.Session

	lockMap sync.Map
}

func NewEtcdLock(discoveryUrls []string, tlsCfg *tls.Config) (lock.Interface, error) {
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: 5 * time.Second,
		Endpoints:   discoveryUrls,
		TLS:         tlsCfg,
	})
	if err != nil {
		return nil, err
	}

	// create concurrency session
	session, err := concurrency.NewSession(cli)
	if err != nil {
		return nil, err
	}
	return &EtcdLock{
		client:  cli,
		session: session,
		lockMap: sync.Map{},
	}, nil
}

// Lock lock a key with distribute-lock in etcd
func (r *EtcdLock) Lock(ctx context.Context, key string) error {
	distributeLock, _ := r.lockMap.LoadOrStore(key, concurrency.NewMutex(r.client, key))
	return distributeLock.(*concurrency.Mutex).Lock(ctx)
}

// UnLock unlock a key with distribute-lock in etcd
func (r *EtcdLock) UnLock(ctx context.Context, key string) error {
	distributeLock, ok := r.lockMap.Load(key)
	if !ok {
		return nil
	}
	return distributeLock.(*concurrency.Mutex).Unlock()
}
