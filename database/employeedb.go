package database

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
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

func (e *EmployeeDao) GetByID(ID uint) (*dto.AllDataEmployee, error) {
	employee := model.Employee{}
	c:= NewClientDao()
	DB().First(&employee, ID)

	_,data,_,_:= c.QueryEmailExists(employee.Email)
	allData := allDataEmployee(data,employee)
	return &allData, nil
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

func allDataEmployee(data model.Client, employee model.Employee) dto.AllDataEmployee {
	allDataEmployee := dto.AllDataEmployee{
		ID      :   data.ID,
		Name       : data.Name,
		LastName   : data.LastName,
		Email      : data.Email,
		Dni        : data.Dni,
		Ruc        : data.Ruc,
		Phone      : data.Phone,
		Picture    : data.Picture,
		Address    : data.Address,
		State      : data.State,
		Uri        : data.Uri,
		BirthDate  : employee.BirthDate,
		Active     : employee.Active,
		Salary     : employee.Salary,
		Turn       : employee.Turn,
		Workdays   : employee.Workdays,
		Profession : employee.Profession,
		BossID     : *employee.BossID,
		RoleID     : employee.RoleID,
	}
	return allDataEmployee
}