package model

import (
	_itemShopModel "tutorial/pkg/itemShop/model"
)

type (
	Inventory struct {
		Item    *_itemShopModel.Item `json:"item"`
		Quanity uint                 `json:"quantity"`
	}

	ItemQuantityCounting struct {
		ItemID   uint64
		Quantity uint
	}
)
