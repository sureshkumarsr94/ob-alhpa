package employee_service

import (
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/dto"
	employeeDto "infopack.co.in/offybox/app/dto/employee"
)

// EmployeeService defines the interface for employee-related operations
type EmployeeService interface {
	// SaveEmployee save the employee
	SaveEmployee(tx *gorm.DB, params employeeDto.EmployeeRequest) (employeeId string, handle dto.HandleError)

	// GetEmployee fetch Employee Details
	GetEmployee(employeeId string) (payload employeeDto.EmployeeResponse)
}

// employeeService is an implementation of EmployeeService
type employeeService struct{}

// NewEmployeeService returns a new instance of EmployeeService
func NewEmployeeService() EmployeeService {
	return &employeeService{}
}
