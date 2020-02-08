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

var (
	EMPLOYEE_COLLECTION = "employees"
)

func (m *Server) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

func (m *Server) FindAllEmployees() ([]Employee, error) {
	var employees []Employee
	err := db.C(EMPLOYEE_COLLECTION).Find(bson.M{}).All(&employees)

	if err != nil {
		log.Fatal(err)
	}

	return employees, err

}

func (m *Server) FindEmployeeById(id string) (Employee, error) {
	var employee Employee
	err := db.C(EMPLOYEE_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&employee)

	if err != nil {
		log.Fatal(err)
	}

	return employee, err
}

func (m *Server) Insert(employee Employee) error {
	err := db.C(EMPLOYEE_COLLECTION).Insert(&employee)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (m *Server) Delete(employee Employee) error {
	err := db.C(EMPLOYEE_COLLECTION).Remove(&employee)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (m *Server) Update(employee Employee) error {
	err := db.C(EMPLOYEE_COLLECTION).UpdateId(employee.ID, &employee)

	if err != nil {
		log.Fatal(err)
	}
	
	return err
}