package main

import (
	"fmt"
	"net/http"
)

func main() {
	servers := []Server{
		newSimpleServer("http://www.linkedin.com/in/rodrigobarenco"),
		newSimpleServer("https://web-portfolio-eight-cyan.vercel.app"),
		newSimpleServer("https://github.com/RodBarenco"),
	}
	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.serverProxy(w, r)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("Serving requests at localhost: %v \n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
