package main

import (
"fmt"
"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe("0.0.0.0:80", nil)
}
