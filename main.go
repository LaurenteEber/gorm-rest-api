package main

import (
	"net/http"

	"github.com/LaurenteEber/gorm-rest-api/db"
	"github.com/LaurenteEber/gorm-rest-api/models"
	"github.com/LaurenteEber/gorm-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConecction()
	// Para crear los modelos de tablas de tarea y usuario
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{i}", routes.GetUserHandler).Methods("GET") //para acceder a algun usuario en particular {i}
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandler).Methods("DELETE")

	// Para inicializar el server
	http.ListenAndServe(":3000", r)
}
