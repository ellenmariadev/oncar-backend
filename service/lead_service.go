package service

import (
	"context"
	"golang-crud/data/request"
	"golang-crud/data/response"
)

type LeadService interface {
	Create(ctx context.Context, request request.LeadCreateRequest)
	FindAll(ctx context.Context) []response.LeadResponse
}