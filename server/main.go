package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/login", headers)
	http.HandleFunc("/register", headers)
	http.HandleFunc("/getEvents", hello)
	http.HandleFunc("/addEvent", hello)
	http.HandleFunc("/deleteEvent", hello)
	http.ListenAndServe(":8888", nil)
}
