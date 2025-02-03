package controller

import (
	"go-gin-api/model"
	"go-gin-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
        productUseCase: usecase,
    }
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
