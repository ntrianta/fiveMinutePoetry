package main

import (
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("$GOPATH/src/github.com/ntrianta/fiveMinutePoetry/poetry/"))
	http.Handle("/", fs)

	http.ListenAndServe(":80", nil)

}
