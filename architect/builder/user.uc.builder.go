package builder

import (
	"architect/common"
	"architect/modules/user/domain"
	"architect/modules/user/infra/cache"
	"architect/modules/user/infra/repository"
	"architect/modules/user/usecase"

	"gorm.io/gorm"
)

type simpleBuilder struct {
	queryDB   *gorm.DB
	commandDB *gorm.DB
	tp        usecase.TokenProvider
}

func NewSimpleBuilder(queryDB *gorm.DB, commandDB *gorm.DB, tp usecase.TokenProvider) simpleBuilder {
	return simpleBuilder{queryDB: queryDB, commandDB: commandDB, tp: tp}
}

func (s simpleBuilder) BuildUserQueryRepo() usecase.UserQueryRepository {
	return repository.NewUserRepo(s.queryDB)
}

func (s simpleBuilder) BuildUserCommandRepo() usecase.UserCommandRepository {
	return repository.NewUserRepo(s.commandDB)
}

func (s simpleBuilder) BuildHasher() usecase.Hasher {
	return &common.Hasher{}
}

func (s simpleBuilder) BuildTokenProvider() usecase.TokenProvider {
	return s.tp
}

func (s simpleBuilder) BuildSessionQueryRepo() usecase.SessionQueryRepository {
	return repository.NewSessionMySQLRepo(s.queryDB)
}

func (s simpleBuilder) BuildSessionCommandRepo() usecase.SessionCommandRepository {
	return repository.NewSessionMySQLRepo(s.commandDB)
}

// Complex builder

type complexBuilder struct {
	simpleBuilder
}

func NewComplexBuilder(simpleBuilder simpleBuilder) complexBuilder {
	return complexBuilder{simpleBuilder: simpleBuilder}
}

func (cb complexBuilder) BuildUserQueryRepo() usecase.UserQueryRepository {
	return cache.NewUserCacheRepo(cb.simpleBuilder.BuildUserQueryRepo(), make(map[string]*domain.User))
}
