package main

import (
	"log"
	"net/http"
	"time"

	"github.com/eduardtomasek/golang-notebook/gorillapi/models"
	"github.com/eduardtomasek/golang-notebook/gorillapi/routes"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	err := models.Init()

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.Use(headersMiddleware)

	router.HandleFunc("/", GetRoot).Methods("GET")
	routes.MakeRouter(router)

	srv := &http.Server{
		Handler: handlers.CORS()(router),
		// Handler:      router,
		Addr:         "127.0.0.1:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// GetRoot handling base request
func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{ "status": "OK" }`))
}
