package gintodo

import (
	"net/http"
	"strconv"
	"todos/common"
	"todos/modules/todo/biz"
	"todos/modules/todo/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewDeleteTodoBiz(store)
		
		

		if err := business.DeleteTodoById(ctx.Request.Context(), id); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}