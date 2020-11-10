package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("https://api.stackexchange.com/")
	log.Printf("forwarding to -> %s://%s\n", target.Scheme, target.Host)

	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		log.Println("req.Host=", req.Host)
		log.Println("req.URL.Host=", req.URL.Host)
		req.Host = req.URL.Host

		proxy.ServeHTTP(w, req)
	})

	err = http.ListenAndServe(":8989", nil)
	if err != nil {
		panic(err)
	}
}
