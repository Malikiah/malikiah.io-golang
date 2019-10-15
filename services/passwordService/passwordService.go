package passwordservice

import (
	"time"


	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
)

//var key = os.Get("SIGNING_KEY")
var key = []byte("Secret")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateJWT( id primitive.ObjectID ) (JSONWebToken string, err error) {
	token := jwt.New(jwt.SigningMethodHS512)

	claims := token.Claims.(jwt.MapClaims)

	claims["privilege"] = "user"
	claims["_id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 180).Unix()
	
	JSONWebToken, err = token.SignedString(key)

	return
}

func CheckJWT(tokenString string) (decodedJWT jwt.Keyfunc, err error){
	
	jwt.Parse(tokenString, decodedJWT)
	
	return
}