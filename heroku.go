package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func herokuHack() {
	go func() {
		http.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(rw, "pong")
		})
		err = http.ListenAndServe(net.JoinHostPort("", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	for range time.Tick(time.Minute * 10) {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		if body, err := io.ReadAll(resp.Body); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(body)
		}
	}

}
