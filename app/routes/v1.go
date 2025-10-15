package routes

import (
	"github.com/gofiber/fiber/v2"
	distributorController "infopack.co.in/offybox/app/controllers/v1/distributor"
	employeeController "infopack.co.in/offybox/app/controllers/v1/employee"
	journeyPlanController "infopack.co.in/offybox/app/controllers/v1/jouney_plan"
	locationController "infopack.co.in/offybox/app/controllers/v1/location"
	masterController "infopack.co.in/offybox/app/controllers/v1/master"
	orderController "infopack.co.in/offybox/app/controllers/v1/order"
	outletController "infopack.co.in/offybox/app/controllers/v1/outlet"
	productController "infopack.co.in/offybox/app/controllers/v1/product"
	territoryController "infopack.co.in/offybox/app/controllers/v1/territory"
	tripController "infopack.co.in/offybox/app/controllers/v1/trip"
	userController "infopack.co.in/offybox/app/controllers/v1/user"
	"infopack.co.in/offybox/app/logger"
	"infopack.co.in/offybox/app/middlewares"
	distributorService "infopack.co.in/offybox/app/services/distributor"
	employeeService "infopack.co.in/offybox/app/services/employee"
	journeyPlanService "infopack.co.in/offybox/app/services/journey_plan"
	locationService "infopack.co.in/offybox/app/services/location"
	masterService "infopack.co.in/offybox/app/services/master"
	orderService "infopack.co.in/offybox/app/services/order"
	outletService "infopack.co.in/offybox/app/services/outlet"
	productService "infopack.co.in/offybox/app/services/product"
	territoryService "infopack.co.in/offybox/app/services/territory"
	tripService "infopack.co.in/offybox/app/services/trip"
	userService "infopack.co.in/offybox/app/services/user"
)

