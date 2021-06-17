package model

type Role struct {
	FirstDataModel
	Employees []Employee
	Clients   []Client
	TimeModel
}
