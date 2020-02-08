package utils

import (
	"log"

	. "github.com/jigar3/docBabu/models"
	"gopkg.in/mgo.v2/bson"
)

var (
	DOCUMENT_COLLECTION = "documents"
)

func (m *Server) FindAllDocuments() ([]Document, error) {
	var documents []Document
	err := db.C(DOCUMENT_COLLECTION).Find(bson.M{}).All(&documents)

	if err != nil {
		log.Fatal(err)
	}

	return documents, err
}
