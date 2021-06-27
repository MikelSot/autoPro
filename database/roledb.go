package database

import "github.com/MikelSot/autoPro/model"


type RoleDao struct {
	roleDao model.Role
}

func NewRoleDao() RoleDao {
	return RoleDao{}
}


func (r *RoleDao) Create(role *model.Role) error {
	DB().Create(&role)
	return nil
}

func (r *RoleDao) Update(ID uint8, role *model.Role) error {
	roleID := model.Role{}
	roleID.ID = ID
	DB().Model(&roleID).Updates(role)
	return nil
}

func (r *RoleDao) GetByID(ID uint8) (*model.Role, error) {
	role :=model.Role{}
	DB().First(&role, ID)
	return &role, nil
}

func (r *RoleDao) GetAll(max int) (*model.Roles, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	roles := model.Roles{}
	DB().Limit(max).First(&roles)
	return &roles, nil
}

func (r *RoleDao) DeleteSoft(ID uint8) error {
	role := model.Role{}
	role.ID = ID
	DB().Delete(&role)
	return nil
}

func (r *RoleDao) DeletePermanent(ID uint8) error {
	role := model.Role{}
	role.ID = ID
	DB().Unscoped().Delete(&role)
	return nil
}
