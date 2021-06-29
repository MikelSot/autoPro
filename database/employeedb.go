package database

import (
	"github.com/MikelSot/autoPro/model"
)


type EmployeeDao struct {
	employeeDao model.Employee
}

func NewEmployeeDao() EmployeeDao {
	return EmployeeDao{}
}

func (e *EmployeeDao) Create(employee *model.Employee) error {
	DB().Create(&employee)
	return nil
}

func (e *EmployeeDao) Update(ID uint, employee *model.Employee) error {
	employeeID := model.Employee{}
	employeeID.ID = ID
	DB().Model(&employeeID).Updates(employee)
	return nil
}

func (e *EmployeeDao) GetByID(ID uint) (*model.Employee, error) {
	employee := model.Employee{}
	DB().First(&employee, ID)
	return &employee, nil
}

func (e *EmployeeDao) GetAll(max int) (*model.Employees, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	employees := model.Employees{}
	DB().Limit(max).First(&employees)
	return &employees, nil
}

func (e *EmployeeDao) DeleteSoft(ID uint) error {
	employee := model.Employee{}
	employee.ID = ID
	DB().Delete(&employee)
	return nil
}

func (e *EmployeeDao) DeletePermanent(ID uint) error {
	employee := model.Employee{}
	employee.ID = ID
	DB().Unscoped().Delete(&employee)
	return nil
}

func (e *EmployeeDao) QueryEmailExists(email string) (bool, model.Client, model.Employee, error) {
	employee := model.Employee{}
	values := DB().Select("Email").Find(&employee, "email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true,model.Client{} ,model.Employee{} ,nil
	}
	return false,model.Client{},model.Employee{},nil
}

func (e *EmployeeDao) QueryDniExists(dni string) (bool, error) {
	employee := model.Employee{}
	values := DB().Select("Email").Find(&employee, "dni = ?", dni)
	if values.RowsAffected != ZeroRowsAffected {
		return true, nil
	}
	return false,nil
}

func (e *EmployeeDao) QueryUriExists(uri string) (bool, error) {
	employee := model.Employee{}
	values := DB().Select("Uri").Find(&employee, "uri = ?", uri)
	if values.RowsAffected != ZeroRowsAffected {
		return true, nil
	}
	return false,nil
}
