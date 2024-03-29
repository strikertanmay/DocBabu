package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/jigar3/docBabu/models"
	. "github.com/jigar3/docBabu/utils"
	"gopkg.in/mgo.v2/bson"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var employee Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	employee.ID = bson.NewObjectId()
	if err := InsertEmployee(employee); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJson(w, http.StatusCreated, map[string]string{"result": "success"})
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	employee, err := FindAllEmployees()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, employee)
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	employee, err := FindEmployeeByName(params["name"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, employee)
}
