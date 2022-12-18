package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/hafizdkn/toko-belanja/apps/domain/category"
	"github.com/hafizdkn/toko-belanja/apps/domain/handler/views"
)

type categoryHandler struct {
	service category.IService
}

func NewCategoryHandler(service category.IService) *categoryHandler {
	return &categoryHandler{service: service}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var input category.CategoryCreateInput

	if err := c.ShouldBindJSON(&input); err != nil {
		resp := views.UnprocessAbleEntityResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	category, err := h.service.CreateCategory(&input)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(category, "Success create category")
	views.WriteJsonRespnse(c, resp)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	var input category.CategoryUpdateInput

	categoryId, err := strconv.Atoi(c.Param("categoryId"))
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

	category, err := h.service.UpdateCategory(&input, categoryId)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(category, "Success update category")
	views.WriteJsonRespnse(c, resp)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	categoryId, err := strconv.Atoi(c.Param("categoryId"))
	if err != nil {
		resp := views.BadRequestResponse(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	err = h.service.DeleteCategory(categoryId)
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(nil, "Success delete category")
	views.WriteJsonRespnse(c, resp)
}

func (h *categoryHandler) GetCategorys(c *gin.Context) {
	categorys, err := h.service.GetCategorys()
	if err != nil {
		resp := views.InternalServerError(err)
		views.WriteJsonRespnse(c, resp)
		return
	}

	resp := views.SuccessCreateResponse(categorys, "Success get categorys")
	views.WriteJsonRespnse(c, resp)
}
