package interfaces

import "github.com/labora-api/ItemAPI/model"

type DBHandler interface {
	GetAll(page int, itemsPerPage int) model.ItemsResponse
	Create(item model.ItemDTO) int64
	GetItemById(id int) *model.ItemResponse
	Update(dto model.ItemDTO, id int) int64
	Delete(id int) int64
}