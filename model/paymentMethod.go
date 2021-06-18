package model

type PaymentMethod struct {
	FirstDataModel
	Invoices []Invoice
	TimeModel
}
