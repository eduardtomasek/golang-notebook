package routes

import (
	"github.com/gorilla/mux"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// MakeRouter create user router
func MakeRouter(r *mux.Router) {
	// opts := &respond.Options{}

	userRouter := r.PathPrefix("/user").Subrouter()

	// userRouter.Path("/{id}").HandlerFunc(getPerson).Methods("GET")
	userRouter.Path("/list").HandlerFunc(listPersonsHandler).Methods("GET")
}
