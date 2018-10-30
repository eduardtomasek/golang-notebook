package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tomasek/golang-notebook/gorillapi/models"
)

func listPersonsHandler(w http.ResponseWriter, r *http.Request) {
	persons, err := models.PersonList()

	if err != nil {
		w.Write([]byte(`{ "status": "ERROR", "message": "` + err.Error() + `" }`))
		return
	}

	json.NewEncoder(w).Encode(Response{Status: "OK", Data: persons})
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.Write([]byte(`{ "status": "ERROR", "message": "` + err.Error() + `" }`))
		return
	}

	person := models.Person{}
	person.Get(int64(id))

	json.NewEncoder(w).Encode(Response{Status: "OK", Data: []models.Person{person}})
}
