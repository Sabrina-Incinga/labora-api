package service

import "github.com/labora-api/ItemAPI/model"

func GetAllItems() []model.Item {
	return model.GetAll()
}