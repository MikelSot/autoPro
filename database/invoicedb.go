package database

import "github.com/MikelSot/autoPro/model"

type IInvoiceCRUD interface {
	Create(invoice *model.Invoice) error
	Update(ID uint, invoice *model.Invoice) error
	GetByID(ID uint) (*model.Invoice, error)
	DeleteSoft(ID uint) error
}

type IQueryInvoice interface {
	AllInvoiceClient(ID, max int) (*model.Invoices, error)
	AllInvoiceWorkshop(ID uint, max int) (*model.Invoices, error)
}

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

func (i InvoiceDao) AllInvoiceClient(ID, max int) (*model.Invoices, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	// incoicecliet
	invoices := model.Invoices{}
	DB().Limit(max).Select(
		"Ruc",
		"Status",
		"InvoiceDate",
	).Find(&invoices, "client_id = ?", ID)

	return &invoices, nil
}

func (i InvoiceDao) AllInvoiceWorkshop(ID uint, max int) (*model.Invoices, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	// incoiceworkshop

	invoices := model.Invoices{}
	DB().Limit(max).Select(
		"Ruc",
		"Status",
		"InvoiceDate",
		"EmployeeID",
		"PaymentMethodID",
	).Find(&invoices, "workshop_id = ?", ID)

	return &invoices, nil
}
