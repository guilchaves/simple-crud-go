package controller

import (
	"go-gin-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

//TODO: Implement the product controller
type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Nintendo Switch 2",
			Price: 500,
		},
	}

	ctx.JSON(http.StatusOK, products)

}
