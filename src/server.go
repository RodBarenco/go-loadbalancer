package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	Adress() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}

func (s *simpleServer) Adress() string { return s.addr }

func (s *simpleServer) IsAlive() bool { return true }

func (s *simpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) serverProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to adress %v \n", targetServer.Adress())
	targetServer.Serve(w, r)
}
