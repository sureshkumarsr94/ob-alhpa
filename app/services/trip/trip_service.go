package trip_service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/guregu/null"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/common/utility"
	"infopack.co.in/offybox/app/dto"
	tripDto "infopack.co.in/offybox/app/dto/trip"
	coreModel "infopack.co.in/offybox/app/models/core"
	saleModel "infopack.co.in/offybox/app/models/sale"
	outletService "infopack.co.in/offybox/app/services/outlet"
	productService "infopack.co.in/offybox/app/services/product"
	userService "infopack.co.in/offybox/app/services/user"
	"time"
)

func (s *tripService) GetTrip(tripID string) (payload tripDto.TripResponse) {
	trip := saleModel.Trip{}
	trip, _ = trip.FindByPrimaryKey(tripID)
	payload.TripID = trip.ID
	payload.UserID = trip.UserID
	user := coreModel.User{}
	payload.User, _ = user.FindByPrimaryKey(trip.UserID)
	payload.StartDate = trip.StartDate.String()
	payload.EndDate = trip.EndDate.Time.String()
	payload.VehicleNumber = trip.VehicleNumber
	payload.VehicleType = trip.VehicleType
	payload.VehicleName = trip.VehicleName
	payload.DriverName = trip.DriverName
	payload.DriverContactNumber = trip.DriverName
	payload.DriverProof = trip.DriverProof
	payload.StartKM = utility.ToString(trip.StartKm)
	payload.EndKM = utility.ToString(trip.EndKm)
	payload.LoadedQuantity = utility.ToString(trip.LoadedQty)
	payload.ReturnedQuantity = utility.ToString(trip.ReturnedQty)
	payload.DamagedQuantity = utility.ToString(trip.DamagedQty)
	payload.Status = trip.Status
	outletSvc := outletService.NewOutletService()
	userSvc := userService.NewUserService()
	productSvc := productService.NewProductService()
	tripItem := saleModel.TripItem{}
	tripItemList, _ := tripItem.FindByTripID(trip.ID)
	for _, tripItem = range tripItemList {
		tripItemResponse := tripDto.TripItemResponse{
			TripItemID:   tripItem.ID,
			InvoiceID:    tripItem.InvoiceID,
			OrderID:      tripItem.OrderID,
			OutletID:     tripItem.OutletID,
			Outlet:       outletSvc.GetOutlet(tripItem.OutletID, userSvc),
			ProductID:    tripItem.ProductID,
			Product:      productSvc.GetProductObject(tripItem.ProductID),
			Quantity:     utility.ToString(tripItem.Qty),
			FreeQuantity: utility.ToString(tripItem.FreeQty),
			Remarks:      tripItem.Remarks,
			Status:       tripItem.Status,
		}
		payload.Items = append(payload.Items, tripItemResponse)
	}
	return
}

func (s *tripService) SaveTrip(tx *gorm.DB, params tripDto.TripRequest, userObject coreModel.User) (tripID string, handle dto.HandleError) {

	trip := saleModel.Trip{}
	if params.TripID != "" {
		trip, _ = trip.FindByPrimaryKey(params.TripID)
		if trip.ID == "" {
			handle.Status = -1
			handle.Errors = fmt.Errorf("trip id is not found")
			return
		}
	} else {
		trip.ID = uuid.New().String()
		trip.CreatedAt = time.Now()
		trip.UpdatedAt = time.Now()
	}

	trip.UserID = userObject.ID
	trip.StartDate = utility.ParseDate(params.StartDate)
	trip.EndDate = null.TimeFrom(utility.ParseDate(params.EndDate))
	trip.VehicleNumber = params.VehicleNumber
	trip.VehicleType = params.VehicleType
	trip.VehicleName = params.VehicleName
	trip.DriverName = params.DriverName
	trip.DriverProof = params.DriverProof
	trip.StartKm = utility.StringToFloat64(params.StartKM)
	trip.EndKm = utility.StringToFloat64(params.EndKM)
	trip.LoadedQty = utility.StringToFloat64(params.LoadedQuantity)
	trip.ReturnedQty = utility.StringToFloat64(params.ReturnedQuantity)
	trip.DamagedQty = utility.StringToFloat64(params.DamagedQuantity)
	trip.Status = params.Status
	if err := tx.Save(&trip).Error; err != nil {
		handle = dto.HandleError{
			Status: -2,
			Errors: err,
		}
		return
	}

	if len(params.Items) == 0 {
		handle = dto.HandleError{
			Status: -3,
			Errors: fmt.Errorf("trip items is empty"),
		}
		return
	}

	for _, element := range params.Items {
		invoice := saleModel.Invoice{}
		invoice, _ = invoice.FindByPrimaryKey(element.InvoiceID)
		if len(invoice.ID) == 0 {
			handle = dto.HandleError{
				Status: -4,
				Errors: fmt.Errorf("invoice id is not found"),
			}
			return
		}
		product := coreModel.Product{}
		product, _ = product.FindByPrimaryKey(element.ProductID)
		if len(product.ID) == 0 {
			handle = dto.HandleError{
				Status: -5,
				Errors: fmt.Errorf("product id is not found"),
			}
		}
		tripItem := saleModel.TripItem{}
		tripItem.TripID = trip.ID
		tripItem.OrderID = invoice.OrderID
		tripItem.InvoiceID = invoice.ID
		tripItem.OutletID = invoice.OutletID
		tripItem.ProductID = element.ProductID
		tripItem.Qty = element.Quantity
		tripItem.FreeQty = element.FreeQuantity
		tripItem.Remarks = element.Remarks
		tripItem.Status = element.Status
		if err := tx.Save(&tripItem).Error; err != nil {
			handle = dto.HandleError{
				Status: -6,
				Errors: err,
			}
			return
		}
	}
	err := tx.Commit().Error
	if err != nil {
		handle = dto.HandleError{
			Status: -7,
			Errors: err,
		}
		return
	}
	return
}
