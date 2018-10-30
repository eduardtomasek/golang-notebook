package routes

import "github.com/gorilla/mux"

// MakeRouter create user router
func MakeRouter(r *mux.Router) {
	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.Path("/{id}").HandlerFunc(getUser).Methods("GET")
	userRouter.Path("/list").HandlerFunc(listUsers).Methods("GET")
}
