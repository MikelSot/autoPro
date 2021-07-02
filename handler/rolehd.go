package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

const (
	roleCreated            = "Role creado correctamente"
	updatedRole            = "Role actualizado correctamente"
	errorStructRole        = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de rol es inválida"
	errorRoleDoesNotExists = "No existe el Role seleccionado"
	errorGetAllRole        = "Hubo un problema al obtener todos los roles"
)

type roleHd struct {
	crud IRoleCRUD
}

func NewRoleHd(c IRoleCRUD) roleHd {
	return roleHd{c}
}

func (r *roleHd) Create(e echo.Context) error {
	data := model.Role{}
	err := e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := NewResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = r.crud.Create(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, roleCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (r *roleHd) Update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Role{}
	err = e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := NewResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = r.crud.Update(uint8(ID), &data)
	if err != nil {
		resp := NewResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, updatedRole, nil)
	return e.JSON(http.StatusOK, resp)
}

func (r *roleHd) GetById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := r.crud.GetByID(uint8(ID))
	if err != nil {
		response := NewResponse(Error, errorRoleDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (r *roleHd) GetAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := r.crud.GetAll(max)
	if err != nil {
		response := NewResponse(Error, errorGetAllRole, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (r *roleHd) DeleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	err = r.crud.DeleteSoft(uint8(ID))
	if err != nil {
		response := NewResponse(Error, errorRoleDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := NewResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}
