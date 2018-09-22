package main

import (
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/help", Help)
	router.HandleFunc("/getAge",GetAge)

	return router
}

