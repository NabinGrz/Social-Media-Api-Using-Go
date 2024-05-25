package authServices

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	userModel "github.com/NabinGrz/SocialMediaApi/src/authentication/models"
	"github.com/NabinGrz/SocialMediaApi/src/dbConnection"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var jwtkey = []byte("N8Sns89nS2ISB09sn290bSkSHJJ2SNoiS09")

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func GenerateJWT(user string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Register(user userModel.User) (*mongo.InsertOneResult, error) {
	var foundUser userModel.User
	filter := bson.M{"email": user.Email}

	result := dbConnection.SocialMediaCollection.FindOne(context.Background(), filter)

	if err := result.Decode(&foundUser); err != nil {
		return nil, err
	}

	if foundUser.Email != "" {
		return nil, errors.New("Email has alread been registered")
	}
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	createUser, err := dbConnection.SocialMediaCollection.InsertOne(context.Background(), user)
	return createUser, err
}

func Login(email string, password string) (map[string]interface{}, error) {
	var foundUser userModel.User

	err := dbConnection.SocialMediaCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&foundUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("User not found")
		}
		return nil, err
	}

	match := VerifyPassword(password, foundUser.Password)

	if match {
		token, _ := GenerateJWT(foundUser.Email)
		return map[string]interface{}{"token": token}, nil
	}
	return map[string]interface{}{"error": "Invalid Credentials"}, nil
}