package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/jigar3/docBabu/config"
	. "github.com/jigar3/docBabu/utils"
)

var config = Config{}
var server = Server{}

func init() {
	config.Read()

	server.Server = config.Server
	server.Database = config.Database
	server.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", abc).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func abc(w http.ResponseWriter, r *http.Request) {
	// server.
	// server.RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
