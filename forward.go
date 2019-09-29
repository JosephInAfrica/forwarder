package main

import (
	"forwarder/read"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	host string
	port string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ServeHTTP(w, r)
}

func modify(origin *http.Response) *http.Response {
	if res, err := Json.Unmarshal(origin.Body); err != nil {
		panic("parse response to Sensor info failed")
	}
}

func modifyResp(resp *http.Response) error {
	origin := resp.Body
	result := modify(origin)
	*origin = *result
}

func startServer() {
	//被代理的服务器host和port
	h := &handle{host: "192.168.0.75", port: "80"}
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
