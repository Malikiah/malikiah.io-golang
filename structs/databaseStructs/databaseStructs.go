package databasestructs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Find struct {
	ID 						primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	MongoCollection 		string 				`json:"mongoCollection,omitempty" bson:"mongoCollection,omitempty"`
	Criteria 				string				`json:"criteria,omitempty" bson:"criteria,omitempty"`
	CriteriaValue 			string 				`json:"criteriaValue,omitempty" bson:"criteriaValue,omitempty"`
	Password  				string				`json:"password,omitempty" bson:"password,omitempty"`
	FindAll 				bool 				`json:"findAll,omitempty" bson:"findAll,omitempty"`
}

type User struct {
	ID			primitive.ObjectID 	`json:"_id,omitempty" bson:"_id,omitempty"`
	Email     	string				`json:"email,omitempty" bson:"email,omitempty"`
	UserName  	string				`json:"username,omitempty" bson:"username,omitempty"`
	Password  	string				`json:"password,omitempty" bson:"password,omitempty"`
	FirstName 	string				`json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  	string				`json:"lastname,omitempty" bson:"lastname,omitempty"`
}
