	package service

	import (
		"context"
		"golang-crud/data/request"
		"golang-crud/data/response"
		"golang-crud/helper"
		"golang-crud/model"
		"golang-crud/repository"
	)

	type VehicleServiceImpl struct {
		VehicleRepository repository.VehicleRepository
	}

	func NewVehicleServiceImpl(VehicleRepository repository.VehicleRepository) VehicleService {
		return &VehicleServiceImpl{VehicleRepository: VehicleRepository}
	}

	// Create implements VehicleService
	func (v *VehicleServiceImpl) Create(ctx context.Context, request request.VehicleCreateRequest) {

		vehicle := model.Vehicle{
			Brand: request.Brand,
			Model: request.Model,
			Price: request.Price,
			Year: request.Year,
		}
		v.VehicleRepository.Save(ctx, vehicle)
	}

	// Delete implements VehicleService
	func (v *VehicleServiceImpl) Delete(ctx context.Context, vehicleId int) {
		vehicle, err := v.VehicleRepository.FindById(ctx, vehicleId)
		helper.PanicIfError(err)
		v.VehicleRepository.Delete(ctx, vehicle.Id)
	}

	// FindAll implements VehicleService
	func (v *VehicleServiceImpl) FindAll(ctx context.Context) []response.VehicleResponse {
		vehicles := v.VehicleRepository.FindAll(ctx)

		var vehicleResp []response.VehicleResponse

		for _, value := range vehicles {
			vehicle := response.VehicleResponse{Id: value.Id, Brand: value.Brand, Model: value.Model, Price: value.Price, Year: value.Year}
			vehicleResp = append(vehicleResp, vehicle)
		}
		return vehicleResp

	}

	// FindById implements VehicleService
	func (v *VehicleServiceImpl) FindById(ctx context.Context, vehicleId int) response.VehicleResponse {
		vehicle, err := v.VehicleRepository.FindById(ctx, vehicleId)
		helper.PanicIfError(err)
		return response.VehicleResponse(vehicle)
	}

	// Update implements VehicleService
	func (v *VehicleServiceImpl) Update(ctx context.Context, request request.VehicleUpdateRequest) {
		vehicle, err := v.VehicleRepository.FindById(ctx, request.Id)
		helper.PanicIfError(err)
		
		vehicle.Brand = request.Brand
		vehicle.Model = request.Model
		vehicle.Price = request.Price
		vehicle.Year = request.Year
		v.VehicleRepository.Update(ctx, vehicle)
	}