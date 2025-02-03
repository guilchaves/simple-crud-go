package usecase

import "go-gin-api/model"

// TODO: Implement the product usecase
type ProductUsecase struct {
    //Repository
}

func NewProductUseCase() ProductUsecase{
    return ProductUsecase{}
}

//TODO
func (pu *ProductUsecase) GetProducts() ([]model.Product, error){
    return []model.Product{}, nil
}
