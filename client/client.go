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

func Client()  {

	client := http.Client{
		Transport: &http2.Transport{
			// So http2.Transport doesn't complain the URL scheme isn't 'https'
			AllowHTTP: true,
			// Pretend we are dialing a TLS endpoint. (Note, we ignore the passed tls.Config)
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	//GET
	resp, err := client.Get("http://127.0.0.1:8888/test")
	log.Println(err, resp)
	//POST
	m:=make(map[string]string)
	m["test"]="test1"
	body,_:=json.Marshal(m)
	read:=bytes.NewReader(body)
	resp, err = client.Post("http://127.0.0.1:8888/test/test1","application/json", read)

	log.Println(err, resp)
}

