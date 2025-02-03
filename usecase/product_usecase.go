package usecase

import "go-gin-api/model"

// TODO: Implement the product usecase
type productUsecase struct {
    //Repository
}

func NewProductUseCase() productUsecase{
    return productUsecase{}
}

//TODO
func (pu *productUsecase) GetProducts() ([]model.Product, error){
    return []model.Product{}, nil
}
