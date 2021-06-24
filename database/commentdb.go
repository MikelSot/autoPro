package database

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
)

type ICommentCRUD interface {
	Create(comment *model.Comment) error
	Update(ID uint,comment *model.Comment) error
	DeleteSoft(ID uint) error
}

type IQueryComment interface {
	AllCommentBlog(ID, max int) (*dto.CommentClients, error) // ID blog
	AllCommentProduct(ID, max int) (*dto.CommentClients, error) // ID product
}

type CommentDao struct {
	commentDao model.Comment
}

func NewCommentDao() ClientDao {
	return ClientDao{}
}

func (c *CommentDao) Create(comment *model.Comment) error {
	DB().Create(&comment)
	return nil
}

func (c *CommentDao) Update(ID uint, comment *model.Comment) error {
	commentID := model.Comment{}
	commentID.ID = ID
	DB().Model(&commentID).Updates(comment)
	return nil
}

func (c *CommentDao) DeleteSoft(ID uint) error {
	comment := model.Comment{}
	comment.ID = ID
	DB().Delete(&comment)
	return nil
}

func (c *CommentDao) AllCommentBlog(ID, max int) (*dto.CommentClients, error) {
	if  max < Maxcomment{
		max = Maxcomment
	}

	commentClients := dto.CommentClients{}
	DB().Table("comments").Limit(max).Select(
		"c.name",
		"c.last_name",
			  "cm.comment",
			  "cm.updated_at",
		).Joins(
			"INNER JOIN clients c on c.id = cm.client_id",
		).Joins(
			"INNER JOIN blogs b on b.id = cm.blog_id",
		).Where("b.id = ?", ID).Scan(&commentClients)

	return &commentClients, nil
}

func (c *CommentDao) AllCommentProduct(ID, max int) (*dto.CommentClients, error) {
	if  max < Maxcomment{
		max = Maxcomment
	}

	commentClients := dto.CommentClients{}
	DB().Table("comments").Limit(max).Select(
		"c.name",
		"c.last_name",
		      "cm.comment",
		      "cm.updated_at",
	).Joins(
		"INNER JOIN clients c on c.id = cm.client_id",
	).Joins(
		"INNER JOIN products p on p.id = cm.product_id",
	).Where("p.id = ?", ID).Scan(&commentClients)

	return &commentClients, nil
}
