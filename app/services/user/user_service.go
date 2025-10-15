package user_service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/utility"
	cfg "infopack.co.in/offybox/app/configs"
	db "infopack.co.in/offybox/app/database"
	"infopack.co.in/offybox/app/dto"
	userDto "infopack.co.in/offybox/app/dto/user"
	"infopack.co.in/offybox/app/models/backup"
	coreModel "infopack.co.in/offybox/app/models/core"
	entityModel "infopack.co.in/offybox/app/models/entity"
	master_service "infopack.co.in/offybox/app/services/master"
	"time"
)

// AccessClaims represents access token JWT claims
type AccessClaims struct {
	Username string `json:"username"`
	UserType string `json:"user_type"`
	UserId   string `json:"user_id"`
	jwt.RegisteredClaims
}

// ValidateCredentials validates the user's credentials for login
// Parameters:
// - request: dto.LoginRequest containing the login details (platform and identifier)
// Returns:
// - dto.LoginResponse with the authentication tokens
// - dto.HandleError with any error that occurred during the process
func (s *userService) ValidateCredentials(request userDto.LoginRequest) (response userDto.LoginResponse, handle dto.HandleError) {
	userModel := coreModel.User{}
	var andCondition, orCondition []db.WhereCondition

	// Add platform validation condition
	andCondition = append(andCondition, db.WhereCondition{
		Key:       coreModel.UserColumns.Type,
		Condition: "=",
		Value:     utility.ValidatePlatform(request.Platform),
	})

	// Add email identifier condition
	orCondition = append(orCondition, db.WhereCondition{
		Key:       coreModel.UserColumns.Email,
		Condition: "=",
		Value:     request.Identifier,
	})

	// Add mobile number identifier condition
	orCondition = append(orCondition, db.WhereCondition{
		Key:       coreModel.UserColumns.Mobile,
		Condition: "=",
		Value:     request.Identifier,
	})

	// Find user by conditions
	userModel, _ = userModel.FindOneByCondition(&andCondition, &orCondition)
	if userModel.ID == "" {
		handle.Status = -2
		handle.Errors = fmt.Errorf("user details not found")
		return
	}

	// Check if the password matches
	if userModel.Password != utility.HashPassword(request.Password) {
		handle.Status = -3
		handle.Errors = fmt.Errorf("user details not found")
		return
	}

	// Generate authentication tokens
	response, err := s.generateAuth(userModel)
	if err != nil {
		handle.Status = -4
		handle.Errors = fmt.Errorf("user details not found")
		return
	}

	return
}

