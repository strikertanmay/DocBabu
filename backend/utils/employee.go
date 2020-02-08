package utils

import (
	"log"

	. "github.com/jigar3/docBabu/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Server struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "employees"
)

func (m *Server) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

func (m *Server) FindAll() ([]Employee, error) {
	var employees []Employee
	err := db.C(COLLECTION).Find(bson.M{}).All(&employees)

	if err != nil {
		log.Fatal(err)
	}

	return employees, err

}
