package transactionhistory

import (
	"errors"

	"github.com/hafizdkn/toko-belanja/apps/domain/product"
	"github.com/hafizdkn/toko-belanja/apps/domain/user"
)

type IService interface {
	BuyItem(userId int, input *TransactionCreateInput) (*Transaction, error)
	GetMyTransactions() ([]*Transaction, error)
	GetUserTransactions() ([]*Transaction, error)
}

type service struct {
	serviceProduct product.IService
	serviceUser    user.IService
	repo           IRepository
}

func NewTransactionService(serviceProdut product.IService, serviceUser user.IService, repo IRepository) IService {
	return &service{serviceProduct: serviceProdut, serviceUser: serviceUser, repo: repo}
}

func (s service) GetMyTransactions() ([]*Transaction, error) {
	transactions, err := s.repo.GetMyTransactions()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s service) GetUserTransactions() ([]*Transaction, error) {
	transactions, err := s.repo.GetUserTransactions()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s service) BuyItem(userId int, input *TransactionCreateInput) (*Transaction, error) {
	var transaction *Transaction

	product, err := s.CheckProduct(input.ProductId, input.Quantity)
	if err != nil {
		return TransactionResponse(transaction), err
	}

	err = s.UpdateItemProduct(product.ID, input.Quantity)
	if err != nil {
		return transaction, err
	}

	totalPrice, err := s.CheckUserBalance(userId, input.Quantity, product.Price)
	if err != nil {
		return TransactionResponse(transaction), err
	}

	transaction = &Transaction{
		ProductID:  input.ProductId,
		Quantity:   input.Quantity,
		TotalPrice: totalPrice,
	}

	transaction, err = s.repo.Create(transaction)
	if err != nil {
		return TransactionResponse(transaction), err
	}

	return TransactionResponse(transaction), nil
}

func (s service) UpdateItemProduct(id, quantity int) error {
	err := s.serviceProduct.UpdateItemProduct(id, quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s service) CheckUserBalance(userId, quantity, price int) (int, error) {
	u, err := s.serviceUser.GetUserById(userId)
	totalPrice := 0

	if err != nil {
		return totalPrice, err
	}

	balanceUser := u.Balance
	totalPrice = price * quantity

	if totalPrice >= balanceUser {
		return totalPrice, errors.New("Duit tidak cukup buat beli barang")
	}

	sisa := balanceUser - totalPrice

	userTopup := user.UserTopUpInput{
		Balance: sisa,
	}

	// update user balanace
	err = s.serviceUser.UpdateBalance(u.Email, &userTopup)
	if err != nil {
		return totalPrice, err
	}

	return totalPrice, nil
}

func (s service) CheckProduct(productId, quantity int) (*product.Product, error) {
	product, err := s.serviceProduct.GetProductById(productId)
	if err != nil {
		return product, err
	}

	err = s.CheckQuantityProduct(productId, quantity)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s service) CheckQuantityProduct(productId, quantity int) error {
	product, err := s.serviceProduct.GetProductById(productId)
	if err != nil {
		return err
	}

	if quantity >= product.Stock {
		return errors.New("Stock melebihi orderan")
	}

	return nil
}
