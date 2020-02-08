package utils

import (
	. "github.com/jigar3/docBabu/models"
	"gopkg.in/mgo.v2/bson"
)

var (
	EMPLOYEE_COLLECTION = "employees"
)

func FindAllEmployees() ([]Employee, error) {
	var employees []Employee
	err := db.C(EMPLOYEE_COLLECTION).Find(bson.M{}).All(&employees)

	if err != nil {
		return []Employee{}, err
	}

	return employees, err

}

func FindEmployeeByID(id string) (Employee, error) {
	var employee Employee
	err := db.C(EMPLOYEE_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&employee)

	if err != nil {
		return Employee{}, err
	}

	return employee, err
}

func FindEmployeeByName(name string) (Employee, error) {
	var employee Employee
	err := db.C(EMPLOYEE_COLLECTION).Find(bson.M{"name": name}).One(&employee)

	if err != nil {
		return Employee{}, err
	}

	return employee, err
}

func InsertEmployee(employee Employee) error {
	err := db.C(EMPLOYEE_COLLECTION).Insert(&employee)

	if err != nil {
		return err
	}

	return err
}

func DeleteEmployee(employee Employee) error {
	err := db.C(EMPLOYEE_COLLECTION).Remove(&employee)

	if err != nil {
		return err
	}

	return err
}

func UpdateEmployee(employee Employee) error {
	err := db.C(EMPLOYEE_COLLECTION).UpdateId(employee.ID, &employee)

	if err != nil {
		return err
	}

	return err
}
