package controller

import (
	"net/http"

	// "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/onizukazaza/tarzer-shop-api-tu/pkg/custom"
	_itemShopModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/model"
	_itemShopService "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemShop/service"

)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService,
) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}
func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)

	if err := customEchoRequest.Bind(itemFilter); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}

	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError,err.Error())
	}
	return pctx.JSON(http.StatusOK, itemModelList)
	// return custom.Error(pctx, http.StatusInternalServerError, (&_itemShopException.ItemListing{}).Error())
}
