package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type VehicleService interface {
	Create(ctx context.Context, request request.VehicleCreateRequest)
	Update(ctx context.Context, request request.VehicleUpdateRequest)
	Delete(ctx context.Context, vehicleId int)
	FindById(ctx context.Context, vehicleId int) response.VehicleResponse
	FindAll(ctx context.Context) []response.VehicleResponse
}