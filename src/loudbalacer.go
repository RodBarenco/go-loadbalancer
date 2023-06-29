package main

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}
