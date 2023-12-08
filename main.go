package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("Server started ... ")
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("k8s method recived = " + r.Method))
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Recived method %s, from host %s", r.Method, r.Host)
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("ListenAndServe log error: %s", err)
	}
}
