package handler

import (
	"github.com/MikelSot/autoPro/model/dto"
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

func (a *appointmentHd) Create(e echo.Context) error {
	data := dto.AppointmentCreate{}
	if err := e.Bind(&data);err != nil {
		resp := NewResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError,resp)
	}

	if err, bool := areDataValidAppointment(&data, *a, e); !bool {
		return err
	}

	err := a.crudQuery.Create(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, appointmentCreated, nil)
	return e.JSON(http.StatusCreated, resp)
}

func (a *appointmentHd) Update(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := dto.AppointmentUpdate{}
	err = e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	dataValidUpdate := isDataValidUpdate(&data)
	if err, bool := areDataValidAppointment(&dataValidUpdate, *a, e); !bool {
		return err
	}

	err = a.crudQuery.Update(uint(ID), &data)
	if err != nil {
		resp := NewResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, updatedAppointment, nil)
	return e.JSON(http.StatusOK, resp)
}


func (a *appointmentHd) UpdateState(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	data := dto.AppointmentUpdateState{}
	err = e.Bind(&data)
	if err != nil {
		resp := NewResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	err = a.crudQuery.UpdateState(uint(ID), &data)
	if err != nil {
		resp := NewResponse(Error, errorStructAppointment, nil)
		return e.JSON(http.StatusInternalServerError, resp)
	}

	resp := NewResponse(Message, updatedAppointment, nil)
	return e.JSON(http.StatusOK, resp)
}


func (a *appointmentHd) GetAll(e echo.Context) error {
	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		response := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	data, err := a.crudQuery.GetAll(max)
	if err != nil {
		response := NewResponse(Error, errorGetAllAppointment, nil)
		return e.JSON(http.StatusInternalServerError, response)
	}

	res := NewResponse(Message, ok, data)
	return e.JSON(http.StatusOK, res)
}

func (a *appointmentHd) DeleteSoft(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = a.crudQuery.DeleteSoft(uint(ID))
	if err != nil {
		response := NewResponse(Error, errorAppointmentIDDoesNotExist, nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	res := NewResponse(Message, ok, nil)
	return e.JSON(http.StatusOK, res)
}

func (a *appointmentHd) AllOrderAttentionAvailable(e echo.Context) error {
	available, err := a.crudQuery.AllOrderAttentionAvailable()
	if err != nil {
		response := NewResponse(Error, attentionOrderError, nil)
		return e.JSON(http.StatusBadRequest, response)
	}
	res := NewResponse(Message, ok, available)
	return e.JSON(http.StatusOK, res)
}

func (a *appointmentHd) AllAppointmentClient(e echo.Context) error {
	ID, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		res := NewResponse(Error, errorId, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	max, err := strconv.Atoi(e.Param("max"))
	if err != nil {
		res := NewResponse(Error, errorGetAllAppointment, nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	appointments, err := a.crudQuery.AllAppointmentClient(uint(ID), max)
	if err != nil {
		response := NewResponse(Error, errorAppointmentIDDoesNotExist, nil)
		return e.JSON(http.StatusBadRequest, response)
	}

	res := NewResponse(Message, ok, appointments)
	return e.JSON(http.StatusOK, res)
}

func areDataValidAppointment(data *dto.AppointmentCreate, a appointmentHd, e echo.Context) (error, bool) {
	data.Workshop = strings.TrimSpace(data.Workshop)
	data.Service = strings.TrimSpace(data.Service)
	data.Description = strings.TrimSpace(data.Description)
	data.OrderAttention = strings.TrimSpace(data.OrderAttention)
	data.VehicleType = strings.TrimSpace(data.VehicleType)

	if !isEmpty(data.Workshop) || !isEmpty(data.Service) || !isNumber(data.OrderAttention) || !isStringVehicle(data.VehicleType) {
		resp := NewResponse(Error, errorContent, nil)
		return e.JSON(http.StatusBadRequest, resp), false
	}

	if exists, _ := a.crudQuery.QueryServiceExists(data.Service); !exists {
		resp := NewResponse(Error, errorServiceDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, resp), false
	}

	if exists, _ := a.crudQuery.QueryWorkshopExists(data.Workshop); !exists {
		resp := NewResponse(Error, errorWorkshopDoesNotExists, nil)
		return e.JSON(http.StatusBadRequest, resp), false
	}

	if data.DateHour.Hour() > maxHour || data.DateHour.Hour() < minHour {
		resp := NewResponse(Error, errorHour, nil)
		return e.JSON(http.StatusBadRequest, resp), false
	}

	if !isDateValid(data.DateHour) {
		resp := NewResponse(Error, errorDate, nil)
		return e.JSON(http.StatusBadRequest, resp), false
	}
	return nil, true
}

func isDataValidUpdate(data *dto.AppointmentUpdate) dto.AppointmentCreate {
	 dataValid := dto.AppointmentCreate{
		Workshop     :data.Workshop,
		Service      :data.Service,
		Description  :data.Description,
		DateHour     :data.DateHour,
		OrderAttention:data.OrderAttention ,
		VehicleType  : data.VehicleType,
		PickUp       : data.PickUp,
		ClientID     : data.ClientID,
	}
	return dataValid
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
