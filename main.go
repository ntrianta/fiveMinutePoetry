package main

import (
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("poetry"))
	http.Handle("/", fs)

	http.ListenAndServe(":80", nil)

}
