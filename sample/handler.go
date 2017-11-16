package sample

import (
	"fmt"
	"net/http"
)

// SayHello HelloWorld出力
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, hello())
}

func hello() string {
	return "Hello"
}
