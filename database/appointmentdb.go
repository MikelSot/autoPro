package database

import (
	"github.com/MikelSot/autoPro/model"
	"strconv"
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