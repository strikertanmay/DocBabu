package routes

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/jigar3/docBabu/models"
	. "github.com/jigar3/docBabu/utils"
	"gopkg.in/mgo.v2/bson"
)

type CreateRequest struct {
	FileName     string        `bson:"filename" json:"filename"`
	CreatorID    bson.ObjectId `bson:"creator_id" json:"creator_id"`
	Associations []Association `bson:"associations" json:"associations"`
	Remarks      string        `bson:"remarks" json:"remarks"`
}
type Association struct {
	Name     string `bson:"name" json:"name"`
	Priority int    `bson:"priority" json:"priority"`
	Remarks  string `bson:"remarks" json:"remarks"`
}

// Name of File, ID of Person Creating, remarks, {name: "", priority: "", remarks: ""} List
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request_body CreateRequest
	json.NewDecoder(r.Body).Decode(&request_body)
	created_by, _ := FindEmployeeByID(request_body.CreatorID.Hex())

	var personAssociated []PersonDetail

	for i := 0; i < len(request_body.Associations); i++ {
		emp, _ := FindEmployeeByName(request_body.Associations[i].Name)
		temp := PersonDetail{
			Person:           emp,
			SignaturePending: true,
			SignedAt:         time.Now(),
			Remarks:          request_body.Associations[i].Remarks,
			Priority:         request_body.Associations[i].Priority,
		}

		personAssociated = append(personAssociated, temp)
	}

	document_id := bson.NewObjectId()
	var document = Document{
		ID:               document_id,
		CreatedBy:        created_by,
		CreatedAt:        time.Now(),
		Associations:     personAssociated,
		FinalDestination: created_by.Name,
		Priority:         0,
		NextEmployees:    []Employee{},
		Remarks:          request_body.Remarks,
		Completed:        false,
		Error:            false,
	}

	err := InsertDocument(document)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	name := CreateQR(document_id.Hex())
	http.ServeFile(w, r, name)
}

func GetAllDocuments(w http.ResponseWriter, r *http.Request) {
	documents, err := FindAllDocuments()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, documents)
}
