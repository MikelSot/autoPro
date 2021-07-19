package database

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
	"strconv"
	"time"
)

const (
	minAppointment = 1
	maxAppointment = 20
)



type AppointmentDao struct {
	appointmentDao model.Appointment
}

func NewAppointmentDao() AppointmentDao {
	return AppointmentDao{}
}


func (a *AppointmentDao) Create(dto *dto.AppointmentCreate) error {
	appointment := model.Appointment{
		Workshop     :dto.Workshop,
		Service      :dto.Service,
		Description  :dto.Description,
		DateHour     :dto.DateHour,
		OrderAttention:dto.OrderAttention ,
		VehicleType  : dto.VehicleType,
		PickUp       : dto.PickUp,
		ClientID     : dto.ClientID,
	}
	DB().Create(&appointment)
	return nil
}

func (a *AppointmentDao) Update(ID uint, dto *dto.AppointmentUpdate) error {
	appointment := model.Appointment{
		Workshop     :dto.Workshop,
		Service      :dto.Service,
		Description  :dto.Description,
		DateHour     :dto.DateHour,
		OrderAttention:dto.OrderAttention ,
		State        : dto.State,
		VehicleType  : dto.VehicleType,
		PickUp       : dto.PickUp,
		ClientID     : dto.ClientID,
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

func (a *AppointmentDao) QueryServiceExists(name string) (bool, error) {
	service := model.Service{}
	values := DB().Limit(1).Select("name").Find(&service, "name = ?", name)
	if values.RowsAffected != ZeroRowsAffected {
		return true, nil
	}
	return false, nil
}

func (a *AppointmentDao) QueryWorkshopExists(name string) (bool, error) {
	workshop := model.Workshop{}
	values := DB().Limit(1).Select("name").Find(&workshop, "name = ?", name)
	if values.RowsAffected != ZeroRowsAffected {
		return true, nil
	}
	return false, nil
}


// signValid firltra el N° que estan disponibles para la cita
func signValid(busy map[string]string)  map[int]string{
	numbers := make(map[int]string)
	for i := minAppointment; i <= maxAppointment; i++ {
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