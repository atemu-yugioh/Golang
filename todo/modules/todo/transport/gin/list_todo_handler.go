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

func GetList(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {

		var pagination common.Paging

		if err := ctx.ShouldBind(&pagination); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}


		pagination.Process()

		var filter model.Filter

		if err := ctx.ShouldBind(&filter); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))

			return 
		}


		store := storage.NewSQLStore(db)
		business := biz.NewListTodoBiz(store)

		result, err := business.ListTodo(ctx.Request.Context(), &filter, &pagination)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}


		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, pagination, filter))
	}
}