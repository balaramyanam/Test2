package routes

import (
	"log"
	"inter/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Register() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/smiles/{name}", controllers.Api).Methods("post")

	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
