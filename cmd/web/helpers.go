package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// 500 internal server error
// change Println -> Output make more cleaner where is error occurred
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.erroLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// 400 -> bad request that user sent
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// clientError function warpeer for 404 not found
func (app *application) NotFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
