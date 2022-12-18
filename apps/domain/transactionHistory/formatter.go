package transactionhistory

func TransactionResponse(tran *Transaction) *Transaction {
	return &Transaction{
		TotalPrice: tran.TotalPrice,
		Quantity:   tran.Quantity,
		ProductID:  tran.ProductID,
	}
}
