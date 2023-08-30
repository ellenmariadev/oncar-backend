package repository

import (
	"context"
	"database/sql"
	"golang-crud/helper"
	"golang-crud/model"
)

type LeadRepositoryImpl struct {
	Db *sql.DB
}

func NewLeadRepository(Db *sql.DB) LeadRepository {
	return &LeadRepositoryImpl{Db: Db}
}

// Save implements LeadRepository
func (l *LeadRepositoryImpl) Save(ctx context.Context, lead model.Lead)  { 

	tx, err := l.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into lead(name, email, phone, vehicleid) values ($1, $2, $3, $4)"
	_, err = tx.ExecContext(ctx, SQL, lead.Name, lead.Email, lead.Phone, lead.VehicleId)
	helper.PanicIfError(err)
}

// FindAll implements LeadRepository
func (v *LeadRepositoryImpl) FindAll(ctx context.Context) []model.Lead {
	tx, err := v.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "select id, name, email, phone, vehicleid from lead"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var leads []model.Lead

	for result.Next() {
		lead := model.Lead{}
		err := result.Scan(&lead.Id, &lead.Name, &lead.Email, &lead.Phone, &lead.VehicleId)
		helper.PanicIfError(err)
		leads = append(leads, lead)
	}

	return leads
}