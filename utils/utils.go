package utils

import (
	"context"
	"database/sql"
	"fmt"
	"medical-clinic/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

func IsAdminAllowed(r *http.Request, db *sql.DB) bool {
	adminIDString := r.Header.Get("ID")

	adminIDInt, err := strconv.Atoi(adminIDString)
	if err != nil {
		return true
	}

	if isUserAllowed(r, db) {
		return true
	}

	adminCheck, err := models.FindAdmin(context.Background(), db, adminIDInt)
	if err != nil {
		return false
	}

	userAdmin, err := models.FindUser(context.Background(), db, adminCheck.UserID)
	if err != nil {
		return false
	}

	if userAdmin.Role != "admin" {
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

	user, err := models.FindUser(context.Background(), db, intID)

	switch user.Role {
	case "doctor":
		doctor, _ := models.FindDoctor(context.Background(), db, user.UserID)
		if doctor.DoctorID == intID {
			return true
		}
	case "patient":
		patient, _ := models.FindPatient(context.Background(), db, user.UserID)
		if patient.PatientID == intID {
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
