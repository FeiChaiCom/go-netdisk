package httpclient

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"go-netdisk/pkg/utils/httpclient/ssl"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// HttpRespone define the information of the http respone
type HttpRespone struct {
	Reply      []byte
	StatusCode int
	Status     string
	Header     http.Header
}

type HttpClient struct {
	caFile   string
	certFile string
	keyFile  string
	header   map[string]string
	httpCli  *http.Client
}

type HeaderSet struct {
	Key   string
	Value string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		httpCli: &http.Client{},
		header:  make(map[string]string),
	}
}

func (client *HttpClient) GetClient() *http.Client {
	return client.httpCli
}

func (client *HttpClient) SetTlsNoVerity() error {
	tlsConf := ssl.ClientTslConfNoVerity()

	trans := client.NewTransPort()
	trans.TLSClientConfig = tlsConf
	client.httpCli.Transport = trans

	return nil
}

func (client *HttpClient) SetTlsVerityServer(caFile string) error {
	client.caFile = caFile

	// load ca cert
	tlsConf, err := ssl.ClientTslConfVerityServer(caFile)
	if err != nil {
		return err
	}

	client.SetTlsVerityConfig(tlsConf)

	return nil
}

func (client *HttpClient) SetTlsVerity(caFile, certFile, keyFile, passwd string) error {
	client.caFile = caFile
	client.certFile = certFile
	client.keyFile = keyFile

	// load cert
	tlsConf, err := ssl.ClientTslConfVerity(caFile, certFile, keyFile, passwd)
	if err != nil {
		return err
	}

	client.SetTlsVerityConfig(tlsConf)

	return nil
}

func (client *HttpClient) SetTlsVerityConfig(tlsConf *tls.Config) {
	trans := client.NewTransPort()
	trans.TLSClientConfig = tlsConf
	client.httpCli.Transport = trans
}

func (client *HttpClient) NewTransPort() *http.Transport {
	return &http.Transport{
		TLSHandshakeTimeout: 5 * time.Second,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		ResponseHeaderTimeout: 30 * time.Second,
	}
}

func (client *HttpClient) SetTimeOut(timeOut time.Duration) {
	client.httpCli.Timeout = timeOut
}

// SetHeader set header for the http client。
// Note：if the header is the same with the parameter(header) which is specified
// in the function GET, POST, PUT,DELETE,Patch and so on. this set header is ignore in the call
func (client *HttpClient) SetHeader(key, value string) {
	client.header[key] = value
}

// SetBatchHeader batch set header for the http client。
// Note：if the header is the same with the parameter(header) which is specified
// in the function GET, POST, PUT,DELETE,Patch and so on. this set header is ignore in the call
func (client *HttpClient) SetBatchHeader(headerSet []*HeaderSet) {
	if headerSet == nil {
		return
	}
	for _, header := range headerSet {
		client.header[header.Key] = header.Value
	}
}

func (client *HttpClient) GET(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "GET", header, data)
}

func (client *HttpClient) POST(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "POST", header, data)
}

func (client *HttpClient) DELETE(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "DELETE", header, data)
}

func (client *HttpClient) PUT(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "PUT", header, data)
}

func (client *HttpClient) PATCH(url string, header http.Header, data []byte) ([]byte, error) {
	return client.Request(url, "PATCH", header, data)
}

func (client *HttpClient) Get(url string, header http.Header, data []byte) (*HttpRespone, error) {
	return client.RequestEx(url, "GET", header, data)
}

func (client *HttpClient) Post(url string, header http.Header, data []byte) (*HttpRespone, error) {
	return client.RequestEx(url, "POST", header, data)
}

func (client *HttpClient) Delete(url string, header http.Header, data []byte) (*HttpRespone, error) {
	return client.RequestEx(url, "DELETE", header, data)
}

func (client *HttpClient) Put(url string, header http.Header, data []byte) (*HttpRespone, error) {
	return client.RequestEx(url, "PUT", header, data)
}

func (client *HttpClient) Patch(url string, header http.Header, data []byte) (*HttpRespone, error) {
	return client.RequestEx(url, "PATCH", header, data)
}

func (client *HttpClient) Request(url, method string, header http.Header, data []byte) ([]byte, error) {
	rsp, err := client.RequestEx(url, method, header, data)
	return rsp.Reply, err
}

func (client *HttpClient) RequestEx(url, method string, header http.Header, data []byte) (*HttpRespone, error) {
	var req *http.Request
	var errReq error
	httpRsp := &HttpRespone{
		Reply:      nil,
		StatusCode: http.StatusInternalServerError,
		Status:     "Internal Server Error",
	}

	if data != nil {
		req, errReq = http.NewRequest(method, url, bytes.NewReader(data))
	} else {
		req, errReq = http.NewRequest(method, url, nil)
	}

	if errReq != nil {
		return httpRsp, errReq
	}

	req.Close = true

	if header != nil {
		req.Header = header
	}

	for key, value := range client.header {
		if req.Header.Get(key) != "" {
			continue
		}
		req.Header.Set(key, value)
	}

	rsp, err := client.httpCli.Do(req)
	if err != nil {
		return httpRsp, err
	}

	defer rsp.Body.Close()

	httpRsp.Status = rsp.Status
	httpRsp.StatusCode = rsp.StatusCode
	httpRsp.Header = rsp.Header

	rpy, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return httpRsp, err
	}

	httpRsp.Reply = rpy
	return httpRsp, nil
}

func Request(url, method string, header http.Header, body io.Reader) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		errMsg := fmt.Sprintf("fail to new a http request. err:%s", err.Error())
		return errMsg, errors.New(errMsg)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	req.Close = true

	// header
	if header != nil {
		req.Header = header
	}

	client := &http.Client{}

	rsp, err := client.Do(req)
	if err != nil {
		return "", errors.New(fmt.Sprintf("fail to do http request. err:%s", err.Error()))
	}

	defer rsp.Body.Close()

	replyData, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", errors.New(fmt.Sprintf("read respone failed. err:%s", err.Error()))
	}

	return string(replyData), nil
}
