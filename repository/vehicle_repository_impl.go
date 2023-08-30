package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-crud/helper"
	"golang-crud/model"
)

type VehicleRepositoryImpl struct {
	Db *sql.DB
}

func NewVehicleRepository(Db *sql.DB) VehicleRepository {
	return &VehicleRepositoryImpl{Db: Db}
}

// Save implements VehicleRepository
func (v *VehicleRepositoryImpl) Save(ctx context.Context, vehicle model.Vehicle) {
	tx, err := v.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into vehicle(brand, model, price, year) values ($1, $2, $3, $4)"
	_, err = tx.ExecContext(ctx, SQL, vehicle.Brand, vehicle.Model, vehicle.Price, vehicle.Year)
	helper.PanicIfError(err)
}

// Delete implements VehicleRepository
func (v *VehicleRepositoryImpl) Delete(ctx context.Context, vehicleId int) {
	tx, err := v.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "delete from vehicle where id =$1"
	_, errExec := tx.ExecContext(ctx, SQL, vehicleId)
	helper.PanicIfError(errExec)
}

// FindAll implements VehicleRepository
func (v *VehicleRepositoryImpl) FindAll(ctx context.Context) []model.Vehicle {
	tx, err := v.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, brand, model, price, year from vehicle"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var vehicles []model.Vehicle

	for result.Next() {
		vehicle := model.Vehicle{}
		err := result.Scan(&vehicle.Id, &vehicle.Brand, &vehicle.Model, &vehicle.Price, &vehicle.Year)
		helper.PanicIfError(err)
		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

// FindById implements VehicleRepository
func (v *VehicleRepositoryImpl) FindById(ctx context.Context, vehicleId int) (model.Vehicle, error) {
	tx, err := v.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id,brand,model,price,year from vehicle where id=$1"
	result, errQuery := tx.QueryContext(ctx, SQL, vehicleId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	vehicle := model.Vehicle{}

	if result.Next() {
		err := result.Scan(&vehicle.Id, &vehicle.Brand, &vehicle.Model, &vehicle.Price, &vehicle.Year)
		helper.PanicIfError(err)
		return vehicle, nil
	} else {
		return vehicle, errors.New("Veículo não encontrado")
	}
}


// Update implements VehicleRepository
func (v *VehicleRepositoryImpl) Update(ctx context.Context, vehicle model.Vehicle) {
	tx, err := v.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "update vehicle set brand=$1, model=$2, price=$3, year=$4 where id=$5"
	_, err = tx.ExecContext(ctx, SQL, vehicle.Brand, vehicle.Model, vehicle.Price, vehicle.Year, vehicle.Id)
	helper.PanicIfError(err)
}