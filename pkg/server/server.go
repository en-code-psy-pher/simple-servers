package server

import (
	"fmt"
	"net/http"
	"simpleserver/pkg/handlers"
)

func Run() {
    http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
	    handlers.EchoMatrix(w, r)
    })

    http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {
	    handlers.InvertMatrix(w, r)
    })

    http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {
	    handlers.FlattenMatrix(w, r)
    })

    http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {
	    handlers.SumMatrix(w, r)
    })

    http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
	    handlers.ProductMatrix(w, r)
    })

    fmt.Println("Server running at port 8080")
    http.ListenAndServe(":8080", nil)
}
