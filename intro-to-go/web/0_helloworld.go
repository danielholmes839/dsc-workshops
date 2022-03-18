package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// send hello world!
		data := []byte("hello world!")
		w.Write(data)
	})

	http.ListenAndServe(":3000", nil)
}
