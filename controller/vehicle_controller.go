package controller

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type VehicleController struct {
	VehicleService service.VehicleService
}

func NewVehicleController(VehicleService service.VehicleService) *VehicleController {
	return &VehicleController{VehicleService: VehicleService}
}

func (controller *VehicleController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	vehicleCreateRequest := request.VehicleCreateRequest{}
	helper.ReadRequestBody(requests, &vehicleCreateRequest)

	controller.VehicleService.Create(requests.Context(), vehicleCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *VehicleController) Update(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	vehicleUpdateRequest := request.VehicleUpdateRequest{}
	helper.ReadRequestBody(requests, &vehicleUpdateRequest)

	vehicleId := params.ByName("vehicleId")
	id, err := strconv.Atoi(vehicleId)
	helper.PanicIfError(err)
	vehicleUpdateRequest.Id = id

	controller.VehicleService.Update(requests.Context(), vehicleUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *VehicleController) Delete(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	vehicleId := params.ByName("vehicleId")
	id, err := strconv.Atoi(vehicleId)
	helper.PanicIfError(err)

	controller.VehicleService.Delete(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)

}

func (controller *VehicleController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.VehicleService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *VehicleController) FindById(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	vehicleId := params.ByName("vehicleId")
	id, err := strconv.Atoi(vehicleId)
	helper.PanicIfError(err)

	result := controller.VehicleService.FindById(requests.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)

}