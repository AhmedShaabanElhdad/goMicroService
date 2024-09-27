package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	address string
	store   Store
}

func NewServer(addr string, store Store) *Server {
	return &Server{
		address: addr,
		store:   store,
	}
}

func (s *Server) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	// here we will add our service
	taskService := NewTaskSerive(s.store)
	taskService.RegisterRouter(router)

	log.Printf("Server Start at %+v \n", s.address)

	log.Fatal(http.ListenAndServe(s.address, subRouter))
}
