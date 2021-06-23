package database

import "github.com/MikelSot/autoPro/model"

type IPaymentMethodCRUD interface {
	Create(method *model.PaymentMethod) error
	Update(ID uint8, method *model.PaymentMethod) error
	GetByID(ID uint8) (*model.PaymentMethod, error)
	GetAll(max int) (*model.PaymentMethods, error)
	DeleteSoft(ID uint8) error
	DeletePermanent(ID uint8) error
}

type PaymentMethodDao struct {
	paymentMethodDao model.PaymentMethod
}

func NewPaymentMethodDao() PaymentMethodDao {
	return PaymentMethodDao{}
}

func (p PaymentMethodDao) Create(method *model.PaymentMethod) error {
	DB().Create(&method)
	return nil
}

func (p PaymentMethodDao) Update(ID uint8, method *model.PaymentMethod) error {
	methodID := model.PaymentMethod{}
	methodID.ID = ID
	DB().Model(&methodID).Updates(method)
	return nil
}

func (p PaymentMethodDao) GetByID(ID uint8) (*model.PaymentMethod, error) {
	method :=model.PaymentMethod{}
	DB().First(&method, ID)
	return &method, nil
}

func (p PaymentMethodDao) GetAll(max int) (*model.PaymentMethods, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	methods := model.PaymentMethods{}
	DB().Limit(max).First(&methods)
	return &methods, nil
}

func (p PaymentMethodDao) DeleteSoft(ID uint8) error {
	method := model.PaymentMethod{}
	method.ID = ID
	DB().Delete(&method)
	return nil
}

func (p PaymentMethodDao) DeletePermanent(ID uint8) error {
	method := model.PaymentMethod{}
	method.ID = ID
	DB().Unscoped().Delete(&method)
	return nil
}


