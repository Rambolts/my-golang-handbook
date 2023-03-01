package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public/"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)

	mux.HandleFunc("/func", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, this is function page.."))
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
