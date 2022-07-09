package routes

import (
	"encoding/json"
	"net/http"

	"github.com/LaurenteEber/gorm-rest-api/db"
	"github.com/LaurenteEber/gorm-rest-api/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// para obtener relacion de usuarios
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
	w.Write([]byte("get users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	// Para que todo lo que se ingrese por el POST se grave en la variable user
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdEuser := db.DB.Create(&user)
	err := createdEuser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
	// w.Write([]byte("post"))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete "))
}
