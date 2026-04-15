package response

import "github.com/Racemir/product-app/domain"

type ErrorResponse struct {
	ErrorDescription string `json:"errorDescription"`
}

type ProdustResponse struct {
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Discount float32 `json:"discount"`
	Store    string  `string:"store"`
}

func ToResponse(product domain.Product) ProdustResponse {
	return ProdustResponse{
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	}
}

func ToResponseList(products []domain.Product) []ProdustResponse {
	var productResponseList = []ProdustResponse{}
	for _, product := range products {
		productResponseList = append(productResponseList, ToResponse(product))
	}
	return productResponseList
}
