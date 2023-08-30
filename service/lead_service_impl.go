package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/model"
	"golang-crud/repository"
)

type LeadServiceImpl struct {
	LeadRepository    repository.LeadRepository
	VehicleRepository repository.VehicleRepository
}

func NewLeadServiceImpl(leadRepository repository.LeadRepository, vehicleRepository repository.VehicleRepository) LeadService {
	return &LeadServiceImpl{
		LeadRepository: leadRepository,
		VehicleRepository: vehicleRepository,
	}
}

func (l *LeadServiceImpl) Create(ctx context.Context, request request.LeadCreateRequest) {
	lead := model.Lead{
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		VehicleId: request.VehicleId,
	}

	l.LeadRepository.Save(ctx, lead)
}

// FindAll implements LeadService
func (l *LeadServiceImpl) FindAll(ctx context.Context) []response.LeadResponse {
	leads := l.LeadRepository.FindAll(ctx)

	var leadResp []response.LeadResponse

	for _, value := range leads {
		lead := response.LeadResponse{
			Id:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			Phone:     value.Phone,
			VehicleId: value.VehicleId,
		}
		leadResp = append(leadResp, lead)
	}
	return leadResp
}