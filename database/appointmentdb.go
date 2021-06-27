package database

import (
	"errors"
	"github.com/MikelSot/autoPro/model"
	"strconv"
	"time"
)

const (
	maxHour = 19
	minHour = 8
	maxAppointment = 20
)

var (
	minAppointment = 0
)



type AppointmentDao struct {
	appointmentDao model.Appointment
}

func NewApointmentDao() AppointmentDao {
	return AppointmentDao{}
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



// signValid firltra el N° que estan disponibles para la cita
func signValid(busy map[string]string)  map[int]string{
	numbers := make(map[int]string)
	for i := minAppointment; i <= maxAppointment; minAppointment++ {
		numbers[i] = strconv.Itoa(i)
	}

	for _, v := range busy {
		for key, _ := range numbers {
			if numbers[key] == v{
				delete(numbers, key)
				delete(busy, v)
			}
		}
	}
	return numbers
}