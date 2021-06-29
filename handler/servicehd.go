package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

const (
	errorStructService = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de servicio es inválida"
	serviceCreated     = "Servicio creado correctamente"
	updatedService     = "Servicio actualizado correctamente"
	errorGetAllService = "Hubo un problema al obtener todos los servicios"
)

type serviceHd struct {
	crudQuery IServiceCRUDQuery
}

func NewServiceHd(cq IServiceCRUDQuery) serviceHd {
	return serviceHd{cq}
}

func (s *serviceHd) create(e echo.Context) error {
	data := model.Service{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructService, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = s.crudQuery.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructService, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, serviceCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (s *serviceHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Service{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructService, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = s.crudQuery.Update(uint8(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructService, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedService, nil)
	return e.JSON(http.StatusOK, resp)
}

func (s *serviceHd) getById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := s.crudQuery.GetByID(uint8(ID))
	if err != nil {
		response := newResponse(Error, errorServiceDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (s *serviceHd) getAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := s.crudQuery.GetAll(max)
	if err != nil {
		response := newResponse(Error, errorGetAllService, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}
