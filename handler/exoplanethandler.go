package handler

import (
	"NTTHomeTestDemo/middleware"
	"NTTHomeTestDemo/model"
	"NTTHomeTestDemo/service"
	"NTTHomeTestDemo/utility"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ExoplanetHandler struct {
	service service.ExoplanetService
}

func (ex *ExoplanetHandler) NewInitExoHandler(svc service.ExoplanetService, handler *mux.Router) {
	ex.service = svc

	handler.Handle("", middleware.Logging(middleware.AuthMiddleware(http.HandlerFunc(ex.CreateExoplanet)))).Methods(http.MethodPost)
	handler.Handle("", middleware.Logging(middleware.AuthMiddleware(http.HandlerFunc(ex.ListExoplanet)))).Methods(http.MethodGet)
	handler.Handle("/{id}", middleware.Logging(middleware.AuthMiddleware(http.HandlerFunc(ex.GetExoplanet)))).Methods(http.MethodGet)
	handler.Handle("/{id}", middleware.Logging(middleware.AuthMiddleware(http.HandlerFunc(ex.UpdateExoplanet)))).Methods(http.MethodPut)
	handler.Handle("/{id}", middleware.Logging(middleware.AuthMiddleware(http.HandlerFunc(ex.DeleteExoplanet)))).Methods(http.MethodDelete)
	handler.Handle("/{id}/fuel", middleware.Logging(middleware.AuthMiddleware(http.HandlerFunc(ex.EstimateFuel)))).Methods(http.MethodGet)

	// handler.Handler("", middleware.Logging(middleware.AuthMiddleware(http.HandlerFunc(ex.CreateExoplanet)).Methods(http.MethodPost)

}

func (ex *ExoplanetHandler) CreateExoplanet(res http.ResponseWriter, req *http.Request) {

	var exoPlanet model.Exoplanet

	err := json.NewDecoder(req.Body).Decode(&exoPlanet)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// called service method
	data, err := ex.service.CreateExoplanet(exoPlanet)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	res.Header().Set(utility.ContentType, utility.ApplicationJSON)

	res.WriteHeader(http.StatusCreated)

	json.NewEncoder(res).Encode(data)

}

func (ex *ExoplanetHandler) ListExoplanet(res http.ResponseWriter, req *http.Request) {
	// called ListExoplanet service method

	res.Header().Set(utility.ContentType, utility.ApplicationJSON)
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(ex.service.ListExoplanet())

}

func (ex *ExoplanetHandler) GetExoplanet(res http.ResponseWriter, req *http.Request) {

	id := mux.Vars(req)["id"]

	// called GetExoplanet service method
	data, err := ex.service.GetExoplanet(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	res.Header().Set(utility.ContentType, utility.ApplicationJSON)

	res.WriteHeader(http.StatusCreated)

	json.NewEncoder(res).Encode(data)

}

func (ex *ExoplanetHandler) UpdateExoplanet(res http.ResponseWriter, req *http.Request) {

	id := mux.Vars(req)["id"]
	var exoPlanet model.Exoplanet

	err := json.NewDecoder(req.Body).Decode(&exoPlanet)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	// called update service method
	data, err := ex.service.UpdateExoplanet(id, exoPlanet)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
		return
	}
	res.Header().Set(utility.ContentType, utility.ApplicationJSON)

	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(data)

}

func (ex *ExoplanetHandler) DeleteExoplanet(res http.ResponseWriter, req *http.Request) {

	id := mux.Vars(req)["id"]
	// called delete service method
	data, err := ex.service.DeleteExoplanet(id)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	res.Header().Set(utility.ContentType, utility.ApplicationJSON)

	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(data)

}

func (ex *ExoplanetHandler) EstimateFuel(res http.ResponseWriter, req *http.Request) {

	id := mux.Vars(req)["id"]

	crewCapacity, err := strconv.Atoi(req.URL.Query().Get("crew_capacity"))
	if err != nil {
		http.Error(res, "Invalid crew capacity", http.StatusBadRequest)
		return
	}
	// called delete service method

	fuelEstimation, err := ex.service.EstimateFuel(id, crewCapacity)
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotFound)
	}
	res.Header().Set(utility.ContentType, utility.ApplicationJSON)

	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(map[string]float64{"fuel_estimation": fuelEstimation})

}
