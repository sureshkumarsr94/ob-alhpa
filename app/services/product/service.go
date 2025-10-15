package product_service

import (
	"infopack.co.in/offybox/app/dto"
	productDto "infopack.co.in/offybox/app/dto/product"
	coreModel "infopack.co.in/offybox/app/models/core"
)

// ProductService defines the interface for product-related operations
type ProductService interface {

	// GetCategoryObject category list
	GetCategoryObject(categoryID string) (payload productDto.CategoryResponse)

	// SaveCategory save the category
	SaveCategory(params productDto.CategoryRequest, userObject coreModel.User) (category coreModel.Category, handle dto.HandleError)

	// GetProductObject product list
	GetProductObject(productID string) (payload productDto.ProductResponse)

	// SaveProduct save the product
	SaveProduct(params productDto.ProductRequest, userObject coreModel.User) (product coreModel.Product, handle dto.HandleError)
}

// productService is an implementation of ProductService
type productService struct{}

// NewProductService returns a new instance of ProductService
func NewProductService() ProductService {
	return &productService{}
}
