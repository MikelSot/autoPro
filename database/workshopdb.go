package database

import "github.com/MikelSot/autoPro/model"


type WorkshopDao struct {
	workshopDao model.Workshop
}

func NewWorkshopDao() WorkshopDao {
	return WorkshopDao{}
}

func (w *WorkshopDao) Create(workshop *model.Workshop) error {
	DB().Create(&workshop)
	return nil
}

func (w *WorkshopDao) Update(ID uint8, workshop *model.Workshop) error {
	workshopID := model.Workshop{}
	workshopID.ID = ID
	DB().Model(&workshopID).Updates(workshop)
	return nil
}

func (w *WorkshopDao) GetByID(ID uint8) (*model.Workshop, error) {
	workshop :=model.Workshop{}
	DB().First(&workshop, ID)
	return &workshop, nil
}

func (w *WorkshopDao) GetAll(max int) (*model.Workshops, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	workshops := model.Workshops{}
	DB().Limit(max).Find(&workshops)
	return &workshops, nil
}

func (w *WorkshopDao) DeleteSoft(ID uint8) error {
	workshop := model.Workshop{}
	workshop.ID = ID
	DB().Delete(&workshop)
	return nil
}

func (w *WorkshopDao) DeletePermanent(ID uint8) error {
	workshop := model.Workshop{}
	workshop.ID = ID
	DB().Unscoped().Delete(&workshop)
	return nil
}
