package transactionhistory

import "gorm.io/gorm"

type IRepository interface {
	Create(transaction *Transaction) (*Transaction, error)
	GetMyTransactions() ([]*Transaction, error)
	GetUserTransactions() ([]*Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) IRepository {
	return &repository{db}
}

func (r *repository) Create(transaction *Transaction) (*Transaction, error) {
	err := r.db.Debug().Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *repository) GetMyTransactions() ([]*Transaction, error) {
	var transactions []*Transaction

	err := r.db.Debug().Preload("Product").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *repository) GetUserTransactions() ([]*Transaction, error) {
	var transactions []*Transaction

	err := r.db.Debug().Preload("Product").Preload("User").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
