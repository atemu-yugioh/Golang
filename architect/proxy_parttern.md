### sử dụng proxy pattern để mở rộng logic 1 method mà không làm thay đổi method gốc

> thực hiện chỉnh sửa logic của hàm FindByEmail mà không làm ảnh hưởng đến code gốc của FindByEmail

> cụ thể là FindByEmail sẽ tìm trên cache, nếu có thì return không thì tìm bằng FindByEmail gốc

> tạo 1 userCacheRepo struct mục đích là để bọc lại (wrap) realRepo usecase.UserQueryRepository (repo gốc) và thêm các thứ mong muốn (vd: cache)

> xong sẽ implement FindByEmail cho thằng userCacheRepo và trong hàm này nó thay đổi hành vi của FindByEmail (có cache)

> bây giờ sẽ dùng 1 complexBuilder để bọc (wrap) simpleBuilder lại với mục đích để thay đổi BuildUserQueryRepo

> bây giờ simpleBuilder.BuildUserQueryRepo() rẽ return về FindByEmail gốc

> bây giờ complexBuilder.BuildUserQueryRepo() rẽ return về FindByEmail có cache

> ==> có thể định nghĩa lại 1 FindByEmail có cache mà không làm thay đổi FindByEmail gốc

> nếu không dùng cách này thì cách thường làm là đi thằng vào FindByEmail update lại code

> bây giờ nếu muốn sử dụng FindByEmail có cache thì call complexBuilder

```go
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

func NewComplexBuilder(simpleBuilder simpleBuilder) complexBuilder {
	return complexBuilder{simpleBuilder: simpleBuilder}
}

type complexBuilder struct {
	simpleBuilder
}

// Proxy design pattern
type userCacheRepo struct {
	realRepo usecase.UserQueryRepository
	cache    map[string]*domain.User
}

func (c userCacheRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	if user, ok := c.cache[email]; ok {
		return user, nil
	}

	user, err := c.realRepo.FindByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	c.cache[email] = user

	return user, nil
}

func (cb complexBuilder) BuildUserQueryRepo() usecase.UserQueryRepository {
	return userCacheRepo{
		realRepo: cb.simpleBuilder.BuildUserQueryRepo(),
		cache:    make(map[string]*domain.User),
	}
}


```
