package product

import "gorm.io/gorm"

type IRepository interface {
	CreateProduct(product *Product) (*Product, error)
	GetProduct() ([]*Product, error)
	UpdateProduct(product *Product) (*Product, error)
	DeleteProduct(id int) error
	GetProductById(id int) (*Product, error)
	UpdateItemProduct(id, stock int) error
}

type repository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IRepository {
	return &repository{db: db}
}

func (r repository) CreateProduct(product *Product) (*Product, error) {
	if err := r.db.Debug().Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r repository) GetProduct() ([]*Product, error) {
	products := make([]*Product, 0)

	if err := r.db.Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (r repository) GetProductById(id int) (*Product, error) {
	var product *Product

	if err := r.db.Debug().Where("id = ?", id).Take(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (r repository) UpdateProduct(product *Product) (*Product, error) {
	_, err := r.GetProductById(product.ID)
	if err != nil {
		return product, err
	}

	// !fixMe: logic error update product
	err = r.db.Debug().Model(&Product{}).Where("id = ?", product.ID).Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r repository) UpdateItemProduct(id, stock int) error {
	err := r.db.Debug().Model(&Product{}).Where("id = ?", id).UpdateColumn("stock", stock).Error
	if err != nil {
		return err
	}

	return err
}

func (r repository) DeleteProduct(id int) error {
	err := r.db.Debug().Delete(&Product{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
