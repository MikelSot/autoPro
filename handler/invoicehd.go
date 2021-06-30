package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

const (
	errorStructInvoice = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de factura es inválida"
	invoiceCreated     = "Factura creada correctamente"
	updatedInvoice = "Factura actualizada correctamente"
)

type invoiceHd struct {
	crudQuery IInvoiceCRUDQuery
}

func NewInvoiceHd(cq IInvoiceCRUDQuery) invoiceHd {
	return invoiceHd{cq}
}

func (i *invoiceHd) Create(e echo.Context) error {
	data := model.Invoice{}
	err := e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = i.crudQuery.Create(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, invoiceCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (i *invoiceHd) Update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Invoice{}
	err = e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = i.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, updatedInvoice, nil)
	return e.JSON(http.StatusOK, resp)
}

func (i *invoiceHd) GetById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data,err := i.crudQuery.GetByID(uint(ID))
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (i *invoiceHd) DeleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = i.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	res := NewResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}

func (i *invoiceHd) AllInvoiceClient(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data ,err := i.crudQuery.AllInvoiceClient(ID, max)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (i *invoiceHd) AllInvoiceWorkshop(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data ,err := i.crudQuery.AllInvoiceWorkshop(uint(ID), max)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoice, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}
