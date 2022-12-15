package main

import "net/http"

func (app *appilcation) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok!"))
	// TODO:
	// Json response!
}
