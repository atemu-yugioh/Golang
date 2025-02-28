package controller

import (
	"architect/modules/product/domain"
	"context"
)

type CreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *domain.ProductCreationDTO) error
}

type APIController struct {
	createUseCase CreateProductUseCase
}

func NewApiController(createUseCase CreateProductUseCase) APIController {
	return APIController{createUseCase}
}
