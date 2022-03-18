package main

import (
	"fmt"
	"net/http"
)

/* 
Article:
https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html
*/

type Server struct {
	name string
}

func (s *Server) wordHandler(word string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200) // set the status code
		fmt.Fprintf(w, "server: %s, word: %s", s.name, word)
	}
}

func main() {
	server := &Server{name: "my server"}
	
	http.HandleFunc("/hello", server.wordHandler("hello"))
	http.HandleFunc("/bye", server.wordHandler("bye"))

	http.ListenAndServe("localhost:3000", nil)
}
