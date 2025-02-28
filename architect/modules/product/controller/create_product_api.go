package controller

import (
	"architect/common"
	"architect/modules/product/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api APIController) CreateProductAPI() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productData domain.ProductCreationDTO

		if err := ctx.Bind(&productData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		productData.Id = common.GenUUID()

		if err := api.createUseCase.CreateProduct(ctx.Request.Context(), &productData); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// response to client
		ctx.JSON(http.StatusCreated, gin.H{"data": productData.Id})

	}
}
