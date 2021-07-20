package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	who string
	hostname string
)

func indexHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, fmt.Sprintf("Hello %s from %s!", who, hostname))
}

func healthHandler(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Don't worry. Puffy is fine!\n")
	fmt.Println("A health check was performed")
}

func main() {
	/*
	var present bool
	who, present = os.LookupEnv("FOO")
	if (!present) {
		who = "foobar"
	}
	*/
	flag.StringVar(&who, "foo", "foobar", "hello to who!?")
	flag.Parse()

	var err error

	// get the hostname
	hostname, err = os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/healthz", healthHandler)
	http.ListenAndServe(":8080", nil)
}
