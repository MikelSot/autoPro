package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
)

type IEmployeeCRUD interface {
	Create(employee *model.Employee) error
	Update(ID uint, employee *model.Employee) error
	GetByID(ID uint) (*model.Employee, error)
	GetAll(Num int) (*model.Employees, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}


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

func (e *EmployeeDao) GetAll(Num int) (*model.Employees, error) {
	employees := model.Employees{}
	DB().Limit(Num).First(&employees)
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

func (e *EmployeeDao) QueryEmailExists(email string) (bool, error) {
	const  ExistsEmail = "Este Email ya existe EMPLEADO"
	employee := model.Employee{}
	values := DB().Select("Email").Find(&employee, "Email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsEmail)
	}
	return false,nil
}

func (e *EmployeeDao) QueryDniExists(dni string) (bool, error) {
	const  ExistsDni = "El DNI ya existe EMPLEADO"
	employee := model.Employee{}
	values := DB().Select("Email").Find(&employee, "Dni = ?", dni)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsDni)
	}
	return false,nil
}

func (e *EmployeeDao) QueryUriExists(uri string) (bool, error) {
	const  ExistsUri = "El DNI ya existe EMPLEADO"
	employee := model.Employee{}
	values := DB().Select("Uri").Find(&employee, "Uri = ?", uri)
	if values.RowsAffected != ZeroRowsAffected {
		return true, errors.New(ExistsUri)
	}
	return false,nil
}
