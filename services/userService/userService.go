package userservice

import (
	"log"
	//"time"
	"net/http"
	//"go/types"
	//"encoding/json"
	//Structs
	databasestructs "malikiah.io/structs/databaseStructs"
	
	//Services
	passwordservice "malikiah.io/services/passwordService"
	databaseservice "malikiah.io/services/databaseService"
)

func Login(databaseQuery databasestructs.Find, request *http.Request) (loggedIn bool, JSONWebToken string, err error) {
	user, _ := databaseservice.Find(databaseQuery)
	log.Println(user)
	loggedIn = passwordservice.CheckPasswordHash(request.FormValue("password"), user.Password)
	
	if loggedIn == true {
		
		JSONWebToken, err = passwordservice.CreateJWT(user.ID)

		log.Println(JSONWebToken)

	}

	return
}
