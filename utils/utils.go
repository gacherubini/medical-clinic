package utils

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"medical-clinic/models"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func IsAdminAllowed(r *http.Request, db *sql.DB) bool {
	UserIdFromToken, err := UserIdFromToken(r)
	fmt.Println(UserIdFromToken)
	if err != nil {
		return false
	}

	if isUserAllowed(r, db) {
		return false
	}

	UserFromToken, err := models.FindUser(context.Background(), db, UserIdFromToken)
	if err != nil {
		return false
	}

	fmt.Println(UserFromToken.Role)
	if UserFromToken.Role != "admin" {
		return false
	}

	return true
}

func isUserAllowed(r *http.Request, db *sql.DB) bool {
	params := mux.Vars(r)
	id := params["id"]
	intID, err := strconv.Atoi(id)

	if err != nil {
		return false
	}

	UserIdFromToken, err := UserIdFromToken(r)
	if err != nil {
		return false
	}

	UserFromToken, err := models.FindUser(context.Background(), db, UserIdFromToken)

	switch UserFromToken.Role {
	case "doctor":
		userDoctor, err := models.FindDoctor(context.Background(), db, UserFromToken.UserID)
		if err != nil {
			return false
		}

		doctor, err := models.FindDoctor(context.Background(), db, intID)
		if err != nil {
			return false
		}

		if doctor.DoctorID == userDoctor.DoctorID {
			return true
		}
	case "patient":
		userPatient, err := models.FindPatient(context.Background(), db, UserFromToken.UserID)
		if err != nil {
			return false
		}

		patient, err := models.FindPatient(context.Background(), db, intID)
		if err != nil {
			return false
		}

		if patient.PatientID == userPatient.PatientID {
			return true
		}
	}
	return false
}

func GenereteToken(userId int) (string, error) {
	token_lifespan := 30
	fmt.Println(userId)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenToReturn, err := token.SignedString([]byte(os.Getenv("API_SECRET")))

	return tokenToReturn, err
}

func extractTokenFromHeader(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(os.Getenv("API_SECRET")), nil
}

func UserIdFromToken(r *http.Request) (int, error) {
	tokenString := extractTokenFromHeader(r)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return int(userId), nil
	}
	return 0, errors.New("invalid token")
}
