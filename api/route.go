package api

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/handlers"
	"btelli-customersupport-app/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {

	utils.LoadEnvFile(".env")

	database.Connect()

	database.SeedData()

	router := mux.NewRouter()

	// Define your route here...

	router.Handle("/", http.HandlerFunc(handlers.Home)).Methods("GET")

	log.Println("Customer Support App API running on Port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
