package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"medical-clinic/models"
	"medical-clinic/utils"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginHandlerContext struct {
	Db *sql.DB
}

func (contextHandler *LoginHandlerContext) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusMethodNotAllowed)
		return
	}

	var payload LoginPayload

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding JSON: ", err), http.StatusBadRequest)
		return
	}

	user, err := models.Users(qm.Where("email = ?", payload.Email)).One(context.Background(), contextHandler.Db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid credentials"), http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(payload.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid credential "), http.StatusBadRequest)
		return
	}

	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error genereting token: ", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "Error marshaling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Token:")
	w.Write(jsonResponse)

}
