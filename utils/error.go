package utils

import (
	"net/http"
)

// NotFoundHandler : NotFoundHandler
var NotFoundHandler = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	Respond(w, Message(false, "This resources was not found on our server"))
}
