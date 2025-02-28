package gintodo

import (
	"net/http"
	"todos/common"
	"todos/modules/todo/biz"
	"todos/modules/todo/model"
	"todos/modules/todo/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var body model.TodoCreate

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateItemBiz(store)

		
		if err := business.CreateNewItem(ctx.Request.Context(), &body); err != nil {
			ctx.JSON(http.StatusBadRequest, err)

			return
		}

		ctx.JSON(http.StatusBadRequest, common.SimpleSuccessResponse(body.Id))
	}
}