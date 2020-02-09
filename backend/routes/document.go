package routes

import (
	"encoding/json"
	"image"
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
	Name         string        `bson:"creator_name" json:"creator_name"`
	Associations []Association `bson:"associations" json:"associations"`
	Remarks      string        `bson:"remarks" json:"remarks"`
}
type Association struct {
	Name     string `bson:"name" json:"name"`
	Priority int    `bson:"priority" json:"priority"`
	Remarks  string `bson:"comment" json:"comment"`
}

func CreateDocument(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var request_body CreateRequest
	json.NewDecoder(r.Body).Decode(&request_body)
	created_by, _ := FindEmployeeByName(request_body.Name)

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
		CurrentlyWith:    created_by,
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
	r.ParseForm()
	file, _, _ := r.FormFile("image")
	img, _, _ := image.Decode(file)
	decoded_id := DecodeQR(img)

	document, _ := FindDocumentByID(decoded_id[0])
	current_user_id := r.FormValue("user_id")

	for i := 0; i < len(document.NextEmployees); i++ {
		if document.NextEmployees[i].Person.ID.Hex() == current_user_id && document.NextEmployees[i].SignaturePending == true {
			data := &document.NextEmployees[i]
			data.SignaturePending = false
			data.SignedAt = time.Now()
			data.Remarks = ""
		}
	}

	for i := 0; i < len(document.Associations); i++ {
		if document.Associations[i].Person.ID.Hex() == current_user_id && document.Associations[i].SignaturePending == true {
			data := &document.Associations[i]
			data.SignaturePending = false
			data.SignedAt = time.Now()
			data.Remarks = ""
		}
	}

	current_user, _ := FindEmployeeByID(current_user_id)
	document.CurrentlyWith = current_user

	result := funk.Filter(document.NextEmployees, func(p PersonDetail) bool {
		return p.Person.ID.Hex() != current_user_id
	})

	document.NextEmployees = result.([]PersonDetail)

	if len(document.NextEmployees) == 0 {
		sort.Slice(document.Associations, func(i, j int) bool {
			return document.Associations[i].Priority < document.Associations[j].Priority
		})

		result := funk.Filter(document.Associations, func(p PersonDetail) bool {
			return p.SignaturePending == true
		})

		x := result.([]PersonDetail)

		if len(x) > 0 {
			min_priority := x[0].Priority

			result = funk.Filter(result, func(x PersonDetail) bool {
				return x.Priority == min_priority
			})

			document.NextEmployees = result.([]PersonDetail)
		}

		if len(document.NextEmployees) == 0 {
			document.Completed = true
		}
	}

	UpdateDocument(document)

	RespondWithJson(w, http.StatusOK, document)
}

func GetDocumentsOfEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	documents, _ := FindAllDocuments()

	var related_documents []Document

	for _, data := range documents {
		if data.CreatedBy.Name == params["user_name"] {
			related_documents = append(related_documents, data)
		} else {
			for _, d := range data.Associations {
				if d.Person.Name == params["user_name"] {
					related_documents = append(related_documents, data)
				}
			}
		}
	}

	RespondWithJson(w, http.StatusOK, related_documents)
}
