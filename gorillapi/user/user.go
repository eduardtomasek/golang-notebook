package user

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tomasek/golang-notebook/gorillapi/db"
	"github.com/tomasek/golang-notebook/gorillapi/resstruct"
)

// MakeRouter create user router
func MakeRouter(r *mux.Router) {
	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.Path("/{id}").HandlerFunc(getUser).Methods("GET")
	userRouter.Path("/list").HandlerFunc(listUsers).Methods("GET")
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	sqliteDB, err := db.Connect()

	if err != nil {
		json.NewEncoder(w).Encode(resstruct.ErrorResponse{Status: "ERROR", Message: err.Error()})
		return
	}

	rows, err := sqliteDB.Query(`SELECT id, login FROM person`)

	if err != nil {
		json.NewEncoder(w).Encode(resstruct.ErrorResponse{Status: "ERROR", Message: err.Error()})
		return
	}

	var users []resstruct.User

	for rows.Next() {
		user := resstruct.User{}

		rows.Scan(&user.ID, &user.Login)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(resstruct.UserResponse{Status: "OK", Data: users})
}

func getUser(w http.ResponseWriter, r *http.Request) {
}
