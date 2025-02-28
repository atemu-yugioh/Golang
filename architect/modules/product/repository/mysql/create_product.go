package productmysql

import (
	"architect/modules/product/domain"
	"context"
)

func (repo MysqlRepo) CreateProduct(ctx context.Context, prod *domain.ProductCreationDTO) error {
	if err := repo.db.Table(prod.TableName()).Create(&prod).Error; err != nil {
		return err
	}

	return nil
}
