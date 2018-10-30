package routes

import (
	"encoding/json"
	"net/http"

	"github.com/tomasek/golang-notebook/gorillapi/models"
)

func listPersonsHandler(w http.ResponseWriter, r *http.Request) {
	persons, err := models.PersonList()

	if err != nil {
		w.Write([]byte(`{ "status": "ERROR", "message": "` + err.Error() + `" }`))
	}

	json.NewEncoder(w).Encode(Response{Status: "OK", Data: persons})
}

//func getPerson(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(`{ "status": "OK", "message": "Not implemented." }`))
//}
