package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

/*
 * 1) Act as an API Gateway
 * 2) Act as a reverse proxy
 * Functionality
 * Backends should be able to register themselves with the backend
 * Backend should simply be able to register a list of ips mapped to the load balancers in a priority queue
 * Experiment with various load balancing algorithms
 */



 func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)


	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.Host = url.Host
	}
	return proxy, nil
 }

 // input is a reverse proxy
 // returns an anononymous function that forwards the 
 func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
 }

func main() {
	fmt.Printf("Hello");

	proxies := map[string]*httputil.ReverseProxy{}

	hostTargets := map[string]string {
		"/shorts": "https://youtube.com/",
		"/internship": "https://hottake.gg/",
		"/us-edu/shop/back-to-school": "https://apple.com/",
	}

	for host, target := range hostTargets {
		fmt.Print(target)
		proxy, err := NewProxy(target)
		if err != nil {
			panic(err)
		}
		// proxies = append(proxies, proxy)
		proxies[host] = proxy
	}

	for host, proxy := range proxies {
		http.HandleFunc(host, ProxyRequestHandler(proxy))
	}


	log.Fatal(http.ListenAndServe(":9999", nil))
}