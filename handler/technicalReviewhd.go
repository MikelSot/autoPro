package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

const (
	reviewCreated            = "Revisión técnica creada correctamente."
	errorStructReview        = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de Revisión técnica es inválida"
	updatedReview            = "Revisión técnica actualizada correctamente"
	errorReviewDoesNotExists = "No existe el Revisión técnica seleccionado"
	errorGetAllReview        = "Hubo un problema al obtener todas revisiónes técnicas"
)

type technicalReviewHd struct {
	crudQuery ITechnicalReviewCRUDQuery
}

func NewTechnicalReviewHd(cq ITechnicalReviewCRUDQuery) technicalReviewHd {
	return technicalReviewHd{cq}
}

func (t *technicalReviewHd) create(e echo.Context) error {
	data := model.TechnicalReview{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructReview, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	areDataValidTechnicalReview(&data)
	err = t.crudQuery.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructReview, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, reviewCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (t *technicalReviewHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.TechnicalReview{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructReview, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	areDataValidTechnicalReview(&data)
	err = t.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructReview, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedReview, nil)
	return e.JSON(http.StatusOK, resp)
}

func (t *technicalReviewHd) getById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := t.crudQuery.GetByID(uint(ID))
	if err != nil {
		response := newResponse(Error, errorReviewDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (t *technicalReviewHd) getAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := t.crudQuery.GetAll(max)
	if err != nil {
		response := newResponse(Error, errorGetAllReview, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (t *technicalReviewHd) allReviewClient(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := t.crudQuery.AllReviewClient(uint(ID) ,max)
	if err != nil {
		response := newResponse(Error, errorGetAllReview, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func areDataValidTechnicalReview(data *model.TechnicalReview) {
	data.CarFeatures = strings.TrimSpace(data.CarFeatures)
	data.CommonFaults = strings.TrimSpace(data.CommonFaults)
	data.TechnicalAnalysis = strings.TrimSpace(data.TechnicalAnalysis)
	data.VehicleType = strings.TrimSpace(data.VehicleType)
	data.Requirements = strings.TrimSpace(data.Requirements)
	data.ServicesVehicle = strings.TrimSpace(data.ServicesVehicle)
	data.Arrangements = strings.TrimSpace(data.Arrangements)
}
