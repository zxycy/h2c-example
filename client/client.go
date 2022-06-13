package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"golang.org/x/net/http2"
	"log"
	"net"
	"net/http"
)

var (
	httpsClient = &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	h2cClient = &http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
)

func Client(scheme string) {
	var client *http.Client
	if scheme == "http" {
		client = h2cClient
	} else if scheme == "https" {
		client = httpsClient
	} else {
		return
	}

	//GET
	resp, err := client.Get(scheme + "://127.0.0.1:8888/test")
	log.Println(err, resp)

	//POST
	m := make(map[string]string)
	m["test"] = "test1"
	body, _ := json.Marshal(m)
	read := bytes.NewReader(body)
	resp, err = client.Post(scheme+"://127.0.0.1:8888/test/test1", "application/json", read)
	log.Println(err, resp)

	//do
	req, _ := http.NewRequest("GET", scheme+"://127.0.0.1:8888/test", nil)
	resp, err = client.Do(req)
	log.Println(err, resp)

}
