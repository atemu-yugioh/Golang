package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"todos/common"
	"todos/middleware"
	"todos/modules/todo/model"
	gintodo "todos/modules/todo/transport/gin"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func main() {
	// HandleJson()

	db := DbConnect()

	if db == nil {
		fmt.Println("Unable connect to DB")
		return
	}


	router := gin.Default()

	router.Use(middleware.Recovery())

	api := router.Group("/api")
	{
		// api.Use(another middleware)
		v1 := api.Group("/v1")
		{
			todos := v1.Group("/todos")
			{
				todos.POST("", gintodo.Create(db))
				todos.GET("/:id", gintodo.GetDetail(db))
				todos.PATCH("/:id", gintodo.Update(db))
				todos.DELETE("/:id", gintodo.Delete(db))
				todos.GET("", gintodo.GetList(db))
			}

		}
	}

	router.GET("/demo-crash-recover", func(ctx *gin.Context) {
	
		// khi ưng dụng crash ở main thresh thì đã có recovery ở middleware control
		fmt.Println([]int{}[0]) // --> demo crash in main goroutine

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/demo-crash-recover-go", func(ctx *gin.Context) {

		/**
		 	- nếu ứng dụng bị crash trong 1 thresh khác (another goroutine) thì không những là thresh đó bị crash
				mà còn crash luôn cả hệ thống (server)
			- cần phải đặt defer common.Recovery() đẻ cover lỗi ở mõi goroutine
			- tại sao ứng dụng đã có middleware Recovery rồi mà vẫn phải cần thêm 1 cái nữa trong common??
				+ middleware Recovery chỉ cover được trong main thresh
				+ muốn cover được ở 1 thresh khác thì cần phài có 1 cái khác
		*/

		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0]) // --> demo crash in another goroutine
		}()

		

		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})


	

	router.Run("localhost:5000")
}


func DbConnect() *gorm.DB {
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


func HandleJson () {
	now := time.Now().UTC()

	status := model.ItemStatusDoing
	todo := model.Todo{
		Title: "todo 1",
		Description: "description todo 1",
		Status: &status,
		SQLModel: common.SQLModel{ // ✅ Khai báo thông qua struct nhúng
			Id:        1,
			CreatedAt: &now,
			UpdatedAt: nil,
		},
	}

	// convert struct to json
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))


	// convert json to struct
	jsonStr := "{\"id\":1,\"title\":\"todo 1\",\"description\":\"description todo 1\",\"status\":\"Doing\",\"created_at\":\"2025-02-21T03:50:11.9116221Z\",\"updated_at\":null}"

	var todoStruct model.Todo
	if err := json.Unmarshal([]byte(jsonStr), &todoStruct); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(todoStruct)

}

