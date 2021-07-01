package handler

import (
	"github.com/MikelSot/autoPro/database"
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	errorStructEmployee         = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura del empleado es invalida"
	employeeCreated             = "Empleado creado correctamente"
	updatedEmployee             = "Empleado actualizado correctamente"
	errorEmployeeIDDoesNotExist = "El ID del empleado no existe"
	errorGetAllEmployee         = "Hubo un problema al obtener todas los Empleados"
)

type employeeHd struct {
	crudExists IEmployeeCRUDExists
}

func NewEmployeeHd(ce IEmployeeCRUDExists) employeeHd {
	return employeeHd{ce}
}

func (e *employeeHd) Create(c echo.Context) error {
	cl := database.NewClientDao()
	data := dto.DataEmployee{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructEmployee, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	areDataValidEmployee(&data)
	if err, bool := isEmailValidEmployee(&data, *e, c); !bool {
		return err
	}

	regexSpace := regexp.MustCompile(` `)
	dniWithoutSpace := regexSpace.ReplaceAllString(data.Dni, "")
	existsDni,_, _ := cl.QueryDniExists(dniWithoutSpace)
	if existsDni {
		resp := NewResponse(Error, errorDniExists, nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	edit := dto.EditClient{}
	editDataEmployeeClient(&edit, data)

	employee := updateDataEmployee(&data)
	err = e.crudExists.Create(&employee)
	if err != nil {
		resp := NewResponse(Error, errorStructEmployee, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, employeeCreated, nil)
	return c.JSON(http.StatusCreated, resp)
}

func (e *employeeHd) Update(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return c.JSON(http.StatusBadRequest, res)
	}

	data := dto.DataEmployee{}
	err = c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructEmployee, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	areDataValidEmployee(&data)
	if err, bool := isEmailValidEmployeeUpdate(uint(ID),&data, *e, c); !bool {
		return err
	}

	employee := updateDataEmployee(&data)
	err = e.crudExists.Update(uint(ID), &employee)
	if err != nil {
		resp := NewResponse(Error, errorStructEmployee, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	// si se cambia el rol a en empleado, tambien que se cambie en el login (al mismo role que se quiera cambiar)
	resp := NewResponse(Message, updatedEmployee, nil)
	return c.JSON(http.StatusOK, resp)
}

func (e *employeeHd) GetById(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// aqui tambien podemos devolver (dto) los datos del login (a traves del email consulta), en el database
	data, err := e.crudExists.GetByID(uint(ID))
	if err != nil {
		response := NewResponse(Error, errorEmployeeIDDoesNotExist, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	res := NewResponse(Message, ok, data)
	return c.JSON(http.StatusOK, res)
}

func (e *employeeHd) GetAll(c echo.Context) error {
	max, err := strconv.Atoi(c.Param("max"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := e.crudExists.GetAll(max) // lo mismo que el anterior del getByid
	if err != nil {
		response := NewResponse(Error, errorGetAllEmployee, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	res := NewResponse(Message, ok, data)
	return c.JSON(http.StatusOK, res)
}


func (e *employeeHd) DataEmployeeHome(c echo.Context) error {
	max, err := strconv.Atoi(c.Param("max"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	data, err := e.crudExists.DataEmployeeHome(max)
	if err != nil {
		response := NewResponse(Error, errorGetAllEmployee, nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	res := NewResponse(Message, ok, data)
	return c.JSON(http.StatusOK, res)
}

func (e *employeeHd) DeleteSoft(c echo.Context) error {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	// si se elimina un empleado que su role cambie a cliente (mismo database)
	err = e.crudExists.DeleteSoft(uint(ID))
	if err != nil {
		response := NewResponse(Error, errorEmployeeIDDoesNotExist, nil)
		return c.JSON(http.StatusBadRequest, response)
	}

	res := NewResponse(Message, ok, nil)
	return c.JSON(http.StatusOK, res)
}



func isEmailValidEmployee(data *dto.DataEmployee, e employeeHd, c echo.Context) (error, bool) {
	cl := database.NewClientDao()
	if !isEmail(data.Email) {
		resp := NewResponse(Error, errorEmailIncorrect, nil)
		return c.JSON(http.StatusBadRequest, resp), false
	}

	exists,_, _ := e.crudExists.QueryEmailExists(data.Email)
	if exists {
		resp := NewResponse(Error, errorEmailExists, nil)
		return c.JSON(http.StatusBadRequest, resp), false
	}

	existsEmail,_,_:= cl.QueryEmailExists(data.Email)
	if existsEmail {
		resp := NewResponse(Error, errorEmailExists + "en el client", nil)
		return c.JSON(http.StatusBadRequest, resp), false
	}
	return nil, true
}

func isEmailValidEmployeeUpdate(ID uint,data *dto.DataEmployee, e employeeHd, c echo.Context) (error,bool) {
	cl := database.NewClientDao()
	if !isEmail(data.Email) {
		resp := NewResponse(Error, errorEmailIncorrect, nil)
		return c.JSON(http.StatusBadRequest, resp), false
	}

	_,employee, _ := e.crudExists.QueryEmailExists(data.Email)
	if employee.Email == data.Email && employee.ID != ID {
		resp := NewResponse(Error, errorEmailExists, nil)
		return c.JSON(http.StatusBadRequest, resp), false
	}

	id,_:= e.crudExists.QueryEmailEqualsClient(employee.Email)
	edit := dto.EditClient{}
	editDataEmployeeClient(&edit, *data)
	cl.Update(id, &edit)
	return nil, true
}

func editDataEmployeeClient(edit *dto.EditClient, data dto.DataEmployee)  {
	edit.Name =  data.Name
	edit.LastName = data.LastName
	edit.Email = data.Email
	edit.Password = data.Password
	edit.Dni = data.Dni
	edit.Phone = data.Phone
	edit.Picture = data.Picture
	edit.Address = data.Address
	edit.Uri = data.Uri
}

func areDataValidEmployee(data *dto.DataEmployee) {
	data.Name = strings.TrimSpace(data.Name)
	data.LastName = strings.TrimSpace(data.LastName)
	data.Email = strings.TrimSpace(data.Email)
	regexSpace := regexp.MustCompile(` `)
	nameWithoutSpace := regexSpace.ReplaceAllString(data.Name, "")
	data.Password = nameWithoutSpace
	data.Dni = strings.TrimSpace(data.Dni)
	data.Ruc = strings.TrimSpace(data.Ruc)
	data.Phone = strings.TrimSpace(data.Phone)
	data.Picture = strings.TrimSpace(data.Picture)
	data.Address = strings.TrimSpace(data.Address)
	data.State = strings.TrimSpace(data.State)
	data.Turn = strings.TrimSpace(data.Turn)
	data.Workdays = strings.TrimSpace(data.Workdays)
	data.Profession = strings.TrimSpace(data.Profession)
}

func updateDataEmployee(data *dto.DataEmployee) model.Employee {
	employee := model.Employee{
		Email: data.Email,
		BirthDate: data.BirthDate,
		Active: data.Active,
		Salary: data.Salary,
		Turn: data.Turn,
		Workdays: data.Workdays,
		Profession: data.Profession,
		BossID: &data.BossID,
		RoleID: data.RoleID,
	}
	return employee
}