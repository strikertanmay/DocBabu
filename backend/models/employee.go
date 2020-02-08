package models

import "gopkg.in/mgo.v2/bson"

type Employee struct {
	ID            bson.ObjectId `bson:"_id" json:"id"`
	Name          string        `bson:"name" json:"name"`
	Designation   string        `bson:"designation" json:"designation"`
	Department    string        `bson:"department" json:"department"`
	OfficeAddress string        `bson:"office_address" json:"office_address"`
	OfficeTimings string        `bson:"office_timings" json:"office_timings"`
}
