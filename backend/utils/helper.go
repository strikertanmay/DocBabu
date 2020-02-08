package utils

import (
	"encoding/json"
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

type Server struct {
	Server   string
	Database string
}

var server = Server{}

var db *mgo.Database

func (m *Server) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

func (m *Server) RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (m *Server) RespondWithError(w http.ResponseWriter, code int, msg string) {
	server.RespondWithJson(w, code, map[string]string{"error": msg})
}
