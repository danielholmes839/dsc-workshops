package main

/*
go mod init example.com
go get -u github.com/go-chi/chi/v5
*/
import (
	"net/http"

	"example.com/greet"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		message := []byte(greet.Greet())
		w.Write(message)
	})
	
	http.ListenAndServe(":3000", r)
}
