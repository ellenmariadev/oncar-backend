package controller

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/repository"
	"golang-crud/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LeadController struct {
	LeadService        service.LeadService
	VehicleRepository repository.VehicleRepository
}

func NewLeadController(LeadService service.LeadService, VehicleRepository repository.VehicleRepository) *LeadController {
	return &LeadController{
		LeadService:        LeadService,
		VehicleRepository: VehicleRepository,
	}
}

func (controller *LeadController) Create(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {

	
	leadCreateRequest := request.LeadCreateRequest{}
	helper.ReadRequestBody(requests, &leadCreateRequest)

	// Validade name
	if !validateName(leadCreateRequest.Name) {
		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "name",
		}
		helper.WriteResponseBody(writer, webResponse)
		return
	}

	// Validate email
	if !validateEmail(leadCreateRequest.Email) {
		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "email",
		}
		helper.WriteResponseBody(writer, webResponse)
		return
	}

	// Validate phone
	if !validatePhone(leadCreateRequest.Phone) {
		webResponse := response.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "phone",
		}
		helper.WriteResponseBody(writer, webResponse)
		return
	}

	controller.LeadService.Create(requests.Context(), leadCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	helper.WriteResponseBody(writer, webResponse)
}

func (controller *LeadController) FindAll(writer http.ResponseWriter, requests *http.Request, params httprouter.Params) {
	result := controller.LeadService.FindAll(requests.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   result,
	}

	helper.WriteResponseBody(writer, webResponse)
}

