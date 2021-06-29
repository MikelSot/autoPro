package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo"
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

func (r *roleHd) create(e echo.Context) error {
	data := model.Role{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = r.crud.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, roleCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (r *roleHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Role{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = r.crud.Update(uint8(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructRole, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedRole, nil)
	return e.JSON(http.StatusOK, resp)
}

func (r *roleHd) getById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := r.crud.GetByID(uint8(ID))
	if err != nil {
		response := newResponse(Error, errorRoleDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (r *roleHd) getAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := r.crud.GetAll(max)
	if err != nil {
		response := newResponse(Error, errorGetAllRole, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (r *roleHd) deleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	err = r.crud.DeleteSoft(uint8(ID))
	if err != nil {
		response := newResponse(Error, errorRoleDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}
