package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

const (
	errorStructPaymentMethod = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de metodo de pago es inválida"
	paymentMethodCreated     = "Metodo de pago creado correctamente"
	updatedPaymentMethod = "Metodo de pago actualizado correctamente"
	errorPaymentMethodDoesNotExists = "No existe el Metodo de pago seleccionado"
	errorGetAllPaymentMethod = "Hubo un problema al obtener todos los Metodos de pago"
)

type paymentMethodHd struct {
	crud IPaymentMethodCRUD
}

func NewPaymentMethodHd(c IPaymentMethodCRUD) paymentMethodHd {
	return paymentMethodHd{c}
}

func (p *paymentMethodHd) create(e echo.Context) error {
	data := model.PaymentMethod{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructPaymentMethod, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = p.crud.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructPaymentMethod, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, paymentMethodCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (p *paymentMethodHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.PaymentMethod{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructPaymentMethod, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = p.crud.Update(uint8(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructPaymentMethod, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedPaymentMethod, nil)
	return e.JSON(http.StatusOK, resp)
}

func (p *paymentMethodHd) getById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := p.crud.GetByID(uint8(ID))
	if err != nil {
		response := newResponse(Error, errorPaymentMethodDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (p *paymentMethodHd) getAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := p.crud.GetAll(max)
	if err != nil {
		response := newResponse(Error, errorGetAllPaymentMethod, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (p *paymentMethodHd) deleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	err = p.crud.DeleteSoft(uint8(ID))
	if err != nil {
		response := newResponse(Error, errorPaymentMethodDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}
