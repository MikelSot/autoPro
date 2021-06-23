package database

import "github.com/MikelSot/autoPro/model"

type IBlogCRUD interface {
	Create( blog *model.Blog) error
	Update(ID uint,  blog *model.Blog) error
	GetByID(ID uint) (*model.Blog, error)
	GetAll(max int) (*model.Blogs, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

type IQueryBlog interface {
	AllBlogCategory(ID, max int) (*model.Blogs, error)
	AllBlogEmployee(ID uint, max int) (*model.Blogs, error)
}

type BlogDao struct {
	blogDao model.Blog
}

func NewBloDao() BlogDao {
	return BlogDao{}
}

func (b BlogDao) Create(blog *model.Blog) error {
	DB().Create(&blog)
	return nil
}

func (b BlogDao) Update(ID uint, blog *model.Blog) error {
	blogID := model.Blog{}
	blogID.ID = ID
	DB().Model(&blogID).Updates(blog)
	return nil
}

func (b BlogDao) GetByID(ID uint) (*model.Blog, error) {
	blog :=model.Blog{}
	DB().First(&blog, ID)
	return &blog, nil
}

func (b BlogDao) GetAll(max int) (*model.Blogs, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	blogs := model.Blogs{}
	DB().Limit(max).First(&blogs)
	return &blogs, nil
}

func (b BlogDao) DeleteSoft(ID uint) error {
	blog := model.Blog{}
	blog.ID = ID
	DB().Delete(&blog)
	return nil
}

func (b BlogDao) DeletePermanent(ID uint) error {
	blog := model.Blog{}
	blog.ID = ID
	DB().Unscoped().Delete(&blog)
	return nil
}

func (b BlogDao) AllBlogCategory(ID, max int) (*model.Blogs, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	blogs := model.Blogs{}
	DB().Limit(max).Find(&blogs, "category_id = ?", ID)
	return &blogs, nil
}

func (b BlogDao) AllBlogEmployee(ID uint, max int) (*model.Blogs, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	blogs := model.Blogs{}
	DB().Limit(max).Find(&blogs, "employee_id = ?", ID)
	return &blogs, nil
}
