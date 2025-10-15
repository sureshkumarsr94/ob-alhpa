package product_service

import (
	"fmt"
	"github.com/guregu/null"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/dto"
	productDto "infopack.co.in/offybox/app/dto/product"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
)

// GetCategoryObject category list
func (s *productService) GetCategoryObject(categoryID string) (payload productDto.CategoryResponse) {
	category := coreModel.Category{}
	category, _ = category.FindByPrimaryKey(categoryID)
	payload.CategoryID = category.ID
	payload.Code = category.Code
	payload.Name = category.Name
	payload.Description = category.Description
	payload.Sequence = category.Sequence
	payload.Status = category.Status
	if !category.ParentCategoryID.IsZero() {
		payload.ParentCategoryID = category.ParentCategoryID.String
		parentCategory, _ := category.FindByPrimaryKey(category.ParentCategoryID.String)
		parentCategoryPayload := productDto.CategoryResponse{
			CategoryID:       parentCategory.ID,
			Code:             parentCategory.Code,
			Name:             parentCategory.Name,
			Description:      parentCategory.Description,
			Sequence:         parentCategory.Sequence,
			Status:           parentCategory.Status,
			ParentCategoryID: parentCategory.ParentCategoryID.String,
		}
		payload.ParentCategory = &parentCategoryPayload
	}
	return
}

// SaveCategory create or update category details
func (s *productService) SaveCategory(params productDto.CategoryRequest, userObject coreModel.User) (category coreModel.Category, handle dto.HandleError) {
	if params.CategoryID != "" {
		category, _ = category.FindByPrimaryKey(params.CategoryID)
		if category.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid category")
			return
		}
	}

	if params.ParentCategoryID != "" {
		parentCategory, _ := category.FindByPrimaryKey(params.ParentCategoryID)
		if parentCategory.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid parent category")
			return
		}
		category.ParentCategoryID = null.StringFrom(params.ParentCategoryID)
	}

	if userObject.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(userObject.ID)
		category.DistributorID = null.StringFrom(distributorUser.DistributorID)
	}

	category.Code = params.Code
	category.Name = params.Name
	category.Description = params.Description
	category.Sequence = params.Sequence
	category.Status = coreModel.StatusActive
	category, err := category.Save()
	if err != nil {
		handle.Status = -1
		handle.Errors = err
		return
	}

	return
}

// GetProductObject product list
func (s *productService) GetProductObject(productID string) (payload productDto.ProductResponse) {
	product := coreModel.Product{}
	product, _ = product.FindByPrimaryKey(productID)
	payload.ProductID = product.ID
	payload.Code = product.Code
	payload.Name = product.Name
	payload.MRP = product.Mrp
	payload.UOM = product.Uom
	payload.Variant = product.Variant
	payload.Description = product.Description
	payload.Status = product.Status
	return
}

// SaveProduct create or update product details
func (s *productService) SaveProduct(params productDto.ProductRequest, userObject coreModel.User) (product coreModel.Product, handle dto.HandleError) {
	if params.ProductID != "" {
		product, _ = product.FindByPrimaryKey(params.ProductID)
		if product.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("invalid product")
			return
		}
	}

	if userObject.Type == constants.UserTypeDistributor {
		distributorUser := entityModel.DistributorUser{}
		distributorUser, _ = distributorUser.FindByUserId(userObject.ID)
		if distributorUser.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("distributor user not found")
			return
		}
		product.DistributorID = null.StringFrom(distributorUser.DistributorID)
	}

	product.Code = params.Code
	product.Name = params.Name
	product.Mrp = params.MRP
	product.Uom = params.UOM
	product.Variant = params.Variant
	product.Description = params.Description
	product.Status = coreModel.UserStatusActive
	product, err := product.Save()
	if err != nil {
		handle.Status = -1
		handle.Errors = err
		return
	}

	return
}
