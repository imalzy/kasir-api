package transaction

import (
	"errors"
)

type TransactionService struct {
	repo *TransactionRepository
}

func NewTransactionService(repo *TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []CheckoutItem) (*Transaction, error) {
	if len(items) == 0 {
		return nil, errors.New("checkout items cannot be empty")
	}

	transaction, err := s.repo.CreateTransaction(items)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
