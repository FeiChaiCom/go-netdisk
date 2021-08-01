package gormstore

import (
	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
	"log"
	"time"
)

type Store interface {
	sessions.Store
	PeriodicCleanup(time.Duration, <-chan struct{})
}

type store struct {
	*GormStore
}

func (c *store) Options(options sessions.Options) {
	c.sessionOptions = options.ToGorillaOptions()
}

func NewStore(db *gorm.DB, opts Options, keyPairs ...[]byte) (Store, error) {
	s := NewGormStore(db, opts, keyPairs...)
	return s, nil
}

// PeriodicCleanup runs Cleanup every interval. Close quit channel to stop.
func (c *store) PeriodicCleanup(interval time.Duration, quit <-chan struct{}) {
	t := time.NewTicker(interval)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			c.Cleanup()
		case <-quit:
			log.Printf("quit PeriodicCleanup task\n")
			return
		}
	}
}
