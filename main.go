package main

import (
	"log"
	"github.com/pabloc33/twittor/handlers"
	"github.com/pabloc33/twittor/bd"
	
)

func main(){

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
   
	handlers.Manejadores()
}