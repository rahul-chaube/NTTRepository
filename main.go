package main

import (
	"NTTHomeTestDemo/handler"
	"NTTHomeTestDemo/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// Invoking user Handler

	userRouter := router.PathPrefix("/user").Subrouter()

	// Initilizing user Handler

	userHandler := handler.UserHandler{}
	userHandler.NewUserHandler(userRouter)

	exoplanets := router.PathPrefix("/exoplanets").Subrouter()

	//init Service
	exoplanetsService := service.InitExpoServiceInit()

	exoplanetHandler := handler.ExoplanetHandler{}
	exoplanetHandler.NewInitExoHandler(*exoplanetsService, exoplanets)

	log.Fatal(http.ListenAndServe(":80", router))

}
