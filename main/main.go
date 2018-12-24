package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

func (s server) routes() {
	s.router.HandleFunc("/asd", s.handleIndex())
	//s.router.HandleFunc("/admin", s.adminOnly(s.handleAdminIndex()))
}
func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "secret stuff")
	}
}

func main() {
	s := server{
		router: http.NewServeMux(),
	}
	s.routes()
	log.Fatal(http.ListenAndServe(":9000", s.router))
}
