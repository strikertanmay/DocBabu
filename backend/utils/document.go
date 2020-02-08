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

func (m *Server) FindDocumentByID(id string) (Document, error) {
	var document Document
	err := db.C(DOCUMENT_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&document)

	if err != nil {
		log.Fatal(err)
	}

	return document, err
}

func (m *Server) InsertDocument(doc Document) error {
	err := db.C(DOCUMENT_COLLECTION).Insert(&doc)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (m *Server) UpdateDocument(doc Document) error {
	err := db.C(DOCUMENT_COLLECTION).UpdateId(doc.ID, &doc)

	if err != nil {
		log.Fatal(err)
	}

	return err
}
