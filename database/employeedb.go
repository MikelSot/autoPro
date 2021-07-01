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

	_,data,_:= c.QueryEmailExists(employee.Email)
	allData := allDataEmployee(data,employee)
	return &allData, nil
}

func (e *EmployeeDao) GetAll(max int) (*model.Employees, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	employees := model.Employees{}
	DB().Limit(max).Find(&employees)
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

func (e *EmployeeDao) DataEmployeeHome(max int) (*dto.DataEmployeeHomes,error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	employees :=dto.DataEmployeeHomes{}
	DB().Table("clients c").Limit(max).Select(
		"c.id",
		"c.name",
		"c.last_name",
		"c.email",
		"c.picture",
		"c.uri",
		"e.workdays",
		"e.profession",
		"e.role_id",
	).Joins(
		"INNER JOIN employees e on c.email = e.email",
	).Scan(&employees)
	return &employees, nil
}

func (e *EmployeeDao) QueryEmailExists(email string) (bool,  model.Employee, error) {
	employee := model.Employee{}
	values := DB().Select("Email").Find(&employee, "email = ?", email)
	if values.RowsAffected != ZeroRowsAffected {
		return true ,employee ,nil
	}
	return false, model.Employee{},nil
}

func (e *EmployeeDao) QueryEmailEqualsClient(email string) (uint,error) {
	var ID uint
	DB().Table("clients c").Limit(1).Select(
		"c.email",
	).Joins(
		"INNER JOIN employees e on c.email = e.email",
	).Where("c.email = ?", email).Scan(&ID)
	return ID, nil
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