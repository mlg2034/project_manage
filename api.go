package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	taskService := NewTaskService(s.store)
	taskService.RegisterRoutes(router)
	log.Println("starting API server at ", s.addr)

	log.Fatal(http.ListenAndServe(s.addr, subRouter))
}
