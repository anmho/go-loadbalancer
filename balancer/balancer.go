package balancer

import (
	"net/http/httputil"
)

/* Enables registering application backeds. Utilizes a configurable load balancing algorithm */


type LoadBalancer struct {
	proxies map[string]*httputil.ReverseProxy
}

func NewLoadBalancer() *LoadBalancer {
	


	return &LoadBalancer{

	}

}

func (lb *LoadBalancer) RegisterBackend(host string, target string) {

}