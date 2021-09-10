package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func herokuHack() {
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("", port), nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	for range time.Tick(time.Minute * 15) {
		_, err = http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
	}

}
