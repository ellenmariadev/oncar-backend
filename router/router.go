package router

import (
	"fmt"
	"golang-crud/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(vehicleController *controller.VehicleController, leadController *controller.LeadController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Oncar API")
	})

	router.POST("/api/vehicle", vehicleController.Create)
	router.GET("/api/vehicles", vehicleController.FindAll)
	router.GET("/api/vehicle/:vehicleId", vehicleController.FindById)
	router.PATCH("/api/vehicle/:vehicleId", vehicleController.Update)
	router.DELETE("/api/vehicle/:vehicleId", vehicleController.Delete)

	
	router.POST("/api/lead", leadController.Create)
	router.GET("/api/leads", leadController.FindAll)

	return router
}