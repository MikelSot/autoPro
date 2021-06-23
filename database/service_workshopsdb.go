package database

import "github.com/MikelSot/autoPro/model"

type IServiceWorkshopCRUD interface {
	Create(workshops *model.Service_Workshops) error
	Update(SID uint8, WID uint8,workshops *model.Service_Workshops) error
	DeleteSoft(SID uint8, WID uint8) error
}

type ServiceWorkshopDao struct {
	serviceWorkshopDao model.Role
}

func NewServiceWorkshopDao() ServiceWorkshopDao {
	return ServiceWorkshopDao{}
}

// para este CRUD hacer pruebas

func (s ServiceWorkshopDao) Create(workshops *model.Service_Workshops) error {
	DB().Create(&workshops)
	return nil
}

func (s ServiceWorkshopDao) Update(SID uint8, WID uint8, workshops *model.Service_Workshops) error {
	workshopID := model.Service_Workshops{}
	workshopID.ServiceID = SID
	workshopID.WorkshopID = WID
	DB().Model(&workshopID).Updates(workshops)
	return nil
}

func (s ServiceWorkshopDao) DeleteSoft(SID uint8, WID uint8) error {
	workshops := model.Service_Workshops{}
	workshops.ServiceID = SID
	workshops.WorkshopID = WID
	DB().Delete(&workshops)
	return nil
}
