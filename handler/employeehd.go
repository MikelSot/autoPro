package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
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
	data := model.Employee{}
	err := c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructEmployee, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = isEmailValidEmployee(&data, *e, c)
	if err != nil {
		return err
	}

	areDataValidEmployee(&data)
	err = e.crudExists.Create(&data)
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

	data := model.Employee{}
	err = c.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructEmployee, nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = isEmailValidEmployee(&data, *e, c)
	if err != nil {
		return err
	}

	areDataValidEmployee(&data)
	err = e.crudExists.Update(uint(ID), &data)
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

func isEmailValidEmployee(data *model.Employee, c employeeHd, e echo.Context) error {
	data.Email = strings.TrimSpace(data.Email)

	if !isEmail(data.Email) {
		resp := NewResponse(Error, errorEmailIncorrect, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	exists, _, _, _ := c.crudExists.QueryEmailExists(strings.TrimSpace(data.Email))
	if exists {
		resp := NewResponse(Error, errorEmailExists, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}
	return nil
}

func areDataValidEmployee(data *model.Employee) {
	data.Email = strings.TrimSpace(data.Email)
	data.Turn = strings.TrimSpace(data.Turn)
	data.Workdays = strings.TrimSpace(data.Workdays)
	data.Profession = strings.TrimSpace(data.Profession)
}
