package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

const (
	errorStructWorkshop = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura del taller es inválida"
	workshopCreated  = "Taller creado correctamente"
	errorGetAllWorkshop = "Hubo un problema al obtener todos los talleres"
)

type workShopHd struct {
	crud IWorkshopCRUD
}

func NewWorkshopHd(c IWorkshopCRUD) workShopHd {
	return workShopHd{c}
}

func (w *workShopHd) Create(e echo.Context) error {
	data := model.Workshop{}
	err := e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructWorkshop, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	data.Name = strings.TrimSpace(data.Name)
	if !isEmpty(data.Name) {
		resp := NewResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = w.crud.Create(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructWorkshop, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, workshopCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}


func (w *workShopHd) GetAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := w.crud.GetAll(max)
	if err != nil {
		response := NewResponse(Error, errorGetAllWorkshop, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}