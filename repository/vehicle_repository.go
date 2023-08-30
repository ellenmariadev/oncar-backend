package repository

import (
	"context"
	"golang-crud/model"
)
	

type VehicleRepository interface {
	Save(ctx context.Context, vehicle model.Vehicle)
	Update(ctx context.Context, vehicle model.Vehicle)
	Delete(ctx context.Context, vehicleId int)
	FindById(ctx context.Context, vehicleId int)(model.Vehicle, error)
	FindAll(ctx context.Context) []model.Vehicle
}