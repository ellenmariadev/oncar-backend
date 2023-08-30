package main

import (
	"fmt"
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"

	"github.com/gorilla/handlers"
)

func main(){
	fmt.Printf("Start server")

	db := config.DatabaseConnection()

	vehicleRepository := repository.NewVehicleRepository(db)
	vehicleService := service.NewVehicleServiceImpl(vehicleRepository)
	vehicleController := controller.NewVehicleController(vehicleService)
	
	leadRepository := repository.NewLeadRepository(db)
	leadService := service.NewLeadServiceImpl(leadRepository, vehicleRepository)
	leadController := controller.NewLeadController(leadService, vehicleRepository)

	routes := router.NewRouter(vehicleController, leadController)

	
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})

	corsRouter := handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(routes)
	
	server := http.Server{Addr: "localhost:5000", Handler: corsRouter}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}