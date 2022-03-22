package restyClient

import (
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

var Client *resty.Client

func InitRestyClient() {

	dialer := &net.Dialer{
		Timeout:   time.Duration(30 * time.Second),
		KeepAlive: time.Duration(30 * time.Second),
		DualStack: true,
	}

	transport := &http.Transport{
		DialContext:         dialer.DialContext,
		ForceAttemptHTTP2:   true,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     time.Duration(90 * time.Second),
		TLSHandshakeTimeout: time.Duration(100 * time.Second),
		MaxConnsPerHost:     100,
	}

	Client = resty.New()
	Client.SetTransport(transport)
}
