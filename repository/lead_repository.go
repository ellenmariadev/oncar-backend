package repository

import (
	"context"
	"golang-crud/model"
)
	

type LeadRepository interface {
	Save(ctx context.Context, lead model.Lead)
	FindAll(ctx context.Context) []model.Lead
}