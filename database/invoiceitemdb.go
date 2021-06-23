package database

import "github.com/MikelSot/autoPro/model"

type IInvoiceItemCRUD interface {
	Create(item *model.InvoiceItem) error
	Update(ID uint, item *model.InvoiceItem) error
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

type IQueryInvoiceItem interface {
	AllInvoiceItemInvoice(ID, max int) (*model.InvoiceItems, error)
}

type InvoiceItemDao struct {
	invoiceItemDao model.InvoiceItem
}

func NewInvoiceItemDao() InvoiceItemDao {
	return InvoiceItemDao{}
}

func (i InvoiceItemDao) Create(item *model.InvoiceItem) error {
	DB().Create(&item)
	return nil
}

func (i InvoiceItemDao) Update(ID uint, item *model.InvoiceItem) error {
	itemID := model.InvoiceItem{}
	itemID.ID = ID
	DB().Model(&itemID).Updates(item)
	return nil
}

func (i InvoiceItemDao) DeleteSoft(ID uint) error {
	item := model.InvoiceItem{}
	item.ID = ID
	DB().Delete(&item)
	return nil
}

func (i InvoiceItemDao) DeletePermanent(ID uint) error {
	item := model.InvoiceItem{}
	item.ID = ID
	DB().Unscoped().Delete(&item)
	return nil
}

func (i InvoiceItemDao) AllInvoiceItemInvoice(ID, max int) (*model.InvoiceItems, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	items := model.InvoiceItems{}
	DB().Limit(max).Find(&items, "invoice_id = ?", ID)

	// hacer dos inner joing con producto y servicio
	return &items, nil
}
