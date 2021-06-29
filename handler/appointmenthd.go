package handler

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	formatDate                     = "2006-01-02"
	maxHour                        = 19
	minHour                        = 8
	regexNumber                    = `([0-9])`
	regexStringVehicle             = `([a-zA-z0-9])`
	appointmentCreated             = "Cita creado correctamente"
	updatedAppointment             = "Cita actualizada correctamente"
	errorStructAppointment         = "Hubo un error con el registro: Registro no permitido a este dominio de Email o la estructura de la cita es inválida"
	errorServiceDoesNotExists      = "No existe el servicio seleccionado"
	errorWorkshopDoesNotExists     = "No existe el taller seleccionado"
	errorHour                      = "La hora seleccionada esta fuera de servicio"
	errorDate                      = "La fecha seleccionada esta fuera de servicio" // fecha pasada
	errorGetAllAppointment         = "Hubo un problema al obtener todas las citas"
	errorAppointmentIDDoesNotExist = "El ID de la cita no existe"
	attentionOrderError            = "El ID de la cita no existe"
)

type appointmentHd struct {
	crudQuery IAppointmentCRUDQuery
}

func NewAppointmentHd(cq IAppointmentCRUDQuery) appointmentHd {
	return appointmentHd{cq}
}

func (a *appointmentHd) create(e echo.Context) error {
	data := model.Appointment{}
	err := e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = areDataValidAppointment(&data, *a, e)
	if err != nil {
		return err
	}

	err = a.crudQuery.Create(&data)
	if err != nil {
		resp := newResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, appointmentCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (a *appointmentHd) update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := model.Appointment{}
	err = e.Bind(&data)
	if err != nil {
		resp := newResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = areDataValidAppointment(&data, *a, e)
	if err != nil {
		return err
	}

	err = a.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := newResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := newResponse(Message, updatedAppointment, nil)
	return e.JSON(http.StatusOK, resp)
}

//func (a *appointmentHd) getById(e echo.Context) error {
//	ID, err := strconv.Atoi(e.Param("id"))
//	if err != nil {
//		res := newResponse(Error, errorId, nil)
//		return e.JSON(http.StatusBadRequest, res)
//	}
//}

func (a *appointmentHd) getAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := a.crudQuery.GetAll(max)
	if err != nil {
		response := newResponse(Error, errorGetAllAppointment, nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	res := newResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (a *appointmentHd) deleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = a.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		response := newResponse(Error, errorAppointmentIDDoesNotExist, nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	res := newResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}

func (a *appointmentHd) allOrderAttentionAvailable(e echo.Context) error {
	available, err := a.crudQuery.AllOrderAttentionAvailable()
	if err != nil {
		response := newResponse(Error, attentionOrderError, nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	res := newResponse(Message, ok, available)
	return e.JSON(http.StatusOK, res)
}

func (a *appointmentHd) allAppointmentClient(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := newResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		res := newResponse(Error, errorGetAllAppointment, nil)
		return e.JSON(http.StatusBadRequest, res)
	}


	appointments, err := a.crudQuery.AllAppointmentClient(uint(ID), max)
	if err != nil {
		response := newResponse(Error, errorAppointmentIDDoesNotExist, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := newResponse(Message, ok, appointments)
	return e.JSON(http.StatusOK, res)
}

func areDataValidAppointment(data *model.Appointment, a appointmentHd, e echo.Context) error {
	data.Workshop = strings.TrimSpace(data.Workshop)
	data.Service = strings.TrimSpace(data.Service)
	data.Description = strings.TrimSpace(data.Description)
	data.OrderAttention = strings.TrimSpace(data.OrderAttention)
	data.VehicleType = strings.TrimSpace(data.VehicleType)

	if !isEmpty(data.Workshop) || !isEmpty(data.Service) || !isNumber(data.OrderAttention) || !isStringVehicle(data.VehicleType) {
		resp := newResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	if exists, _ := a.crudQuery.QueryServiceExists(data.Service); !exists {
		resp := newResponse(Error, errorServiceDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	if exists, _ := a.crudQuery.QueryWorkshopExists(data.Workshop); !exists {
		resp := newResponse(Error, errorWorkshopDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	if data.DateHour.Hour() > maxHour || data.DateHour.Hour() < minHour {
		resp := newResponse(Error, errorHour, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}

	if !isDateValid(data.DateHour) {
		resp := newResponse(Error, errorDate, nil)
		return e.JSON(http.StatusBadRequest, resp)
	}
	return nil
}

func isNumber(num string) bool {
	r, _ := regexp.Compile(regexNumber)
	if !r.MatchString(num) {
		return false
	}
	return true
}

func isStringVehicle(vehicle string) bool {
	r, _ := regexp.Compile(regexStringVehicle)
	if !r.MatchString(vehicle) {
		return false
	}
	return true
}

func isDateValid(date time.Time) bool {
	now := time.Now().Format(formatDate)
	dateString := date.Format(formatDate)

	if dateString < now {
		return false
	}
	return true
}
