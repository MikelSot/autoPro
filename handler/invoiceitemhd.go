package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

const (
	errorStructInvoiceItem = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura del item de factura es inválida"
	invoiceItemCreated     = "item de factura creado correctamente"
	updatedInvoiceItem = "item de factura actualizada correctamente"
	errorInvoiceItemDoesNotExists = "No existe el item de factura seleccionado"
)

type invoiceItemHd struct {
	crudQuery IInvoiceItemCRUDQuery
}

func NewInvoiceItemHd(cq IInvoiceItemCRUDQuery) invoiceItemHd {
	return invoiceItemHd{cq}
}

func (i *invoiceItemHd) create(e echo.Context) error {
	data := model.InvoiceItem{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = i.crudQuery.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	// corregir stock de productos cuando se elija uno

	resp := newResponse(Message, invoiceItemCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (i *invoiceItemHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.InvoiceItem{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = i.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	// corregir stock de productos cuando sea actualizado y si es menor o mayor corregir

	resp := newResponse(Message, updatedInvoiceItem, nil)
	return e.JSON(http.StatusOK, resp)
}

func (i *invoiceItemHd) deleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = i.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		res := newResponse(Error, errorInvoiceItemDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, ok,nil)
	return e.JSON(http.StatusOK, res)
}

func (i *invoiceItemHd) allInvoiceItemInvoice(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data ,err := i.crudQuery.AllInvoiceItemInvoice(ID, max)
	if err != nil {
		res := newResponse(Error, errorInvoiceItemDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}
