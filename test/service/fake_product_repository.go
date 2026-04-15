package service

import (
	"github.com/Racemir/product-app/domain"
	"github.com/Racemir/product-app/persistence"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistence.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeRepository.products
}

func (fakeRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	// todo: your turn

	responseProducts := []domain.Product{}

	for _, x := range fakeRepository.products {
		if x.Store == storeName {
			responseProducts = append(responseProducts, x)
		}
	}
	return responseProducts
}

func (fakeRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeRepository.products = append(fakeRepository.products, domain.Product{
		Id:       int64(len(fakeRepository.products)) + 1,
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

func (fakeRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	// todo: your turn

	for _, x := range fakeRepository.products {
		if x.Id == productId {
			return x, nil
		}
	}
	return domain.Product{}, nil
}

func (fakeRepository *FakeProductRepository) DeleteById(productId int64) error {

	for _, x := range fakeRepository.products {
		if x.Id == productId {
			fakeRepository.products = append(fakeRepository.products[:0], fakeRepository.products[1:]...)
		}
	}
	return nil
}

func (fakeRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {

	for i := range fakeRepository.products {
		if fakeRepository.products[i].Id == productId {
			fakeRepository.products[i].Price = newPrice
		}
	}
	return nil
}
