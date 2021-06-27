package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
)

type IEmployeeCRUD interface {
	Create(employee *model.Employee) error
	Update(ID uint, employee *model.Employee) error
	GetByID(ID uint) (*model.Employee, error)
	GetAll(max int) (*model.Employees, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

// todas las citas en que el usuario intervido (ir a recojer su auto)
//

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
	const  ExistsEmail = "Este Email ya existe EMPLEADO"
	employee := model.Employee{}
	values := DB().Select("Email").Find(&employee, "email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true,model.Client{} ,model.Employee{} ,errors.New(ExistsEmail)
	}
	return false,model.Client{},model.Employee{},nil
}

func (e *EmployeeDao) QueryDniExists(dni string) (bool, error) {
	const  ExistsDni = "El DNI ya existe EMPLEADO"
	employee := model.Employee{}
	values := DB().Select("Email").Find(&employee, "dni = ?", dni)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsDni)
	}
	return false,nil
}

func (e *EmployeeDao) QueryUriExists(uri string) (bool, error) {
	const  ExistsUri = "El DNI ya existe EMPLEADO"
	employee := model.Employee{}
	values := DB().Select("Uri").Find(&employee, "uri = ?", uri)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsUri)
	}
	return false,nil
}
