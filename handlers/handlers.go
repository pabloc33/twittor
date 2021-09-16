package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	middlew "github.com/pabloc33/twittor/middleW"
	"github.com/pabloc33/twittor/routers"

	"github.com/rs/cors"
)

func Manejadores(){
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	// router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.Login)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}