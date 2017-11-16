package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ryutah/circle-ci-sample/sample"
)

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", sample.SayHello)

	http.Handle("/", r)
}
