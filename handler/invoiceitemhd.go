package handler

import (
	"github.com/MikelSot/autoPro/database"
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
	quantityError = "Su pedido excede el stock del producto"
)

type invoiceItemHd struct {
	crudQuery IInvoiceItemCRUDQuery
}

func NewInvoiceItemHd(cq IInvoiceItemCRUDQuery) invoiceItemHd {
	return invoiceItemHd{cq}
}

func (i *invoiceItemHd) Create(e echo.Context) error {
	data := model.InvoiceItem{}
	err := e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err,bool:=updateStock(data, *i,e) ; !bool{
		return err
	}
	err = i.crudQuery.Create(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	// corregir stock de productos cuando se elija uno
	resp := NewResponse(Message, invoiceItemCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (i *invoiceItemHd) Update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.InvoiceItem{}
	err = e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err,bool:=updateStockTwo(uint(ID), data, *i,e) ; !bool{
		return err
	}

	err = i.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := NewResponse(Error, errorStructInvoiceItem, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	// corregir stock de productos cuando sea actualizado y si es menor o mayor corregir
	resp := NewResponse(Message, updatedInvoiceItem, nil)
	return e.JSON(http.StatusOK, resp)
}

func (i *invoiceItemHd) DeleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = i.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		res := NewResponse(Error, errorInvoiceItemDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok,nil)
	return e.JSON(http.StatusOK, res)
}

func (i *invoiceItemHd) AllInvoiceItemInvoice(e echo.Context) error {
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

	data ,err := i.crudQuery.AllInvoiceItemInvoice(ID, max)
	if err != nil {
		res := NewResponse(Error, errorInvoiceItemDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}

func updateStock(data model.InvoiceItem,i invoiceItemHd ,e echo.Context) (error, bool) {
	product := database.NewProductDao()
	dataProduct,_:= product.GetByID(data.ProductID)
	switch  {
	case dataProduct.Stock == data.Quantity:
		stock := dataProduct.Stock - data.Quantity
		i.crudQuery.UpdateStock(data.ProductID, stock)
		return nil, true
	case dataProduct.Stock > data.Quantity:
		stock := dataProduct.Stock - data.Quantity
		i.crudQuery.UpdateStock(data.ProductID, stock)
		return nil, true
	case dataProduct.Stock < data.Quantity:
		stock := data.Quantity - dataProduct.Stock
		res := NewResponse(Error, quantityError, stock)
		return e.JSON(http.StatusBadRequest, res), false
	default:
		res := NewResponse(Error, quantityError, nil)
		return e.JSON(http.StatusBadRequest, res), false
	}
	return nil, true
}

func updateStockTwo(ID uint, data model.InvoiceItem,i invoiceItemHd ,e echo.Context) (error, bool) {
	product := database.NewProductDao()
	dataProduct,_:= product.GetByID(data.ProductID)
	itemPrevious, _:= i.crudQuery.GetByID(ID)
	switch {
	case itemPrevious.Quantity == data.Quantity:
		return nil, true
	case itemPrevious.Quantity > data.Quantity:
		stock := dataProduct.Stock + (itemPrevious.Quantity-data.Quantity)
		i.crudQuery.UpdateStock(data.ID, stock)
		return nil, true
	case itemPrevious.Quantity < data.Quantity:
		if dataProduct.Stock < (data.Quantity-itemPrevious.Quantity){
			res := NewResponse(Error, quantityError, nil)
			return e.JSON(http.StatusBadRequest, res), false
		}
		stock := dataProduct.Stock - (itemPrevious.Quantity-data.Quantity)
		i.crudQuery.UpdateStock(data.ProductID, stock)
		return nil, true
	default:
		res := NewResponse(Error, quantityError, nil)
		return e.JSON(http.StatusBadRequest, res), false
	}
}