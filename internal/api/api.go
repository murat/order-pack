package api

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusFound)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}