// SetupRoutesV1 sets up the version 1 routes for the aspire-lms API
// Parameters:
// - app: *fiber.App representing the Fiber application instance
func SetupRoutesV1(app *fiber.App) {
	// Create a new route group for version 1 endpoints with optional authentication middleware
	v1 := app.Group("/v1", middlewares.OptionalAuth())

	// Define a health check endpoint
	v1.Get("/", func(c *fiber.Ctx) error {
		message := "Welcome to offybox api"
		logger.Sugar.Info(message)
		return c.JSON(fiber.Map{
			"status":  1,
			"message": message,
		})
	})

	userSvc := userService.NewUserService()
	territorySvc := territoryService.NewTerritoryService()
	locationSvc := locationService.NewLocationService()
	employeeSvc := employeeService.NewEmployeeService()
	masterSvc := masterService.NewMasterService()
	productSvc := productService.NewProductService()
	distributorSvc := distributorService.NewDistributorService()
	outletSvc := outletService.NewOutletService()
	orderSvc := orderService.NewUserService()
	tripSvc := tripService.NewTripService()
	journeyPlanSvc := journeyPlanService.NewJourneyPlanService()

	// Define the user-related routes
	userRoute := v1.Group("user")
	userRoute.Post("/auth", func(c *fiber.Ctx) error {
		return userController.Login(c, userSvc)
	})

	userRoute.Post("/change-password", func(c *fiber.Ctx) error {
		return userController.ChangePassword(c)
	})

	/*// Define the application-related routes
	applicationRoute := v1.Group("/application")

	// Route for creating a new loan application
	applicationRoute.Post("/", func(c *fiber.Ctx) error {
		return loanController.CreateLoanApplication(c, loanSvc, userSvc, repaymentSvc)
	})

	// Define a restricted route group for application-related operations that require authentication
	restrictedApplicationRoute := applicationRoute.Group("/", middlewares.RequireLoggedIn())

	restrictedApplicationRoute.Get("/", func(c *fiber.Ctx) error {
		return loanController.GetApplicationList(c, loanSvc, userSvc)
	})

	// Route for getting loan application details
	restrictedApplicationRoute.Get("/:applicationId", func(c *fiber.Ctx) error {
		return loanController.GetLoanApplication(c, loanSvc, userSvc)
	})

	// Route for making a repayment
	restrictedApplicationRoute.Post("/:applicationId/repayment", func(c *fiber.Ctx) error {
		return repaymentController.PayRepayment(c, repaymentSvc, userSvc)
	})

	adminApplicationRoute := restrictedApplicationRoute.Group("/:applicationId/approve",
		middlewares.RequireAdmin)

	// Route for approving a loan application
	adminApplicationRoute.Post("/", func(c *fiber.Ctx) error {
		return loanController.ApproveLoanApplication(c, loanSvc, userSvc, repaymentSvc)
	})*/

	distributorRoute := v1.Group("distributor", middlewares.RequireLoggedIn())

	// List all distributor
	distributorRoute.Get("/", func(c *fiber.Ctx) error {
		return distributorController.GetDistributors(c, distributorSvc, userSvc)
	})

	// List all distributor
	distributorRoute.Post("/", func(c *fiber.Ctx) error {
		return distributorController.CreateUpdateDistributor(c, distributorSvc, userSvc)
	})

	// Get distributor Users List
	distributorRoute.Get("/user", func(c *fiber.Ctx) error {
		return distributorController.GetDistributorUsers(c, userSvc)
	})

	// Save distributor User
	distributorRoute.Post("/user", func(c *fiber.Ctx) error {
		return distributorController.CreateUpdateDistributorUser(c, distributorSvc, userSvc)
	})

	distributorRoute.Get("/:distributorId", func(c *fiber.Ctx) error {
		return distributorController.GetDistributor(c, distributorSvc, userSvc)
	})

	orderRoute := v1.Group("order")

	// Create Order
	orderRoute.Post("/", func(c *fiber.Ctx) error {
		return orderController.CreateUpdateOrder(c, orderSvc, userSvc)
	})

	masterRoute := v1.Group("master")
	// Route for territory Type list
	masterRoute.Get("/territory-type", func(c *fiber.Ctx) error {
		return territoryController.TerritoryTypeList(c, userSvc, territorySvc)
	})
	// Route for create Territory Type
	masterRoute.Post("/territory-type", func(c *fiber.Ctx) error {
		return territoryController.CreateUpdateTerritoryType(c, userSvc, territorySvc)
	})
	// Route for territory list
	masterRoute.Get("/territory", func(c *fiber.Ctx) error {
		return territoryController.TerritoryList(c, userSvc, territorySvc)
	})
	// Route for create Territory Type
	masterRoute.Post("/territory", func(c *fiber.Ctx) error {
		return territoryController.CreateUpdateTerritory(c, userSvc, territorySvc)
	})

	// Route for Country List
	masterRoute.Get("/country", func(c *fiber.Ctx) error {
		return locationController.CountryList(c)
	})
	// Route for create Country
	masterRoute.Post("/country", func(c *fiber.Ctx) error {
		return locationController.CreateUpdateCountry(c, locationSvc)
	})

	// Route for State List
	masterRoute.Get("/state", func(c *fiber.Ctx) error {
		return locationController.StateList(c)
	})
	// Route for create State
	masterRoute.Post("/state", func(c *fiber.Ctx) error {
		return locationController.CreateUpdateState(c, locationSvc)
	})

	// Route for City List
	masterRoute.Get("/city", func(c *fiber.Ctx) error {
		return locationController.CityList(c)
	})
	// Route for create City
	masterRoute.Post("/city", func(c *fiber.Ctx) error {
		return locationController.CreateUpdateCity(c, locationSvc)
	})

	// Route for Area List
	masterRoute.Get("/area", func(c *fiber.Ctx) error {
		return locationController.AreaList(c)
	})
	// Route for create Area
	masterRoute.Post("/area", func(c *fiber.Ctx) error {
		return locationController.CreateUpdateArea(c, locationSvc)
	})
	masterRoute.Get("/pin-code/suggest", func(c *fiber.Ctx) error {
		return locationController.PinCodeSuggest(c)
	})

	// Route for Role List
	masterRoute.Get("/role", func(c *fiber.Ctx) error {
		return masterController.RoleList(c, userSvc, masterSvc)
	})
	// Route for Create Role
	masterRoute.Post("/role", func(c *fiber.Ctx) error {
		return masterController.CreateUpdateRole(c, userSvc, masterSvc)
	})

	// Route for Warehouse List
	masterRoute.Get("/warehouse", func(c *fiber.Ctx) error {
		return masterController.WarehouseList(c, userSvc, masterSvc)
	})
	// Route for Warehouse Role
	masterRoute.Post("/warehouse", func(c *fiber.Ctx) error {
		return masterController.CreateUpdateWarehouse(c, userSvc, masterSvc)
	})

	employeeRoute := v1.Group("employee")
	// Route for create Employee
	employeeRoute.Post("/", func(c *fiber.Ctx) error {
		return employeeController.CreateUpdateEmployee(c, employeeSvc)
	})
	// Route for Employee List
	employeeRoute.Get("/", func(c *fiber.Ctx) error {
		return employeeController.EmployeeList(c, employeeSvc)
	})

	productRoute := v1.Group("product")

	// Route for Product list
	productRoute.Get("/", func(c *fiber.Ctx) error {
		return productController.ProductList(c, userSvc, productSvc)
	})
	// Route for create Product
	productRoute.Post("/", func(c *fiber.Ctx) error {
		return productController.CreateUpdateProduct(c, userSvc, productSvc)
	})

	// Route for Product Category list
	productRoute.Get("/category", func(c *fiber.Ctx) error {
		return productController.CategoryList(c, userSvc, productSvc)
	})
	// Route for create product category
	productRoute.Post("/category", func(c *fiber.Ctx) error {
		return productController.CreateUpdateCategory(c, userSvc, productSvc)
	})

	outletRoute := v1.Group("outlet")
	outletRoute.Get("/", func(c *fiber.Ctx) error {
		return outletController.OutletList(c, userSvc, outletSvc)
	})
	// Route for Create Outlet
	outletRoute.Post("/", func(c *fiber.Ctx) error {
		return outletController.CreateUpdateOutlet(c, userSvc, outletSvc)
	})

	// Route for list Outlet Category
	outletRoute.Get("/category", func(c *fiber.Ctx) error {
		return outletController.OutletCategoryList(c, userSvc, outletSvc)
	})
	// Route for Create Outlet Category
	outletRoute.Post("/category", func(c *fiber.Ctx) error {
		return outletController.CreateUpdateOutletCategory(c, userSvc, outletSvc)
	})

	// Route for list Outlet Address
	outletRoute.Get("/:outletID/address", func(c *fiber.Ctx) error {
		return outletController.OutletAddressList(c, outletSvc)
	})
	// Route for Create Outlet Address
	outletRoute.Post("/:outletID/address", func(c *fiber.Ctx) error {
		return outletController.CreateUpdateOutletAddress(c, userSvc, outletSvc)
	})

	tripRoute := v1.Group("trip")
	tripRoute.Get("/", func(c *fiber.Ctx) error {
		return tripController.TripList(c, userSvc, tripSvc)
	})
	// Route for Create Outlet
	tripRoute.Post("/", func(c *fiber.Ctx) error {
		return tripController.CreateUpdateTrip(c, userSvc, tripSvc)
	})

	journeyPlanRoute := v1.Group("journey-plan")
	journeyPlanRoute.Get("/", func(c *fiber.Ctx) error {
		return journeyPlanController.JourneyPlanList(c, journeyPlanSvc)
	})
	// Route for Create Journey Plan
	journeyPlanRoute.Post("/", func(c *fiber.Ctx) error {
		return journeyPlanController.CreateUpdateJourneyPlan(c, journeyPlanSvc)
	})
}
