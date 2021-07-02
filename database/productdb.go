package database

import (
	"github.com/MikelSot/autoPro/model"
	"github.com/MikelSot/autoPro/model/dto"
	"regexp"
	"strings"
)


type ProductDao struct {
	productDao model.Product
}

func NewProductDao() ProductDao {
	return ProductDao{}
}

func (p *ProductDao) Create(product *model.Product) error {
	regexSpace := regexp.MustCompile(` `)
	nameWithoutSpace := regexSpace.ReplaceAllString(product.Name, "")
	url := at + strings.ToLower(nameWithoutSpace)
	product.Uri = url
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
	DB().Limit(max).Find(&products)
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