package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type TodoItem struct {
	Id			int			`json:"id"`
	Title		string		`json:"title"`
	Description	string 		`json:"description"`
	Status		string		`json:"status"`
	CreatedAt	*time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at,omitempty"`
}

func (TodoItem) TableName() string {return "todo_items"}

type TodoItemCreate struct {
	Id			int			`json:"-" gorm:"column:id;"`
	Title		string		`json:"title" gorm:"column:title;"`			
	Description	string		`json:"description" gorm:"column:description;"`			
	Status		string		`json:"status" gorm:"column:status;"`
}

func (TodoItemCreate) TableName() string {return TodoItem{}.TableName()}

type TodoItemUpdated struct {
	Title		*string		`json:"title" gorm:"column:title;"`			
	Description	*string		`json:"description" gorm:"column:description;"`			
	Status		*string		`json:"status" gorm:"column:status;"`
}

func (TodoItemUpdated) TableName() string {return TodoItem{}.TableName()}

type Paging struct {
	Page		int		`json:"page"`
	Limit		int		`json:"limit"`
	Total		int64 		`json:"total"`
}

func (p *Paging) Process() {
	if p.Page < 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}

func main() {
	// HandleJson()

	db := DBConnect()

	if db == nil {
		fmt.Println("Unable connect to DB")
		return 
	}

	router := gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			todos := v1.Group("/todos")
			{
				todos.POST("", CreateTodo(db))
				todos.GET("/:id", GetDetailTodo(db))
				todos.PATCH("/:id", UpdateTodo(db))
				todos.DELETE("", DeleteTodo(db))
				todos.GET("", GetListTodo(db))
			}
		}
	}

	router.Run("localhost: 5001")
}

func GetListTodo(db * gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var paging Paging

		if err := ctx.ShouldBind(&paging); err != nil {
			ReturnError(err, ctx)
		}

		paging.Process()
		var result []TodoItem

		db = db.Where("status <> ?", "Deleted")

		if err := db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
			ReturnError(err, ctx)
		}

		if err := db.Order("id desc").
					Offset((paging.Page - 1) * paging.Limit).
					Limit(paging.Limit).Find(&result).Error; 
		err != nil {
			ReturnError(err,ctx)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data": result,
			"paging": paging,
		})
	}
}


func DeleteTodo(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ReturnError(err, ctx)
		}

		if err := db.Table(TodoItem{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
			"status": "Deleted",
		}).Error; err != nil {
			ReturnError(err, ctx)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data": true,
		})
	}
}

func UpdateTodo(db* gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		var body TodoItemUpdated


		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ReturnError(err, ctx)
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Where("id = ?", id).Updates(&body).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":err.Error(),
			})

			return
		}


		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data": true,
		})

	}
}

func GetDetailTodo(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var response TodoItem

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ReturnError(err, ctx)
		}

		response.Id = id

		if err := db.First(&response).Error; err != nil {
			ReturnError(err, ctx)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data": response,
		})
	}
}

func ReturnError(err error, ctx *gin.Context) {

	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})

	return 
}

func CreateTodo(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var body TodoItemCreate

		if err := ctx.ShouldBindJSON(&body);  err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		if err := db.Create(&body).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Created",
			"data": body.Id,
		})
	}
}

func DBConnect() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return nil
	}

	dsn := os.Getenv("DB_CONN_STR")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}


func HandleJson() {
	// convert struct to json using json.Marshal
	now := time.Now().UTC()

	todo := TodoItem{
		Id: 1,
		Title: "this is todo one",
		Description: "this is todo one",
		Status: "Doing",
		CreatedAt: &now,
		UpdatedAt: nil,
	}

	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	// convert json to struct using json.Unmarshal

	strJson := "{\"id\":1,\"title\":\"todo 1\",\"description\":\"description todo 1\",\"status\":\"Doing\",\"created_at\":\"2025-02-21T03:50:11.9116221Z\",\"updated_at\":null}"

	var todoStruct TodoItem

	if err := json.Unmarshal([]byte(strJson), &todoStruct); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(todoStruct)
}

