package controllers

import (
	"encoding/json"
	"exoplanets/models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var exoplanets = make(map[string]models.Exoplanet)
var mu sync.Mutex
var idCounter int

func AddExoplanet(w http.ResponseWriter, r *http.Request) {
	var exoplanet models.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if exoplanet.Type != "GasGiant" && exoplanet.Type != "Terrestrial" {
		http.Error(w, "Invalid exoplanet type", http.StatusBadRequest)
		return
	}

	if exoplanet.Distance <= 10 || exoplanet.Distance >= 1000 {
		http.Error(w, "Distance must be between 10 and 1000 light years", http.StatusBadRequest)
		return
	}

	if exoplanet.Radius <= 0.1 || exoplanet.Radius >= 10 {
		http.Error(w, "Radius must be between 0.1 and 10 Earth-radius units", http.StatusBadRequest)
		return
	}

	if exoplanet.Type == "Terrestrial" && (exoplanet.Mass <= 0.1 || exoplanet.Mass >= 10) {
		http.Error(w, "Mass must be between 0.1 and 10 Earth-mass units", http.StatusBadRequest)
		return
	}

	mu.Lock()
	idCounter++
	exoplanet.ID = strconv.Itoa(idCounter)
	exoplanets[exoplanet.ID] = exoplanet
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exoplanet)
}

func ListExoplanets(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	var list []models.Exoplanet
	for _, exoplanet := range exoplanets {
		list = append(list, exoplanet)
	}

	json.NewEncoder(w).Encode(list)
}

func GetExoplanetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mu.Lock()
	exoplanet, exists := exoplanets[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exoplanet)
}

func UpdateExoplanet(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var updatedExoplanet models.Exoplanet
	err := json.NewDecoder(r.Body).Decode(&updatedExoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	mu.Lock()
	exoplanet, exists := exoplanets[id]
	if !exists {
		mu.Unlock()
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	updatedExoplanet.ID = exoplanet.ID
	exoplanets[id] = updatedExoplanet
	mu.Unlock()

	json.NewEncoder(w).Encode(updatedExoplanet)
}

func DeleteExoplanet(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	mu.Lock()
	_, exists := exoplanets[id]
	if exists {
		delete(exoplanets, id)
	}
	mu.Unlock()

	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func EstimateFuel(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	crewCapacityStr := r.URL.Query().Get("crewCapacity")
	crewCapacity, err := strconv.Atoi(crewCapacityStr)
	if err != nil {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}

	mu.Lock()
	exoplanet, exists := exoplanets[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Exoplanet not found", http.StatusNotFound)
		return
	}

	var gravity float64
	if exoplanet.Type == "GasGiant" {
		gravity = 0.5 / (exoplanet.Radius * exoplanet.Radius)
	} else {
		gravity = exoplanet.Mass / (exoplanet.Radius * exoplanet.Radius)
	}

	fuel := float64(exoplanet.Distance) / (gravity * gravity) * float64(crewCapacity)
	json.NewEncoder(w).Encode(map[string]float64{"fuel": fuel})
}
