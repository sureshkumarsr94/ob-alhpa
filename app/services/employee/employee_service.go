package employee_service

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"infopack.co.in/offybox/app/common/constants"
	"infopack.co.in/offybox/app/common/passwordutil"
	db "infopack.co.in/offybox/app/database"
	"infopack.co.in/offybox/app/dto"
	employeeDto "infopack.co.in/offybox/app/dto/employee"
	masterDto "infopack.co.in/offybox/app/dto/master"
	coreModel "infopack.co.in/offybox/app/models/core"
)

func (s *employeeService) SaveEmployee(tx *gorm.DB, params employeeDto.EmployeeRequest) (employeeID string, handle dto.HandleError) {

	user := coreModel.User{}
	// employee duplicate validation
	if len(params.EmployeeID) > 0 {
		user, _ = user.FindByEmailOrMobile(constants.UserTypeEmployee, params.Email, params.Mobile)
		if len(user.ID) > 0 {
			handle = dto.HandleError{
				Status: -1,
				Errors: fmt.Errorf("employee already exists"),
			}
			return
		}
	}

	employee := coreModel.Employee{}
	role := coreModel.Role{}

	if len(params.EmployeeID) > 0 {
		employee, _ = employee.FindByPrimaryKey(params.EmployeeID)
	}

	// Create user if not available
	if len(employee.UserID) > 0 {
		user, _ = user.FindByPrimaryKey(employee.UserID)
	}
	if len(user.ID) == 0 {
		user.ID = uuid.New().String()
	}
	user.FirstName = params.Name
	user.Username = params.Name
	user.Type = constants.UserTypeEmployee
	user.Email = params.Email
	user.Mobile = params.Mobile
	user.Status = params.Status
	if len(params.Password) > 0 {
		user.Password = passwordutil.HashPassword(params.Password)
	}

	// Save Employee Information
	if len(employee.ID) == 0 {
		employee.ID = uuid.New().String()
	}
	employee.UserID = user.ID
	employee.Name = params.Name
	employee.Code = params.Code
	employee.Mobile = params.Mobile
	employee.Email = params.Email
	employee.Status = params.Status

	if err := tx.Save(&user).Error; err != nil {
		handle = dto.HandleError{
			Status: -4,
			Errors: err,
		}
		return
	}

	if err := tx.Save(&employee).Error; err != nil {
		handle = dto.HandleError{
			Status: -5,
			Errors: err,
		}
		return
	}

	// Validate supervisor ID in employee
	if len(params.SupervisorID) > 0 {
		supervisor, _ := employee.FindByPrimaryKey(params.SupervisorID)
		if len(supervisor.ID) == 0 {
			handle = dto.HandleError{
				Status: -3,
				Errors: fmt.Errorf("supervisor details not found"),
			}
			return
		}

		//Save Supervisor Information
		employeeHierarchy := coreModel.EmployeeHierarchy{}
		employeeHierarchy, _ = employeeHierarchy.FindByEmployeeSupervisor(employee.ID, params.SupervisorID)
		employeeHierarchy.EmployeeID = employee.ID
		employeeHierarchy.SupervisorID = params.SupervisorID
		employeeHierarchy.Status = constants.StatusActive

		if err := tx.Save(&employeeHierarchy).Error; err != nil {
			handle = dto.HandleError{
				Status: -6,
				Errors: err,
			}
			return
		}
	}

	// Validate Role Id and Create User Role Mapping
	if len(params.RoleId) > 0 {
		role, _ = role.FindByPrimaryKey(params.RoleId)
		if len(role.ID) == 0 {
			handle = dto.HandleError{
				Status: -2,
				Errors: fmt.Errorf("role details not found"),
			}
			return
		}
		userRole := coreModel.UserRole{}
		userRole, _ = userRole.FindByUserRole(user.ID, role.ID)
		userRole.UserID = user.ID
		userRole.RoleID = role.ID
		userRole.Status = constants.StatusActive
		if err := tx.Save(&userRole).Error; err != nil {
			handle = dto.HandleError{
				Status: -7,
				Errors: err,
			}
			return
		}
	}

	err := tx.Commit().Error
	if err != nil {
		handle = dto.HandleError{
			Status: -8,
			Errors: err,
		}
	}
	employeeID = employee.ID
	return
}

func (s *employeeService) GetEmployee(employeeID string) (payload employeeDto.EmployeeResponse) {
	employee := coreModel.Employee{}
	employee, _ = employee.FindByPrimaryKey(employeeID)
	payload.EmployeeID = employee.ID
	payload.Code = employee.Code
	payload.Name = employee.Name
	payload.Mobile = employee.Mobile
	payload.Email = employee.Email
	payload.Status = employee.Status

	employeeHierarchy := coreModel.EmployeeHierarchy{}
	var conditions []db.WhereCondition
	conditions = append(conditions, db.WhereCondition{
		Key: fmt.Sprintf("%s.%s", employeeHierarchy.TableName(),
			coreModel.EmployeeHierarchyColumns.EmployeeID),
		Condition: "=",
		Value:     employee.ID,
	})
	conditions = append(conditions, db.WhereCondition{
		Key: fmt.Sprintf("%s.%s", employeeHierarchy.TableName(),
			coreModel.EmployeeHierarchyColumns.Status),
		Condition: "=",
		Value:     constants.StatusActive,
	})
	employeeHierarchyList, _ := employeeHierarchy.FindAll(conditions)
	if len(employeeHierarchyList) > 0 {
		employee, _ = employee.FindByPrimaryKey(employeeHierarchyList[0].EmployeeID)
		supervisorPayload := employeeDto.EmployeeResponse{
			EmployeeID: employee.ID,
			Code:       employee.Code,
			Name:       employee.Name,
			Mobile:     employee.Mobile,
			Email:      employee.Email,
			Status:     employee.Status,
		}
		payload.Supervisor = &supervisorPayload
	}

	userRole := coreModel.UserRole{}
	conditions = []db.WhereCondition{}
	conditions = append(conditions, db.WhereCondition{
		Key: fmt.Sprintf("%s.%s", userRole.TableName(),
			coreModel.UserRoleColumns.UserID),
		Condition: "=",
		Value:     employee.UserID,
	})
	conditions = append(conditions, db.WhereCondition{
		Key: fmt.Sprintf("%s.%s", userRole.TableName(),
			coreModel.UserRoleColumns.Status),
		Condition: "=",
		Value:     constants.StatusActive,
	})
	userRoleList, _ := userRole.FindAll(conditions)
	if len(userRoleList) > 0 {
		role := coreModel.Role{}
		role, _ = role.FindByPrimaryKey(userRoleList[0].RoleID)
		rolePayload := masterDto.RoleResponse{
			RoleID:       role.ID,
			Code:         role.Code,
			Name:         role.Name,
			Description:  role.Description,
			DataAccess:   role.DataAccess,
			ParentRoleID: role.RoleID.Ptr(),
			Status:       role.Status,
		}
		payload.Role = rolePayload
	}
	return
}
