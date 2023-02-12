package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var nextServerIndex int32 = 0

func main() {
	var mu sync.Mutex

	originServerList := []string{
		"http://localhost:8081",
		"http://localhost:8082",
	}

	loadBalancerHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		mu.Lock()

		originServerURL, _ := url.Parse(originServerList[(nextServerIndex)%2])

		nextServerIndex++
		mu.Unlock()

		reverseProxy := httputil.NewSingleHostReverseProxy(originServerURL)

		reverseProxy.ServeHTTP(w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", loadBalancerHandler))
}
