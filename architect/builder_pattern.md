### sử dụng builder pattern để build tham số của Contructor UserUseCase (khi có nhiều tham số đầu vào)

> ưu điểm:
> khi có thêm các usecase mới (logout, changePassword ...) hàm dựng NewUseCase nhận thêm trực tiếp nhiều tham số ==> khó đọc, bảo trì

> tất cả những nơi gọi NewUseCase đều phải cập nhật tham số tuân theo contructor của nó

> sử dụng builder pattern

> tham số chỉ nhận vào 1 builder chứa tất cả các dependency của constructor

> khi có thêm các usecase mới (logout, changePassword ...) chỉ cần khởi tạo thêm các dependency trong builder

> tất cả những nơi gọi UseCaseWithBuilder ko phải cập nhật tham số (bởi vì nó nhận vào 1 builder, chỉ việc thay đổi builder là được)

```js

type Builder interface {
	BuildUserQueryRepo() UserQueryRepository
	BuildUserCommandRepo() UserCommandRepository
	BuildHasher() Hasher
	BuildTokenProvider() TokenProvider
	BuildSessionQueryRepo() SessionQueryRepository
	BuildSessionCommandRepo() SessionCommandRepository
}

func UseCaseWithBuilder(b Builder) UseCase {

	return &useCase{
		registerUC: NewRegisterUC(b.BuildUserQueryRepo(), b.BuildUserCommandRepo(), b.BuildHasher()),
		LoginUC:    NewLoginUC(b.BuildUserQueryRepo(), b.BuildSessionCommandRepo(), b.BuildTokenProvider(), b.BuildHasher()),
	}

}

func NewUseCase(userRepo UserRepository, sessionRepo SessionRepository, tokenProvider TokenProvider, hasher Hasher) UseCase {
	return &useCase{
		registerUC: NewRegisterUC(userRepo, userRepo, hasher),
		LoginUC:    NewLoginUC(userRepo, sessionRepo, tokenProvider, hasher),
	}
}

```

```js
package builder


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



```
