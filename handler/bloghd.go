package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

const (
	minLenContent          = 200
	errorStructBlog        = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura del blog es inválida"
	blogCreated            = "Blog creado correctamente"
	errorContentBlog       = "El contenido del blog debe tener más de 200 caracteres"
	updatedBlog            = "Blog actualizada correctamente"
	errorBlogDoesNotExists = "No existe el Blog seleccionado"
	errorGetAllBlog        = "Hubo un problema al obtener todos los Blogs"
)

type blogHd struct {
	crudQuery IBlogCRUDQuery
}

func NewBlogHd(cq IBlogCRUDQuery) blogHd {
	return blogHd{cq}
}

func (b *blogHd) create(e echo.Context) error {
	data := model.Blog{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err = areDataValidBlog(&data, e); err != nil {
		return err
	}

	err = b.crudQuery.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, blogCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (b *blogHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Blog{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err = areDataValidBlog(&data, e); err != nil {
		return err
	}

	err = b.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedBlog, nil)
	return e.JSON(http.StatusOK, resp)
}

func (b *blogHd) getById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data, err := b.crudQuery.GetByID(uint(ID))
	if err != nil {
		res := newResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) getAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data, err := b.crudQuery.GetAll(max)
	if err != nil {
		res := newResponse(Error, errorGetAllBlog, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) deleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = b.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		res := newResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, ok,nil)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) allBlogCategory(e echo.Context) error {
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

	data, err := b.crudQuery.AllBlogCategory(ID, max)
	if err != nil {
		res := newResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) allBlogEmployee(e echo.Context) error {
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

	data, err := b.crudQuery.AllBlogEmployee(uint(ID), max)
	if err != nil {
		res := newResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}

func areDataValidBlog(data *model.Blog, e echo.Context) error {
	data.Author = strings.TrimSpace(data.Author)
	data.Tittle = strings.TrimSpace(data.Tittle)
	data.Synthesis = strings.TrimSpace(data.Synthesis)
	data.Content = strings.TrimSpace(data.Content)

	if !isEmpty(data.Author) || !isEmpty(data.Tittle) || !isEmpty(data.Synthesis) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	if len(data.Content) < minLenContent {
		resp := newResponse(Error, errorContentBlog, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}
	return nil
}
