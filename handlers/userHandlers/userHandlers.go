package userhandlers

import (
	"log"
	"net/http"

	//Services
	databaseservice "malikiah.io/services/databaseService"
)

type User struct {
	email string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	body := User{
		email: r.FormValue("email"),
	}
	//body, err := ioutil.ReadAll(r.Body)
	/*if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}*/

	//r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	//byteArray, err := json.Marshal(r.Form)

	/*if err != nil {
		log.Println(err)
	}*/

	//jd := json.NewDecoder(r.GetBody)
	log.Println(r.FormValue("email"))
	log.Println(body.email)
	databaseservice.Insert("user", body)
	//password := "secret"
	//hash, _ := passwordService.HashPassword(password) // ignore error for the sake of simplicity
	//log.Println(hash)

}
