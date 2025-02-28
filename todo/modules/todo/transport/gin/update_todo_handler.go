package gintodo

import (
	"net/http"
	"strconv"
	"todos/common"
	"todos/modules/todo/biz"
	"todos/modules/todo/model"
	"todos/modules/todo/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Update(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var body model.TodoUpdate
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}


		store := storage.NewSQLStore(db)
		business := biz.NewUpdateTodoBiz(store)

		

		if err := business.UpdateTodoById(ctx.Request.Context(), id, &body); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}