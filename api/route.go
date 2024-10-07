package api

import (
	"btelli-customersupport-app/database"
	"btelli-customersupport-app/handlers"
	"btelli-customersupport-app/middlewares"
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

	// Auth
	router.Handle("/login", http.HandlerFunc(handlers.Login)).Methods("POST")

	// Admin Only
	router.Handle("/assingrole", middlewares.Auth("admin")(http.HandlerFunc(handlers.AssingRoleToUser))).Methods("POST")

	// User
	router.Handle("/register", http.HandlerFunc(handlers.CreateUser)).Methods("POST")
	// Comments
	router.Handle("/comments/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.GetComments))).Methods("GET")
	router.Handle("/createcomment", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.CreateComment))).Methods("POST")
	router.Handle("/updatecommment/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.UpdateComment))).Methods("PUT")
	router.Handle("/removecomment/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.RemoveComment))).Methods("DELETE")

	// Categories
	router.Handle("/categories", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.GetCategories))).Methods("GET")
	router.Handle("/category/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.GetCategory))).Methods("GET")
	router.Handle("/createcategory", middlewares.Auth("admin")(http.HandlerFunc(handlers.CreateCategory))).Methods("POST")
	router.Handle("/updatecategory/{id}", middlewares.Auth("admin")(http.HandlerFunc(handlers.UpdateCategory))).Methods("PUT")
	router.Handle("/removecategory/{id}", middlewares.Auth("admin")(http.HandlerFunc(handlers.RemoveCategory))).Methods("DELETE")

	// Complaints
	router.Handle("/complaints", middlewares.Auth("admin")(http.HandlerFunc(handlers.GetComplaints))).Methods("GET")
	router.Handle("/complaints/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.GetComplaintsById))).Methods("GET")
	router.Handle("/complaint/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.GetComplaint))).Methods("GET")
	router.Handle("/createcomplaint", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.CreateComplaint))).Methods("POST")
	router.Handle("/updatecomplaint/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.UpdateComplaint))).Methods("PUT")
	router.Handle("/removecomplaint/{id}", middlewares.Auth("admin", "help desk", "customer")(http.HandlerFunc(handlers.RemoveComplaint))).Methods("DELETE")

	log.Println("Customer Support App API running on Port", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}
