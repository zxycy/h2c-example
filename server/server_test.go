package server

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"golang.org/x/net/http2"
	"log"
	"net"
	"net/http"
	"testing"
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

func Test_h2c_Client(t *testing.T) {

	client := h2cClient
	//GET
	resp, err := client.Get("http://127.0.0.1:8888/test")
	resp.Body.Close()
	log.Println(err, resp)

	//POST
	m := make(map[string]string)
	m["test"] = "test1"
	body, _ := json.Marshal(m)
	read := bytes.NewReader(body)
	resp, err = client.Post("http://127.0.0.1:8888/test/test1", "application/json", read)
	resp.Body.Close()
	log.Println(err, resp)

	//do
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8888/test", nil)
	resp, err = client.Do(req)
	resp.Body.Close()
	log.Println(err, resp)
}
func Test_h2_Client(t *testing.T) {

	client := httpsClient
	//GET
	resp, err := client.Get("https://127.0.0.1:8888/test")
	resp.Body.Close()
	log.Println(err, resp)

	//POST
	m := make(map[string]string)
	m["test"] = "test1"
	body, _ := json.Marshal(m)
	read := bytes.NewReader(body)
	resp, err = client.Post("https://127.0.0.1:8888/test/test1", "application/json", read)
	resp.Body.Close()
	log.Println(err, resp)

	//do
	req, _ := http.NewRequest("GET", "https://127.0.0.1:8888/test", nil)
	resp, err = client.Do(req)
	resp.Body.Close()
	log.Println(err, resp)
}
