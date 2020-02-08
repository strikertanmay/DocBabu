package routes

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
	. "github.com/jigar3/docBabu/models"
	. "github.com/jigar3/docBabu/utils"
	"github.com/thoas/go-funk"
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

	sort.Slice(personAssociated, func(i, j int) bool {
		return personAssociated[i].Priority < personAssociated[j].Priority
	})

	result := funk.Filter(personAssociated, func(p PersonDetail) bool {
		return p.SignaturePending == true
	})

	x := result.([]PersonDetail)

	min_priority := x[0].Priority

	result = funk.Filter(result, func(x PersonDetail) bool {
		return x.Priority == min_priority
	})

	y := result.([]PersonDetail)

	document_id := bson.NewObjectId()
	var document = Document{
		ID:               document_id,
		FileName:         request_body.FileName,
		CreatedBy:        created_by,
		CreatedAt:        time.Now(),
		Associations:     personAssociated,
		FinalDestination: created_by.Name,
		Priority:         0,
		NextEmployees:    y,
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

func GetDocumentByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	documents, err := FindDocumentByName(params["name"])
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, documents)
}

func EditDocument(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
}
