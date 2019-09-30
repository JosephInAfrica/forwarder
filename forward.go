package main

import (
	"bytes"
	"encoding/json"
	// "errors"
	"fmt"
	"forwarder/read"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	host string
	port string
}

type normalHandle struct {
	host string
	port string
}

func (this *normalHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse("http://" + this.host + ":" + this.port)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ModifyResponse = modifyResp
	proxy.ServeHTTP(w, r)
}

func modifyResp(resp *http.Response) error {
	// 疑似这里已经被 read了。
	origin := resp.Body

	res, err := ioutil.ReadAll(origin)
	defer origin.Close()
	if err != nil {
		return err
	}
	status := read.Status{}
	err = json.Unmarshal(res, &status)

	if err != nil {
		fmt.Println("unmarshal failed")
		fmt.Println("res", res)
		return nil
	}

	fmt.Println("res", res)
	// fmt.Printf("%s %v", "status", status)

	s, err := json.Marshal(&status)
	if err != nil {
		return err
	}

	fmt.Println("marshalld ", string(s))
	// fmt.Println("")
	// in case content length does not match .There would be fatal error.
	resp.Header["Content-Length"] = []string{}
	// 找了半天就在找这个。 bytes.NewReader(s)将[]byte转成Reader. NopCloser将Reader转成ReadCloser.
	resp.Body = ioutil.NopCloser(bytes.NewReader(s))
	//
	fmt.Println("response", resp)
	// fmt.Println("resp.body", resp.Body)
	// defer resp.Body.Close()
	// defer origin.Close()
	return nil

}

func startServer() {
	//被代理的服务器host和port

	h_temp := &handle{host: "192.168.0.75", port: "80"}
	h_normal := &handle{host: "192.168.0.75", port: "80"}
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}

func main() {
	startServer()
}
