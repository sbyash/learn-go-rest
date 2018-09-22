package main

import (
    "log"
    "net/http"
)

func main(){

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 //        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
 //    })
    router := NewRouter()
	log.Fatal(http.ListenAndServe(":8090", router))
}