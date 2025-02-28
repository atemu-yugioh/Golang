package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Done bool   `json:"done"`
}


type Todos []Todo


func main() {
	router := gin.Default()

	router.Static("static_file", "./assets")
	router.GET("/heath-check", HeathCheck)

	router.Use(LoggerMiddleware())

	/* UPLOAD HANDLER*/
	// upload singer file
	router.MaxMultipartMemory = 8 // 8 MiB
	router.POST("/upload", func(ctx *gin.Context) {
		// get file from request
		file,_ := ctx.FormFile("file")

		// Upload the file to destination
		ctx.SaveUploadedFile(file, "./assets/upload/" + uuid.New().String() + filepath.Ext(file.Filename))

		ctx.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// upload multiple file
	router.POST("/upload_multiple_file", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			fmt.Println(file.Filename)

			// upload the file to specific des
			ctx.SaveUploadedFile(file, "./assets/upload/" + file.Filename)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": "files uploaded",
		})
	})

	// CRUD TODO

	// router.GET("/api/v1/todos", GetAll)
	// router.GET("/api/v1/todos/:id",GetDetail)
	// router.POST("api/v1/todos", Create)
	// router.PATCH("/api/v1/todos/:id", Update)

	/* GROUP ROUTER */
	api := router.Group("/api")

		v1 := api.Group("/v1")

			v1.GET("todos", GetAll)
			v1.GET("todos/:id", GetDetail)
			v1.POST("todos", Update)
			v1.PATCH("todos/:id", Create)
		

		v2 := api.Group("/v2")
			v2.GET("todos", GetAll)
			v2.GET("todos/:id", GetDetail)
			v2.POST("todos", Update)
			v2.PATCH("todos/:id", Create)

		
	

	router.Run(":5000")
}

func Update (ctx *gin.Context) {
	var body map[string]interface{}
	
	id := ctx.Param("id")
	println("update todo have id", id)

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data": body,
	})
}

func Create(ctx *gin.Context) {
	var req Todo

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	fmt.Println(req)

	var data = gin.H{
		"message": "OK",
		"data": req,
	}

	ctx.JSON(http.StatusOK, data)
}

func GetDetail(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetAll(ctx *gin.Context) {
	var name = ctx.DefaultQuery("name", "none")
	fmt.Println(name)
	var todos = Todos{
		Todo{Name: "learn golang", Id: 1, Done: false},
		Todo{Name: "workout", Id: 2, Done: true},
		Todo{Name: "Check mail", Id: 3, Done: true},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "OK",
		"data": todos,
	})
}

func HeathCheck (ctx *gin.Context) {
	var data = gin.H{
		"message": "OK",
		"data": gin.H{
			"my-sql": "OK",
			"redis": "OK",
		},
	}

	ctx.JSON(http.StatusOK, data)
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("This is global middleware")
		ctx.Next()
	}
}

