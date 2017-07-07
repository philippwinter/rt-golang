package rt

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"
)

type HttpClient struct {
	http.Client
}

func Init(t *testing.T, config *tls.Config) (*Tester, *HttpClient) {
	return InitT(t), InitClient(config)
}

func InitClient(config *tls.Config) *HttpClient {
	return &HttpClient{
		http.Client{
			Transport: &http.Transport{
				TLSClientConfig: config,
			},
			Timeout: time.Second * 10,
		},
	}
}
func InitT(t *testing.T) *Tester {
	return &Tester{
		*t,
	}
}

func StandardTLSConfig() *tls.Config {
	return &tls.Config{}
}

func InsecureTLSConfig() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}
