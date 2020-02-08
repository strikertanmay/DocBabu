package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/jigar3/docBabu/config"
	. "github.com/jigar3/docBabu/routes"
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
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employee/all", GetAllEmployees).Methods("GET")
	r.HandleFunc("/employee/{name}", GetEmployeeByName).Methods("GET")

	r.HandleFunc("/document", CreateDocument).Methods("POST")
	r.HandleFunc("/document/all", GetAllDocuments).Methods("GET")
	r.HandleFunc("/document/{name}", GetDocumentByName).Methods("GET")
	r.HandleFunc("/document/{id}", EditDocument).Methods("GET")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
