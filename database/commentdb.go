package database

import "github.com/MikelSot/autoPro/model"

type ICommentCRUD interface {
	Create(comment *model.Comment) error
	Update(ID uint,comment *model.Comment) error
	DeleteSoft(ID uint) error
}


type CommentDao struct {
	commentDao model.Comment
}

func NewCommentDao() ClientDao {
	return ClientDao{}
}

func (c CommentDao) Create(comment *model.Comment) error {
	DB().Create(&comment)
	return nil
}

func (c CommentDao) Update(ID uint, comment *model.Comment) error {
	commentID := model.Comment{}
	commentID.ID = ID
	DB().Model(&commentID).Updates(comment)
	return nil
}

func (c CommentDao) DeleteSoft(ID uint) error {
	comment := model.Comment{}
	comment.ID = ID
	DB().Delete(&comment)
	return nil
}
