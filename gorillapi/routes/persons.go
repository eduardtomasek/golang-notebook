package routes

import (
	"net/http"
)

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{ "status": "OK", "message": "Not implemented." }`))
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{ "status": "OK", "message": "Not implemented." }`))
}
