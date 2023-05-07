package service

import "github.com/labora-api/ItemAPI/model"

func GetAllItems() []model.Item {
	return model.GetAll()
}

func GetItemById(id int) *model.Item{
	return model.GetItemById(id)
}

func CreateItem(dto model.ItemDTO) int64{
	return model.Create(dto)
}

func UpdateItem(dto model.ItemDTO, id int) bool{
	return model.Update(dto, id) > 0
}

func DeleteItem(id int) bool{
	return model.Delete(id) > 0
}