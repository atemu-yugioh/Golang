package productmysql

import "gorm.io/gorm"

type MysqlRepo struct {
	db *gorm.DB
}

func NewMysqlRepo(db *gorm.DB) MysqlRepo {
	return MysqlRepo{db}
}
