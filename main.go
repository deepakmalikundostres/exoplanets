package main

import (
	"fmt"
	"net/http"

	"exoplanets/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/exoplanets", controllers.AddExoplanet).Methods("POST")
	router.HandleFunc("/exoplanets", controllers.ListExoplanets).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", controllers.GetExoplanetByID).Methods("GET")
	router.HandleFunc("/exoplanets/{id}", controllers.UpdateExoplanet).Methods("PUT")
	router.HandleFunc("/exoplanets/{id}", controllers.DeleteExoplanet).Methods("DELETE")
	router.HandleFunc("/exoplanets/{id}/fuel", controllers.EstimateFuel).Methods("GET")

	port := ":8080"
	fmt.Println("Server is spinning on port", port)
	http.ListenAndServe(port, router)
}
