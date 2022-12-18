package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hafizdkn/toko-belanja/apps/domain/handler/views"
	"github.com/hafizdkn/toko-belanja/apps/domain/product"
)

type productHandler struct {
	service product.IService
}

func NewProductHandler(service product.IService) *productHandler {
	return &productHandler{service: service}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.ProductCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := views.UnprocessAbleEntityResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	product, err := h.service.CreateProduct(&input)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(product, "Success create category")
	views.WriteJsonRespnse(c, resp)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	var input product.ProductUpdateInput

	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		resp := views.BadRequestResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := views.UnprocessAbleEntityResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	product, err := h.service.UpdateProduct(productId, &input)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(product, "Success update category")
	views.WriteJsonRespnse(c, resp)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		resp := views.BadRequestResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	err = h.service.DeleteProduct(productId)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(nil, "Product has been successfully deleted")
	views.WriteJsonRespnse(c, resp)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	products, err := h.service.GetProduct()
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return

	}

	resp := views.SuccessCreateResponse(products, "Success get products")
	views.WriteJsonRespnse(c, resp)
}
