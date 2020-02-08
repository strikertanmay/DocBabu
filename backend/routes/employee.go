package routes

import (
	"encoding/json"
	"net/http"

	. "github.com/jigar3/docBabu/models"
	"gopkg.in/mgo.v2/bson"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var employee Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		server.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	employee.ID = bson.NewObjectId()
	if err := server.InsertEmployee(employee); err != nil {
		server.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	server.RespondWithJson(w, http.StatusCreated, employee)
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employee, err := server.FindAllEmployees()
	if err != nil {
		server.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	server.RespondWithJson(w, http.StatusOK, employee)
}
