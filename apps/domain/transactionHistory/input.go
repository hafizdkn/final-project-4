package transactionhistory

type TransactionCreateInput struct {
	ProductId int `json:"product_id,omitempty"`
	Quantity  int `json:"quantity,omitempty"`
}
