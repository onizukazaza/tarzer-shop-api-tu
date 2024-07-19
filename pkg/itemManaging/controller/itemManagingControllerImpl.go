package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/onizukazaza/tarzer-shop-api-tu/pkg/custom"
	_itemManagingModel "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/model"
	_itemManagingService "github.com/onizukazaza/tarzer-shop-api-tu/pkg/itemManaging/service"
	"net/http"
)

type itemManagingControllerImpl struct {
	itemManagingService _itemManagingService.ItemManagingService
}

func NewItemManagingControllerImpl(
	itemManagingService _itemManagingService.ItemManagingService,
) ItemManagingController {
	return &itemManagingControllerImpl{itemManagingService}
}

func (c *itemManagingControllerImpl) Creating(pctx echo.Context) error {
	itemCreatingReq := new(_itemManagingModel.ItemCreatingReq)
	customEchoRequest := custom.NewCustomEchoRequest(pctx)
	if err := customEchoRequest.Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err.Error())
	}
	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusCreated, item)
}
