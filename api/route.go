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

	// Comments
	router.Handle("/comments/{id}", http.HandlerFunc(handlers.GetComments)).Methods("GET")
	router.Handle("/createcomment", http.HandlerFunc(handlers.CreateComment)).Methods("POST")
	router.Handle("/updatecommment/{id}", http.HandlerFunc(handlers.UpdateComment)).Methods("PUT")
	router.Handle("/removecomment/{id}", http.HandlerFunc(handlers.RemoveComment)).Methods("DELETE")

	// Categories
	router.Handle("/categories", http.HandlerFunc(handlers.GetCategories)).Methods("GET")
	router.Handle("/category/{id}", http.HandlerFunc(handlers.GetCategory)).Methods("GET")
	router.Handle("/createcategory", http.HandlerFunc(handlers.CreateCategory)).Methods("POST")
	router.Handle("/updatecategory/{id}", http.HandlerFunc(handlers.UpdateCategory)).Methods("PUT")
	router.Handle("/removecategory/{id}", http.HandlerFunc(handlers.RemoveCategory)).Methods("DELETE")

	// Complaints
	router.Handle("/complaints", http.HandlerFunc(handlers.GetComplaints)).Methods("GET")
	router.Handle("/complaint/{id}", http.HandlerFunc(handlers.GetComplaint)).Methods("GET")
	router.Handle("/createcomplaint", http.HandlerFunc(handlers.CreateComplaint)).Methods("POST")
	router.Handle("/updatecomplaint/{id}", http.HandlerFunc(handlers.UpdateComplaint)).Methods("PUT")
	router.Handle("/removecomplaint/{id}", http.HandlerFunc(handlers.RemoveComplaint)).Methods("DELETE")

	log.Println("Customer Support App API running on Port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
