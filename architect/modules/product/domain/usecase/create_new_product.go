package usecase

import (
	"architect/modules/product/domain"
	"context"
	"strings"
)

// IUseCase
type ICreateProductUseCase interface {
	CreateProduct(ctx context.Context, prod *domain.ProductCreationDTO) error
}

// Business Logic

func NewCreateProductUseCase(repo CreateProductRepository) CreateProductUseCase {
	return CreateProductUseCase{
		repo: repo,
	}
}

type CreateProductUseCase struct {
	repo CreateProductRepository
}

func (uc CreateProductUseCase) CreateProduct(ctx context.Context, prod *domain.ProductCreationDTO) error {
	prod.Name = strings.TrimSpace(prod.Name)

	if prod.Name == "" {
		return domain.ErrProductNameCannotBeBlank
	}

	if err := uc.repo.CreateProduct(ctx, prod); err != nil {
		return err
	}

	return nil
}

// IRepo

type CreateProductRepository interface {
	CreateProduct(ctx context.Context, prod *domain.ProductCreationDTO) error
}