// GenerateAuth generates JWT access and refresh tokens for the authenticated user
// Parameters:
// - users: models.User representing the authenticated user
// Returns:
// - dto.LoginResponse with the authentication tokens and user details
// - error if any error occurred during token generation
func (s *userService) generateAuth(users coreModel.User) (detail userDto.LoginResponse, err error) {
	// Set the access token expiration time to 14 hours
	expireTime := time.Now().Add(time.Hour * 14)

	// Create access claims with user information and token metadata
	accessClaims := AccessClaims{
		fmt.Sprintf("%v %v", users.FirstName, users.LastName),
		users.Type,
		utility.ToString(users.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    cfg.GetConfig().Tenant,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Generate the access token with the specified signing method and secret key
	accessClaimWithSecret := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := accessClaimWithSecret.SignedString([]byte(cfg.GetConfig().JWTAccessSecret))
	if err != nil {
		return
	}

	// Create refresh claims with user information and token metadata
	refreshClaims := AccessClaims{
		fmt.Sprintf("%v %v", users.FirstName, users.LastName),
		users.Type,
		utility.ToString(users.ID),
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    cfg.GetConfig().Tenant,
			NotBefore: jwt.NewNumericDate(expireTime),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
	}

	// Generate the refresh token with the specified signing method and secret key
	refreshClaimWithSecret := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refreshClaimWithSecret.SignedString([]byte(cfg.GetConfig().JWTRefreshSecret))
	if err != nil {
		return
	}

	// Return the generated tokens and user details in the response
	return userDto.LoginResponse{
		Status: 1,
		Data: userDto.TokenDetail{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			UserID:       utility.ToString(users.ID),
			UserName:     users.Username,
			UserType:     users.Type,
			TokenExpires: expireTime.Unix(),
		},
	}, nil
}

// GetUserModel retrieves the user object from the JWT token in the request context
// Parameters:
// - c: *fiber.Ctx representing the Fiber request context
// Returns:
// - models.User representing the user extracted from the JWT token
func (s *userService) GetUserModel(c *fiber.Ctx) (userModel coreModel.User) {
	defer func() {
		if r := recover(); r != nil {
			// Handle any panic that might occur
		}
	}()

	// Initialize an empty AccessClaims struct to store token claims
	access := AccessClaims{}

	// Retrieve the JWT token from the request context
	user := c.Locals("user").(*jwt.Token)
	// Extract claims from the JWT token
	claims := user.Claims.(jwt.MapClaims)

	// Convert the claims map to JSON
	jsonData, err := utility.MapToJSON(claims)
	if err != nil {
		fmt.Println("Error converting map to JSON:", err)
		return
	}

	// Convert the JSON data to AccessClaims struct
	err = utility.JSONToStruct(jsonData, &access)
	if err != nil {
		fmt.Println("Error converting JSON to struct:", err)
		return
	}

	// Find the user by primary key (user ID) extracted from the claims
	userModel, _ = userModel.FindByPrimaryKey(access.UserId)
	return
}

func (s *userService) GetUserObject(userId string) (object userDto.UserObject) {
	userModel := coreModel.User{}
	userModel, _ = userModel.FindByPrimaryKey(userId)

	if userModel.ID != "" {
		object.UserID = userModel.ID
		object.AssociateUserId = userModel.AssociateUserId
		object.Username = userModel.Username
		object.Firstname = userModel.FirstName
		object.Lastname = userModel.LastName
		object.Mobile = userModel.Mobile
		object.UserType = userModel.Type
		object.Email = userModel.Email

		if userModel.Type == constants.UserTypeDistributor {
			distributorUser := entityModel.DistributorUser{}
			distributorUser, _ = distributorUser.FindByUserId(userModel.ID)
			object.DistributorId = distributorUser.DistributorID
		}

		userRole := coreModel.UserRole{}
		userRole, _ = userRole.FindByUser(userModel.ID)
		if userRole.RoleID != "" {
			masterService := master_service.NewMasterService()
			object.Role = masterService.GetRoleObject(userRole.RoleID)
		}

	}

	return
}

// AllocateEmployeeForProcess assigns the least loaded employee to the loan application
// Parameters:
// - application: models.LoanApplication representing the loan application
// Returns:
// - models.LoanApplicationParticipant representing the assigned employee as a participant
// - error if an employee could not be found or assigned
func (s *userService) AllocateEmployeeForProcess(application backup.LoanApplication) (participant backup.LoanApplicationParticipant, err error) {
	/*user := coreModel.User{}

	// Find the least loaded employee
	employeeId, _ := user.FindLeastLoadedEmployee()

	// Retrieve the employee details by primary key
	user, _ = user.FindByPrimaryKey(employeeId)
	if user.ID == "" {
		err = fmt.Errorf("employee not configured")
		return
	}

	// Create a new loan application participant for the employee
	participant = backup.LoanApplicationParticipant{
		ParticipantID:   uuid.New().String(),
		ApplicationID:   application.ApplicationID,
		ParticipantType: constants.UserTypeEmployee,
		UserID:          user.ID,
	}*/

	return
}

// GetApplicationUser retrieves or creates a user based on the provided UserObject
// Parameters:
// - object: dto.UserObject containing user details
// Returns:
// - models.User representing the user found or created
func (s *userService) GetApplicationUser(object dto.UserObject) (userModel coreModel.User) {
	var andCondition, orCondition []db.WhereCondition

	// Add condition to match the user type as 'Customer'
	andCondition = append(andCondition, db.WhereCondition{
		Key:       coreModel.UserColumns.Type,
		Condition: "=",
		Value:     constants.UserTypeDistributor,
	})

	// Add conditions to match either user email or mobile number
	orCondition = append(orCondition, db.WhereCondition{
		Key:       coreModel.UserColumns.Email,
		Condition: "=",
		Value:     object.UserEmail,
	})

	orCondition = append(orCondition, db.WhereCondition{
		Key:       coreModel.UserColumns.Mobile,
		Condition: "=",
		Value:     object.MobileNumber,
	})

	// Find the user by the specified conditions
	userModel, _ = userModel.FindOneByCondition(&andCondition, &orCondition)
	return
}
