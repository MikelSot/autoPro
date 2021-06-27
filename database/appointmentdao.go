package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
	"time"
)

type IAppointmentCRUD interface {
	Create(appointment *model.Appointment) error
	Update(ID uint, appointment *model.Appointment) error
	GetAll(max int) (*model.Appointments, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

type IQueryAppointment interface {
	AllOrderAttentionAvailable() (map[int]string, error)
	AllAppointmentClient(ID uint, max int) (*model.Appointments, error)
}


func (a *AppointmentDao) Create(appointment *model.Appointment) error {
	const TimeError = "horario de cita no aceptada"
	if appointment.DateHour.Hour() > maxHour || appointment.DateHour.Hour() < minHour{
		return errors.New(TimeError)
	}

	DB().Create(&appointment)
	return nil
}

func (a *AppointmentDao) Update(ID uint, appointment *model.Appointment) error {
	const TimeError = "horario de cita no aceptada"
	if appointment.DateHour.Hour() > maxHour || appointment.DateHour.Hour() < minHour{
		return errors.New(TimeError)
	}

	appointmentID := model.Appointment{}
	appointmentID.ID = ID
	DB().Model(&appointmentID).Updates(appointment)
	return nil
}

func (a *AppointmentDao) GetAll(max int) (*model.Appointments, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	appointments := model.Appointments{}
	DB().Limit(max).Find(&appointments)
	return &appointments, nil
}

func (a *AppointmentDao) DeleteSoft(ID uint) error {
	appointment := model.Appointment{}
	appointment.ID = ID
	DB().Delete(&appointment)
	return nil
}

func (a *AppointmentDao) DeletePermanent(ID uint) error {
	appointment := model.Appointment{}
	appointment.ID = ID
	DB().Unscoped().Delete(&appointment)
	return nil
}

func (a *AppointmentDao) AllOrderAttentionAvailable() (map[int]string, error) {
	ordersBusy := make(map[string]string)
	now := time.Now().Format("2006-01-02")
	appointments := model.Appointments{}
	DB().Select("OrderAttention").Find(&appointments, "date(updated_at) = ?", now)

	for _, appointment := range appointments {
		ordersBusy[appointment.OrderAttention] = appointment.OrderAttention
	}

	available:= signValid(ordersBusy)
	return available, nil
}


func (a *AppointmentDao) AllAppointmentClient(ID uint, max int) (*model.Appointments, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	appointments := model.Appointments{}
	DB().Limit(max).Select(
		"Workshop",
		"Service",
		"Description",
		"DateHour",
		"OrderAttention",
		"VehicleType",
		"PickUp",
	).Find(&appointments, "client_id = ?", ID)

	return &appointments, nil
}

