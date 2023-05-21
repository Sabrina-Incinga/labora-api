package service

import (
	"github.com/labora-api/ItemAPI/model"
	"github.com/labora-api/ItemAPI/repository"
)

func GetAllItems(page int, itemsPerPage int) model.ItemsResponse {
	return repository.GetAll(page, itemsPerPage)
}

func GetItemById(id int) *model.ItemResponse{
	return repository.GetItemById(id)
}

func CreateItem(dto model.ItemDTO) int64{
	return repository.Create(dto)
}

func UpdateItem(dto model.ItemDTO, id int) bool{
	return repository.Update(dto, id) > 0
}

func DeleteItem(id int) bool{
	return repository.Delete(id) > 0
}