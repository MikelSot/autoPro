package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
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

func (b *blogHd) Create(e echo.Context) error {
	data := model.Blog{}
	err := e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err, bool := areDataValidBlog(&data, e); !bool {
		return err
	}

	UploadBlog(&data, e)
	err = b.crudQuery.Create(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, blogCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (b *blogHd) Update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Blog{}
	err = e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	if err, bool := areDataValidBlog(&data, e); !bool {
		return err
	}
	UploadBlog(&data, e)
	err = b.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := NewResponse(Error, errorStructBlog, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, updatedBlog, nil)
	return e.JSON(http.StatusOK, resp)
}

func (b *blogHd) GetById(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data, err := b.crudQuery.GetByID(uint(ID))
	if err != nil {
		res := NewResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) GetAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data, err := b.crudQuery.GetAll(max)
	if err != nil {
		res := NewResponse(Error, errorGetAllBlog, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) DeleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = b.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		res := NewResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok,nil)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) AllBlogCategory(e echo.Context) error {
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

	data, err := b.crudQuery.AllBlogCategory(ID, max)
	if err != nil {
		res := NewResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}

func (b *blogHd) AllBlogEmployee(e echo.Context) error {
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

	data, err := b.crudQuery.AllBlogEmployee(uint(ID), max)
	if err != nil {
		res := NewResponse(Error, errorBlogDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := NewResponse(Message, ok,data)
	return e.JSON(http.StatusOK, res)
}

func UploadBlog(data *model.Blog, e echo.Context) error  {
	file, err := e.FormFile("file-blog")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	archive := "uploads/blogs/" +"@"+file.Filename
	// ruta destino
	dst, err := os.Create(archive)
	if err != nil {
		return err
	}
	defer dst.Close()
	// Copia ruta archivo a ruta destino
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	data.Pictures = archive
	return nil
}

func areDataValidBlog(data *model.Blog, e echo.Context) (error, bool) {
	data.Author = strings.TrimSpace(data.Author)
	data.Tittle = strings.TrimSpace(data.Tittle)
	data.Synthesis = strings.TrimSpace(data.Synthesis)
	data.Content = strings.TrimSpace(data.Content)

	if !isEmpty(data.Author) || !isEmpty(data.Tittle) || !isEmpty(data.Synthesis) {
		resp := NewResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp), false
	}

	if len(data.Content) < minLenContent {
		resp := NewResponse(Error, errorContentBlog, nil)
		return e.JSON(http.StatusBadRequest, resp), false
	}
	return nil, true
}
