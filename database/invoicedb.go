package database

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
)

type InvoiceDao struct {
	invoiceDao model.Invoice
}

func NewInvoiceDao() InvoiceDao {
	return InvoiceDao{}
}

func (i InvoiceDao) Create(invoice *model.Invoice) error {
	DB().Create(&invoice)
	return nil
}

func (i InvoiceDao) Update(ID uint, invoice *model.Invoice) error {
	invoiceID := model.Invoice{}
	invoiceID.ID = ID
	DB().Model(&invoiceID).Updates(invoice)
	return nil
}

func (i InvoiceDao) GetByID(ID uint) (*model.Invoice, error) {
	invoice :=model.Invoice{}
	DB().First(&invoice, ID)
	return &invoice, nil
}

func (i InvoiceDao) DeleteSoft(ID uint) error {
	invoice := model.Invoice{}
	invoice.ID = ID
	DB().Delete(&invoice)
	return nil
}

func (i InvoiceDao) AllInvoiceClient(ID, max int) (*dto.InvoiceClients, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}

	invoices := dto.InvoiceClients{}
	DB().Table("invoices").Limit(max).Select(
		"Ruc",
		"Status",
		"InvoiceDate",
	).Find("client_id = ?", ID).Scan(&invoices)

	return &invoices, nil
}

func (i InvoiceDao) AllInvoiceWorkshop(ID uint, max int) (*dto.InvoiceWorkshops, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}

	invoices := dto.InvoiceWorkshops{}
	DB().Table("invoices").Limit(max).Select(
		"Ruc",
		"Status",
		"InvoiceDate",
		"EmployeeID",
		"PaymentMethodID",
	).Find( "workshop_id = ?", ID).Scan(&invoices)

	//invoices := model.Invoices{}
	//DB().Limit(max).Select(
	//	"Ruc",
	//	"Status",
	//	"InvoiceDate",
	//	"EmployeeID",
	//	"PaymentMethodID",
	//).Find(&invoices, "workshop_id = ?", ID)
	return &invoices, nil
}
