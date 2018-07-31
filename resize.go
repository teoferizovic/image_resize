package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"willnorris.com/go/imageproxy"
)

func main() {
	p := imageproxy.NewProxy(nil, nil)

	router := mux.NewRouter()
	router.NewRoute().Name("proxy").Methods("GET").PathPrefix("/proxy/").Handler(http.StripPrefix("/proxy", p))

	server := &http.Server{
		Addr:    ":8099",
		Handler: router,
	}

	//server.Handler = p

	log.Printf("Listening at %v", server.Addr)
	log.Fatal(server.ListenAndServe())
}

//http://127.0.0.1:8099/proxy/50x,/https:/octodex.github.com/images/codercat.jpg