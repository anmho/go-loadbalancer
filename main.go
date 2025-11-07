package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

/*
 * 1) Act as an API Gateway
 * 2) Act as a reverse proxy
 * Functionality
 * Backends should be able to register themselves with the backend
 * Backend should simply be able to register a list of ips mapped to the load balancers in a priority queue
 * Experiment with various load balancing algorithms
 */

func main() {
	proxies := map[string]*httputil.ReverseProxy{}

	hostTargets := map[string]string{
		"/shorts":                     "https://youtube.com/",
		"/internship":                 "https://hottake.gg/",
		"/us-edu/shop/back-to-school": "https://apple.com/",
	}

	for host, target := range hostTargets {
		//fmt.Print(target)
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
