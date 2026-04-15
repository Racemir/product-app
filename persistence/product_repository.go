package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/Racemir/product-app/domain"
	"github.com/Racemir/product-app/persistence/common"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

// Ürünleri getiren bir yapıya ihtiyacım var
type IProductRepository interface {
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
	AddProduct(product domain.Product) error
	GetById(productId int64) (domain.Product, error)
	DeleteById(productId int64) error
	UpdatePrice(productId int64, newPrice float32) error
}

// PostgreSQL'e bağlanıp, SQL sorgusu atıp verileri çekmeyi bilen asıl yapı
type ProductRepository struct {
	dbPool *pgxpool.Pool
}

func NewProductRepository(dbPool *pgxpool.Pool) IProductRepository {
	return &ProductRepository{
		dbPool: dbPool,
	}
}

func (productRepository *ProductRepository) GetAllProducts() []domain.Product {
	ctx := context.Background()
	productRows, err := productRepository.dbPool.Query(ctx, "Select * from products")

	if err != nil {
		log.Errorf("Error while getting all products %v", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	ctx := context.Background()

	getProductsByStoreNameSql := `Select * from products where store = $1`

	productRows, err := productRepository.dbPool.Query(ctx, getProductsByStoreNameSql, storeName)

	if err != nil {
		log.Errorf("Error while getting all products %v", err)
		return []domain.Product{}
	}
	return extractProductsFromRows(productRows)
}

func (productRepository *ProductRepository) AddProduct(product domain.Product) error {
	ctx := context.Background()

	insert_sql := `Insert into products (name,price,discount,store) VALUES ($1,$2,$3,$4)`

	addNewProduct, err := productRepository.dbPool.Exec(ctx, insert_sql, product.Name, product.Price, product.Discount, product.Store)

	if err != nil {
		log.Error("Failed to add new product", err)
		return err
	}
	log.Info(fmt.Printf("Product added with %v", addNewProduct))
	return nil
}

func extractProductsFromRows(productRows pgx.Rows) []domain.Product {
	var products = []domain.Product{}
	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	for productRows.Next() {
		productRows.Scan(&id, &name, &price, &discount, &store)
		products = append(products, domain.Product{
			Id:       id,
			Name:     name,
			Price:    price,
			Discount: discount,
			Store:    store,
		})
	}
	return products
}

func (productRepository *ProductRepository) GetById(productId int64) (domain.Product, error) {
	ctx := context.Background()

	getByIdSql := `Select * from products where id = $1`

	queryRow := productRepository.dbPool.QueryRow(ctx, getByIdSql, productId)

	var id int64
	var name string
	var price float32
	var discount float32
	var store string

	scanErr := queryRow.Scan(&id, &name, &price, &discount, &store)

	if scanErr != nil && scanErr.Error() == common.NOT_FOUND {
		return domain.Product{}, fmt.Errorf("Product not found with id %v", productId)
	}

	if scanErr != nil {
		return domain.Product{}, fmt.Errorf("Error while getting product with id %v", productId)
	}

	return domain.Product{
		Id:       id,
		Name:     name,
		Price:    price,
		Discount: discount,
		Store:    store,
	}, nil
}

func (productRepository *ProductRepository) DeleteById(productId int64) error {
	ctx := context.Background()

	_, getErr := productRepository.GetById(productId)

	if getErr != nil {
		return errors.New("Product not found")
	}

	delete_sql := `Delete from products where id = $1`

	_, deleteErr := productRepository.dbPool.Exec(ctx, delete_sql, productId)
	if deleteErr != nil {
		return fmt.Errorf("Error while deleting product with id %d : %w", productId, deleteErr)
	}

	log.Info("Product delete succesful")
	return nil
}

func (productRepository *ProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	ctx := context.Background()

	update_sql := `Update products set price = $1 where id = $2`

	_, updateErr := productRepository.dbPool.Query(ctx, update_sql, newPrice, productId)
	if updateErr != nil {
		return fmt.Errorf("Error while updating product with id %d: %w", productId, updateErr)
	}

	log.Infof("Product %d price updated with new price %v", productId, newPrice)
	return nil
}
