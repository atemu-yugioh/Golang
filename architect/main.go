package main

import (
	"architect/builder"
	"architect/component"
	"architect/modules/product/controller"
	productUC "architect/modules/product/domain/usecase"
	productmysql "architect/modules/product/repository/mysql"
	"architect/modules/user/infra/httpservice"
	userUC "architect/modules/user/usecase"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := DBConnect()

	if db == nil {
		fmt.Println("Unable connect to DB")
		return
	}

	router := gin.Default()

	router.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	// Setup dependencies
	repo := productmysql.NewMysqlRepo(db)
	useCase := productUC.NewCreateProductUseCase(repo)
	api := controller.NewApiController(useCase)

	v1 := router.Group("/api/v1")
	{
		product := v1.Group("/products")
		{
			product.POST("", api.CreateProductAPI())
		}
	}

	tokenProvider := component.NewJWTProvider("jwtSecret", 60*60*24*7, 60*60*24*14)

	userBuilder := builder.NewSimpleBuilder(db, db, tokenProvider)

	// userUseCase := userUC.NewUseCase(repository.NewUserRepo((db)), repository.NewSessionMySQLRepo((db)), tokenProvider, &common.Hasher{})

	userUseCase := userUC.UseCaseWithBuilder(userBuilder)

	httpservice.NewUserService(userUseCase).Routes(v1)

	router.Run("localhost:5000")
}

func DBConnect() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	dsn := os.Getenv("DB_CONN_STR")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(db)

	return db
}
