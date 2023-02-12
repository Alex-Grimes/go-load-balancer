package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	portFlag := flag.Int("port", 8081, "listening port")
	flag.Parse()
	port := fmt.Sprintf(":%d", *portFlag)

	originServerHandler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Printf("[origin server] recived request: %s\n", port)
	})
	log.Fatal(http.ListenAndServe(port, originServerHandler))
}
