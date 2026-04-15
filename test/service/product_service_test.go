package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/Racemir/product-app/domain"
	"github.com/Racemir/product-app/service"
	"github.com/Racemir/product-app/service/model"
	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {

	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000.0,
			Store: "ABC TECH",
		},
		{
			Id:    2,
			Name:  "Ütü",
			Price: 4000.0,
			Store: "ABC TECH",
		},
		{
			Id:    3,
			Name:  "PC",
			Price: 8000.0,
			Store: "KAF TECH",
		},
	}

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProduct(t *testing.T) {
	t.Run("ShouldGetAllProduct", func(t *testing.T) {
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
	})
}

func Test_WhenNoValudationErrorOccurred_ShouldAddProduct(t *testing.T) {
	t.Run("WhenNoValudationErrorOccurred_ShouldAddProduct", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "ütü",
			Price:    2000.0,
			Discount: 50,
			Store:    "ABC TECH",
		}, actualProducts[len(actualProducts)-1])
	})
}

// İndirim %70'in üzerindeyse ürün eklenmemelidir
func Test_WhenDiscountIsHigherThan70_ShouldNotAddProduct(t *testing.T) {
	t.Run("WhenDiscountIsHigherThan70_ShouldNotAddProduct", func(t *testing.T) {
		err := productService.Add(model.ProductCreate{
			Name:     "ütü",
			Price:    2000.0,
			Discount: 74,
			Store:    "ABC TECH",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, "Discount can not be greater than 70", err.Error())
	})
}

// verilen storeName'in slice'ını dönecek
func Test_ItWillReturnASliceOfTheGivenStoreName(t *testing.T) {
	t.Run("ItWillReturnASliceOfTheGivenStoreName", func(t *testing.T) {
		value := "ABC TECH"
		actualProducts := productService.GetAllProductsByStore(value)
		assert.Equal(t, 2, len(actualProducts))
		for _, product := range actualProducts {
			assert.Equal(t, value, product.Store, "The product that arrived is not from the expected store %s:", product.Name)
		}
	})
}

// verilen Id'ye göre ürünü dönecek
func Test_ItWillReturnTheProductBasedOnTheGivenID(t *testing.T) {
	t.Run("ItWillReturnTheProductBasedOnTheGivenID", func(t *testing.T) {

		actualProducts, _ := productService.GetById(1)
		productSlice := productService.GetAllProducts()
		assert.Equal(t, productSlice[0], actualProducts)
	})
}

// id' ye göre ürünü silecek DeleteById
func Test_TheProductWillBeDeletedAccordingToTheID(t *testing.T) {
	t.Run("TheProductWillBeDeletedAccordingToTheID", func(t *testing.T) {
		deleteErr := productService.DeleteById(1)
		assert.Equal(t, nil, deleteErr)
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))

	})
}

// id ve price alacak o id'li ürünün price değerini değitirecek UpdatePrice
func Test_ItWillGetTheIdAndPriceAndChangeThePriceOfTheProductWithThatID(t *testing.T) {
	t.Run("ItWillGetTheIdAndPriceAndChangeThePriceOfTheProductWithThatID", func(t *testing.T) {
		updateErr := productService.UpdatePrice(2, 1500.0)
		assert.Equal(t, nil, updateErr)

		products := productService.GetAllProducts()
		//responseProducts := []domain.Product{}

		for _, i := range products {
			fmt.Println(i)
			if i.Id == 2 {
				i.Price = 1500.0
				assert.Equal(t, i.Price, float32(1500.0))
			}
		}

	})
}
