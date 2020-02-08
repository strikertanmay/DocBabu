package utils

import (
	. "github.com/jigar3/docBabu/models"
	"gopkg.in/mgo.v2/bson"
)

var (
	DOCUMENT_COLLECTION = "documents"
)

func FindAllDocuments() ([]Document, error) {
	var documents []Document
	err := db.C(DOCUMENT_COLLECTION).Find(bson.M{}).All(&documents)

	if err != nil {
		return []Document{}, err
	}

	return documents, err
}

func FindDocumentByID(id string) (Document, error) {
	var document Document
	err := db.C(DOCUMENT_COLLECTION).Find(bson.M{"_id": id}).One(&document)

	if err != nil {
		return Document{}, err
	}

	return document, err
}

func FindDocumentByName(name string) (Document, error) {
	var document Document
	err := db.C(DOCUMENT_COLLECTION).Find(bson.M{"filename": name}).One(&document)

	if err != nil {
		return Document{}, err
	}

	return document, err
}

func InsertDocument(doc Document) error {
	err := db.C(DOCUMENT_COLLECTION).Insert(&doc)

	if err != nil {
		return err
	}

	return err
}

func UpdateDocument(doc Document) error {
	err := db.C(DOCUMENT_COLLECTION).UpdateId(doc.ID, &doc)

	if err != nil {
		return err
	}

	return err
}
