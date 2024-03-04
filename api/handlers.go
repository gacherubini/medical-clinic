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

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oi"))
}

func HandleCreateUniversity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var university models.University

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&university); err != nil {
		http.Error(w, "Error to Decode JSON", http.StatusBadRequest)
		return
	}

	err := university.Insert(context.Background(), db, boil.Infer())

	if err != nil {
		http.Error(w, "Error to Insert university", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Post successful")
}

func HandleGetAllUniversity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	universities, err := models.Universities().All(context.Background(), db)
	if err != nil && universities == nil {
		http.Error(w, "Failed to retrieve universities", http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(universities)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func HandleDeleteUniversity(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id := params["id"]
	intID, _ := strconv.Atoi(id)

	universities, err := models.FindUniversity(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve universitie", http.StatusInternalServerError)
		return
	}

	_, err = universities.Delete(context.Background(), db)
	if err != nil {
		http.Error(w, "Error deleting this universitie", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "University deleted successfully")
}

func HandleUpdateUniversity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
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

	university, err := models.FindUniversity(context.Background(), db, intID)
	if err != nil {
		http.Error(w, "Failed to retrieve university", http.StatusInternalServerError)
		return
	}

	_, err = university.Update(context.Background(), db, boil.Infer())
	if err != nil {
		http.Error(w, "Error updating university", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(university)
}
