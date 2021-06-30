package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

const (
	minLenComment             = 10
	errorStructComment        = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de comentario es inválida"
	errorContentComment       = "El contenido del comentario debe tener más de 10 caracteres"
	commentCreated            = "Comentario creado correctamente"
	updatedComment            = "Comentario actualizada correctamente"
	errorCommentDoesNotExists = "No existe el comentario seleccionado"
)

type commentHd struct {
	crudQuery ICommentCRUDQuery
}

func NewCommentHd(cq ICommentCRUDQuery) commentHd {
	return commentHd{cq}
}

func (c *commentHd) Create(e echo.Context) error {
	data := model.Comment{}
	err := e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructComment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if len(data.Comment) < minLenComment {
		resp := NewResponse(Error, errorContentComment, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = c.crudQuery.Create(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructComment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, commentCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (c *commentHd) Update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Comment{}
	err = e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructComment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if len(data.Comment) < minLenComment {
		resp := NewResponse(Error, errorContentComment, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = c.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := NewResponse(Error, errorStructComment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, updatedComment, nil)
	return e.JSON(http.StatusOK, resp)
}

func (c *commentHd) DeleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		resp := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	err = c.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		resp := NewResponse(Error, errorCommentDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	res := NewResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}

func (c *commentHd) AllCommentBlog(e echo.Context) error {
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

	data, err := c.crudQuery.AllCommentBlog(ID, max)
	if err != nil {
		res := NewResponse(Error, errorCommentDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}

func (c *commentHd) AllCommentProduct(e echo.Context) error {
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

	data, err := c.crudQuery.AllCommentProduct(ID, max)
	if err != nil {
		res := NewResponse(Error, errorCommentDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}
