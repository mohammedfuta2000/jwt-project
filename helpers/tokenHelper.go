package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mohammedfuta2000/jwt-project/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// "golang-jwt-project/database"

//
// "github.com/golang-jwt/jwt"

// jwt token uses a hashing mechanism to convert all the fields u feed it, rg: email, phone, first_name etc
// into a single string
// it also consists of a secret key

type SignedDetails struct{
	Email string
	First_name string
	Last_name string
	Uid string
	User_type string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client,"users")

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, first_name string, last_name string, user_type string, uid string) (token, refresh_token string, err error) {
	claims := &SignedDetails{
		Email: email,
		First_name: first_name,
		Last_name: last_name,
		Uid: uid,
		User_type: user_type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(24)).Unix(),
		},
	}

	refresh_claims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour*time.Duration(168)).Unix(),
		},
	}

	token, err =jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY) )
	refresh_token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims).SignedString([]byte(SECRET_KEY))

	if err!=nil{
		log.Panic(err)
		return
	}
	return token,refresh_token,err
}

func ValidateToken(singedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		singedToken,
		&SignedDetails{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err!=nil {
		msg = err.Error()
		return
	}
	claims, ok:= token.Claims.(*SignedDetails)
	if !ok{
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("toekn is expired")
		msg = err.Error()
		return
	}
	return claims, msg
	
}

func UpdateAllTokens(signedToken, signedRefreshToken, userId string)  {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	updated_at, _:= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at",Value: updated_at})

	upsert:=true
	filter:= bson.M{"user_id":userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err :=userCollection.UpdateOne(
		ctx,
		filter,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)
	if err!=nil {
		log.Panic(err)
		return 
	}
	return

}