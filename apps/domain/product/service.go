package product

import (
	"errors"

	"github.com/hafizdkn/toko-belanja/apps/domain/category"
)

type IService interface {
	CreateProduct(input *ProductCreateInput) (*Product, error)
	UpdateProduct(updateId int, input *ProductUpdateInput) (*Product, error)
	DeleteProduct(id int) error
	GetProduct() ([]*Product, error)
	GetProductById(id int) (*Product, error)
	UpdateItemProduct(id, quantity int) error
}

type service struct {
	repo            IRepository
	serviceCategory category.IService
}

func NewServiceProduct(repo IRepository, serviceCategory category.IService) IService {
	return &service{repo: repo, serviceCategory: serviceCategory}
}

func (s *service) UpdateItemProduct(id, quantity int) error {
	product, err := s.repo.GetProductById(id)
	stock := product.Stock

	if err != nil {
		return err
	}

	if stock == 0 || quantity >= stock {
		return errors.New("Order lebih besar dariapa stock")
	}

	newStock := stock - quantity
	if newStock <= 0 {
		return errors.New("Order lebih besar dariapa stock")
	}

	err = s.repo.UpdateItemProduct(id, newStock)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetProductById(categoryId int) (*Product, error) {
	category, err := s.repo.GetProductById(categoryId)
	if err != nil {
		return category, err
	}

	return category, nil
}

func (s service) CreateProduct(input *ProductCreateInput) (*Product, error) {
	_, err := s.serviceCategory.GetCategoryById(input.CategoryId)
	if err != nil {
		return &Product{}, err
	}

	product := &Product{
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryId: input.CategoryId,
	}

	product, err = s.repo.CreateProduct(product)
	if err != nil {
		return product, err
	}

	return ProductResponseRequest(product), err
}

func (s service) UpdateProduct(productId int, input *ProductUpdateInput) (*Product, error) {
	_, err := s.serviceCategory.GetCategoryById(input.CategoryId)
	if err != nil {
		return &Product{}, err
	}

	product := &Product{
		ID:         productId,
		Title:      input.Title,
		Price:      input.Price,
		Stock:      input.Stock,
		CategoryId: input.CategoryId,
	}

	product, err = s.repo.UpdateProduct(product)
	if err != nil {
		return product, err
	}

	return ProductResponseRequest(product), nil
}

func (s service) GetProduct() ([]*Product, error) {
	products, err := s.repo.GetProduct()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s service) DeleteProduct(id int) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
