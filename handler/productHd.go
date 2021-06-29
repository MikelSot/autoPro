package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

const (
	updatedProduct              = "Producto actualizado correctamente"
	errorStructProduct          = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de productos es inválida"
	productCreated              = "Producto creado correctamente"
	errorProductIDDoesNotExists = "No existe el producto seleccionado"
	errorGetAllProduct          = "Hubo un problema al obtener todos los productos"
	errorGetAllProductCategory  = "Hubo un problema al obtener todos los productos por categoria"
	errorGetAllProductWorkshop  = "Hubo un problema al obtener todos los productos por taller"
)

type productHd struct {
	crudQuery IProductCRUDQuery
}

func NewProductHd(cq IProductCRUDQuery) productHd {
	return productHd{cq}
}

func (p *productHd) create(e echo.Context) error {
	data := model.Product{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStruct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err = areDataValidProduct(&data, *p, e); err != nil {
		return err
	}

	err = p.crudQuery.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructProduct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, productCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (p *productHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Product{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructProduct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err = areDataValidProduct(&data, *p, e); err != nil {
		return err
	}

	err = p.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructProduct, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedProduct, nil)
	return e.JSON(http.StatusOK, resp)
}

func (p *productHd) getById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := p.crudQuery.GetByID(uint(ID))
	if err != nil {
		response := newResponse(Error, errorProductIDDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (p *productHd) getAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := p.crudQuery.GetAll(max)
	if err != nil {
		response := newResponse(Error, errorGetAllProduct, nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (p *productHd) allProductsCategory(e echo.Context) error {
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

	data, err := p.crudQuery.AllProductsCategory(uint(ID), max)
	if err != nil {
		response := newResponse(Error, errorGetAllProductCategory, nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (p *productHd) allProductsWorkshop(e echo.Context) error {
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

	data, err := p.crudQuery.AllProductsWorkshop(uint(ID), max)
	if err != nil {
		response := newResponse(Error, errorGetAllProductWorkshop, nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func areDataValidProduct(data *model.Product, p productHd, e echo.Context) error {
	data.Name = strings.TrimSpace(data.Name)
	data.SKU = strings.TrimSpace(data.SKU)
	data.ProductCode = strings.TrimSpace(data.ProductCode)
	data.Measures = strings.TrimSpace(data.Measures)
	data.Description = strings.TrimSpace(data.Description)

	if !isEmpty(data.Name) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	return nil
}
