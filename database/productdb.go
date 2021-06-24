package database

import "github.com/MikelSot/autoPro/model"

type IProductCRUD interface {
	Create( product *model.Product) error
	Update(ID uint,  product *model.Product) error
	GetByID(ID uint) (*model.Product, error)
	GetAll(max int) (*model.Products, error)
	DeleteSoft(ID uint) error
	DeletePermanent(ID uint) error
}

type IQueryProduct interface {
	AllProductsCategory(ID uint, max int) (model.Products, error)
	AllProductsWorkshop(ID uint, max int) (model.Products, error)
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

func (p *ProductDao) AllProductsCategory(ID uint, max int) (model.Products, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}

	// usar su dto para guardar info retorno
	products := model.Products{}
	DB().Limit(max).Find(&products, "category_id = ?", ID)
	return products, nil
}

func (p *ProductDao) AllProductsWorkshop(ID uint, max int) (model.Products, error) {
	if  max < MaxGetAll{
		max = MaxGetAll
	}
	products := model.Products{}
	DB().Limit(max).Find(&products, "workshop_id = ?", ID)
	return products, nil
}
