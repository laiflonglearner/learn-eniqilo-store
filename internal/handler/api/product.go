package api

import (
	"net/http"

	"github.com/backend-magang/eniqilo-store/models"
	"github.com/backend-magang/eniqilo-store/models/entity"
	"github.com/backend-magang/eniqilo-store/utils/helper"
	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func (h *handler) GetListProduct(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.GetListProductRequest{
		ID:          c.QueryParam("id"),
		Limit:       c.QueryParam("limit"),
		Offset:      c.QueryParam("offset"),
		Name:        c.QueryParam("name"),
		IsAvailable: c.QueryParam("isAvailable"),
		Category:    c.QueryParam("category"),
		SKU:         c.QueryParam("sku"),
		InStock:     c.QueryParam("inStock"),
		Price:       c.QueryParam("price"),
		CreatedAt:   c.QueryParam("createdAt"),
	}

	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	if cast.ToInt(request.Limit) == 0 {
		request.Limit = "5"
	}

	if cast.ToInt(request.Offset) == 0 {
		request.Offset = "0"
	}

	resp := h.usecase.GetListProduct(ctx, request)
	return helper.WriteResponse(c, resp)
}

func (h *handler) CheckoutProduct(c echo.Context) (err error) {
	ctx, cancel := helper.GetContext()
	defer cancel()

	request := entity.CheckoutProductRequest{}
	err = pkg.BindValidate(c, &request)
	if err != nil {
		return helper.WriteResponse(c, models.StandardResponseReq{Code: http.StatusBadRequest, Error: err})
	}

	resp := h.usecase.CheckoutProduct(ctx, request)
	return helper.WriteResponse(c, resp)
}
