package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hafizdkn/toko-belanja/apps/domain/handler/views"
	transactionhistory "github.com/hafizdkn/toko-belanja/apps/domain/transactionHistory"
)

type transactionHandler struct {
	service transactionhistory.IService
}

func NewTransactionHandler(service transactionhistory.IService) *transactionHandler {
	return &transactionHandler{service: service}
}

func (h *transactionHandler) Transaction(c *gin.Context) {
	var input transactionhistory.TransactionCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := views.UnprocessAbleEntityResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}
	userID := GetCurrentUserID(c)

	transaction, err := h.service.BuyItem(userID, &input)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(transaction, "You have succesfully purchased the product")
	views.WriteJsonRespnse(c, resp)
}

func (h *transactionHandler) GetMyTransactions(c *gin.Context) {
	transactions, err := h.service.GetMyTransactions()
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(transactions, "Success get my transaction")
	views.WriteJsonRespnse(c, resp)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	transactions, err := h.service.GetUserTransactions()
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(transactions, "Success get user transaction")
	views.WriteJsonRespnse(c, resp)
}
