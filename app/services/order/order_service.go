package order_service

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	db "infopack.co.in/offybox/app/database"
	"infopack.co.in/offybox/app/dto"
	orderDto "infopack.co.in/offybox/app/dto/order"
	userDto "infopack.co.in/offybox/app/dto/user"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
	saleModel "infopack.co.in/offybox/app/models/sale"
	"log"
	"time"
)

func (o orderService) CreateUpdateOrder(tx *gorm.DB, request orderDto.OrderRequest, userObject userDto.UserObject) (orderId string, handle dto.HandleError) {

	order := saleModel.Order{}

	outlet := entityModel.Outlet{}
	outlet, _ = outlet.FindByPrimaryKey(request.OutletID)
	log.Println("userObject.DistributorId :: ", userObject.DistributorId)
	if outlet.ID == "" || (!outlet.DistributorID.IsZero() && (outlet.DistributorID.String != userObject.DistributorId)) {
		handle = dto.HandleError{
			Status: -2,
			Errors: fmt.Errorf("outlet critiria does not match"),
		}
		return
	}

	outletAddress := entityModel.OutletAddress{}
	outletAddress, _ = outletAddress.FindByPrimaryKey(request.OutletAddressID)
	if outletAddress.ID == "" || outletAddress.OutletID != outlet.ID {
		handle = dto.HandleError{
			Status: -2,
			Errors: fmt.Errorf("outlet address critiria does not match"),
		}
		return
	}

	if request.OrderID != "" {
		order, _ = order.FindByPrimaryKey(request.OrderID)
		if order.ID == "" {
			handle = dto.HandleError{
				Status: -1,
				Errors: fmt.Errorf("order id does not exist"),
			}
			return
		}
	} else {
		order.ID = uuid.New().String()
		order.OutletID = outlet.ID
		order.CreatedBy = userObject.UserID
		order.DistributorID = userObject.DistributorId
		order.Status = saleModel.OrderStatusPending
	}

	order.OutletAddressID = outletAddress.ID
	if order.Type == saleModel.OrderTypeOrder && request.Type == saleModel.OrderTypeQuotation {
		handle = dto.HandleError{
			Status: -1,
			Errors: fmt.Errorf("order is already passed quotation stage"),
		}
		return

	}

	order.Type = request.Type

	if request.Type == saleModel.OrderTypeQuotation && order.QuotationDate.IsZero() {
		order.QuotationDate = time.Now()
	}

	order.Remarks = request.Remarks
	order.Longitude = request.Longitude
	order.Latitude = request.Latitude

	var itemCount int
	var netAmount, totalAmount, taxAmount float64
	var orderItems []saleModel.OrderItem
	for _, itemRequest := range request.Item {
		item := saleModel.OrderItem{}
		item, _ = item.FindByPrimaryKey(itemRequest.OrderItemID)
		if item.ID == "" {
			item.ID = uuid.New().String()
			item.OrderID = order.ID
		}

		product := coreModel.Product{}
		var productCondition []db.WhereCondition
		productCondition = append(productCondition, db.WhereCondition{Key: coreModel.ProductColumns.ID, Value: itemRequest.ProductID, Condition: "="})

		if userObject.DistributorId != "" {
			productCondition = append(productCondition, db.WhereCondition{Key: coreModel.ProductColumns.DistributorID, Value: userObject.DistributorId, Condition: "="})
		}

		product, _ = product.FindOneWithCondition(productCondition)

		if product.ID == "" {
			handle = dto.HandleError{
				Status: -1,
				Errors: fmt.Errorf("product does not belonged to you"),
			}
			return
		}

		item.ProductID = product.ID
		item.Price = product.Mrp
		item.Qty = itemRequest.Quantity
		item.NetAmount = float64(item.Qty) * item.Price
		item.TotalAmount = item.NetAmount + item.TaxAmount
		item.Status = itemRequest.Status

		itemCount += item.Qty
		netAmount += item.NetAmount
		taxAmount += item.TaxAmount
		totalAmount += item.TotalAmount

		orderItems = append(orderItems, item)
	}

	order.ItemCount = itemCount
	order.NetAmount = netAmount
	order.TaxAmount = taxAmount
	order.TotalAmount = totalAmount

	if err := tx.Save(&order).Error; err != nil {
		handle = dto.HandleError{
			Status: -1,
			Errors: err,
		}
		return
	}

	if err := tx.Save(&orderItems).Error; err != nil {
		handle = dto.HandleError{
			Status: -1,
			Errors: err,
		}
		return
	}

	if err := tx.Commit().Error; err != nil {
		handle = dto.HandleError{
			Status: -1,
			Errors: err,
		}
		return
	}

	orderId = order.ID

	return
}
