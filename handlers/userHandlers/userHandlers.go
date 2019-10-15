package userhandlers

import (
	"log"
	"net/http"
	//"fmt"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//Services
	databaseservice "malikiah.io/services/databaseService"
	passwordservice "malikiah.io/services/passwordService"
	userservice "malikiah.io/services/userService"

	//Structs
	//databasestructs "malikiah.io/structs/databasestructs"
	databasestructs "malikiah.io/structs/databaseStructs"
)

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	objectID, _ := primitive.ObjectIDFromHex(request.FormValue("_id"))
	databaseQuery := databasestructs.Find{
		ID: objectID,
		MongoCollection: "user",
		Criteria: "email",
		CriteriaValue: request.FormValue("email"),
		FindAll: false,
	}

	userservice.Login(databaseQuery, request)
}

func RegistrationHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	if request.FormValue("password") != request.FormValue("confirmpassword") {
		log.Println("Passwords do not match")
	} else {

		password, _ := passwordservice.HashPassword(request.FormValue("password"))

		body := databasestructs.User{
			Email:     request.FormValue("email"),
			UserName: request.FormValue("username"),
			Password:  password,
			FirstName: request.FormValue("firstname"),
			LastName:  request.FormValue("lastname"),
		}
		
		log.Println(body)

		databaseservice.Insert("user", body)
	}

}
