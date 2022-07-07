package main

import (
	"net/http"

	"github.com/LaurenteEber/gorm-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	// Para inicializar el server
	http.ListenAndServe(":3000", r)
}
