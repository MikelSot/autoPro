package database

import (
	"github.com/MikelSot/autoPro/model"
	"regexp"
	"strings"
)

type ServiceDao struct {
	serviceDao model.Service
}

func NewServiceDao() ServiceDao {
	return ServiceDao{}
}

func (s ServiceDao) Create(service *model.Service) error {
	regexSpace := regexp.MustCompile(` `)
	nameWithoutSpace := regexSpace.ReplaceAllString(service.Name, "")
	url := at + strings.ToLower(nameWithoutSpace)
	service.Uri = url
	DB().Create(&service)
	return nil
}

func (s ServiceDao) Update(ID uint8, service *model.Service) error {
	serviceID := model.Service{}
	serviceID.ID = ID
	DB().Model(&serviceID).Updates(service)
	return nil
}

func (s ServiceDao) GetByID(ID uint8) (*model.Service, error) {
	service :=model.Service{}
	DB().First(&service, ID)
	return &service, nil
}

func (s ServiceDao) GetAll(max int) (*model.Services, error) {
	if  max < MaxGetAllHome {
		max = MaxGetAllHome
	}
	services := model.Services{}
	DB().Limit(max).Find(&services)
	return &services, nil
}

func (s ServiceDao) DeleteSoft(ID uint8) error {
	service := model.Service{}
	service.ID = ID
	DB().Delete(&service)
	return nil
}

func (s ServiceDao) DeletePermanent(ID uint8) error {
	service := model.Service{}
	service.ID = ID
	DB().Unscoped().Delete(&service)
	return nil
}
