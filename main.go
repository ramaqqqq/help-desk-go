package main

import (
	"help-desk/controllers"
	helper "help-desk/helpers"
	middleware "help-desk/middlewares"

	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helper.Logger("error", "Error getting env")
	}

	middleUrl := os.Getenv("MIDDLE_URL")

	router := mux.NewRouter()
	router.Use(middleware.JwtAuthentication)

	router.HandleFunc(middleUrl+"/auth/login", controllers.C_Login).Methods("POST")
	router.HandleFunc(middleUrl+"/users/create", controllers.C_AddUsers).Methods("POST")
	router.HandleFunc(middleUrl+"/users/fetch", controllers.C_GetAllUsers).Methods("GET")
	router.HandleFunc(middleUrl+"/users/fetch/{id}", controllers.C_GetSingleUsers).Methods("GET")
	router.HandleFunc(middleUrl+"/users/update/{id}", controllers.C_UpdateUsers).Methods("POST")
	router.HandleFunc(middleUrl+"/users/delete/{id}", controllers.C_DeleteUsers).Methods("DELETE")

	// Category
	router.HandleFunc(middleUrl+"/category/create", controllers.C_AddCategory).Methods("POST")
	router.HandleFunc(middleUrl+"/category/fetch", controllers.C_GetAllCategory).Methods("GET")
	router.HandleFunc(middleUrl+"/category/fetch/{id}", controllers.C_GetSingleCategory).Methods("GET")
	router.HandleFunc(middleUrl+"/category/update/{id}", controllers.C_UpdateCategory).Methods("POST")
	router.HandleFunc(middleUrl+"/category/delete/{id}", controllers.C_DeleteCategory).Methods("DELETE")

	// Sub Category
	router.HandleFunc(middleUrl+"/sub-category/create", controllers.C_AddSubCategory).Methods("POST")
	router.HandleFunc(middleUrl+"/sub-category/fetch", controllers.C_GetAllSubCategory).Methods("GET")
	router.HandleFunc(middleUrl+"/sub-category/fetch/{id}", controllers.C_GetSingleSubCategory).Methods("GET")
	router.HandleFunc(middleUrl+"/sub-category/update/{id}", controllers.C_UpdateSubCategory).Methods("POST")
	router.HandleFunc(middleUrl+"/sub-category/delete/{id}", controllers.C_DeleteSubCategory).Methods("DELETE")

	// Request
	router.HandleFunc(middleUrl+"/request/create", controllers.C_AddRequest).Methods("POST")
	router.HandleFunc(middleUrl+"/request/fetch", controllers.C_GetAllRequest).Methods("GET")
	router.HandleFunc(middleUrl+"/request/fetch/{id}", controllers.C_GetSingleRequest).Methods("GET")
	router.HandleFunc(middleUrl+"/request/update/{id}", controllers.C_UpdateRequest).Methods("POST")
	router.HandleFunc(middleUrl+"/request/delete/{id}", controllers.C_DeleteRequest).Methods("DELETE")
	router.HandleFunc(middleUrl+"/request/summary", controllers.C_GetSummaryRequest).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	port := os.Getenv("PORT")
	handler := c.Handler(router)
	server := new(http.Server)
	server.Handler = handler
	server.Addr = ":" + port
	fmt.Println("Starting server at", server.Addr)
	server.ListenAndServe()
}
