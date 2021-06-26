package database

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
)

type IProductCRUD interface {
	Create( product *model.Product) error
	Update(ID uint,  product *model.Product) error
	GetByID(ID uint) (*model.Product, error)
	GetAll(max int) (*model.Products, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

type IQueryProduct interface {
	AllProductsCategory(ID uint, max int) (dto.ProductClients, error)
	AllProductsWorkshop(ID uint, max int) (dto.ProductClients, error)
}

type ProductDao struct {
	productDao model.Product
}

func NewProductDao() ProductDao {
	return ProductDao{}
}

func (p *ProductDao) Create(product *model.Product) error {
	DB().Create(&product)
	return nil
}

func (p *ProductDao) Update(ID uint, product *model.Product) error {
	productID := model.Product{}
	productID.ID = ID
	DB().Model(&productID).Updates(product)
	return nil
}

func (p *ProductDao) GetByID(ID uint) (*model.Product, error) {
	product :=model.Product{}
	DB().First(&product, ID)
	return &product, nil
}

func (p *ProductDao) GetAll(max int) (*model.Products, error) {
	if  max < MaxGetAll {
		max = MaxGetAll
	}
	products := model.Products{}
	DB().Limit(max).First(&products)
	return &products, nil
}

func (p *ProductDao) DeleteSoft(ID uint) error {
	product := model.Product{}
	product.ID = ID
	DB().Delete(&product)
	return nil
}

func (p *ProductDao) DeletePermanent(ID uint) error {
	product := model.Product{}
	product.ID = ID
	DB().Unscoped().Delete(&product)
	return nil
}

func (p *ProductDao) AllProductsCategory(ID uint, max int) (dto.ProductClients, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}

	products := dto.ProductClients{}
	DB().Table("products").Limit(max).Find( "category_id = ?", ID).Scan(&products)
	return products, nil
}

func (p *ProductDao) AllProductsWorkshop(ID uint, max int) (dto.ProductClients, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	products := dto.ProductClients{}
	DB().Table("products").Limit(max).Find( "workshop_id = ?", ID).Scan(&products)
	return products, nil
}
