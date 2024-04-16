package repository

import (
	_itemShopModel "github.com/Shiryuji/ryuji-shop-api/pkg/itemShop/model"

	"github.com/Shiryuji/ryuji-shop-api/entities"
)

type ItemShopRepository interface{
	Listing(itemFilter *_itemShopModel.ItemFilter) ([]*entities.Item, error) 
	Counting(itemFilter *_itemShopModel.ItemFilter) (int64, error)
	FindByID(itemID uint64) (*entities.Item, error)
}