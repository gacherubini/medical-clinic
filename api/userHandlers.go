package api

import (
	"context"
	"encoding/json"
	"fmt"
	"medical-clinic/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var user models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Error to Decode JSON", http.StatusBadRequest)
		return
	}

	err := user.Insert(context.Background(), db, boil.Infer())

	if err != nil {
		http.Error(w, "Error to Insert user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Post successful")
}

func HandleGetAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	user, err := models.Users().All(context.Background(), db)
	if err != nil && user == nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	intID, _ := strconv.Atoi(id)

	user, err := models.FindUser(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	_, err = user.Delete(context.Background(), db)
	if err != nil {
		http.Error(w, "Error deleting this user", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "User deleted successfully")
}

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	intID, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	user, err := models.FindUser(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	var userToUpdate = user

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userToUpdate); err != nil {
		http.Error(w, "Error to Decode JSON", http.StatusBadRequest)
		return
	}

	_, err = userToUpdate.Update(context.Background(), db, boil.Infer())
	if err != nil {
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "University updated successfully")
}
