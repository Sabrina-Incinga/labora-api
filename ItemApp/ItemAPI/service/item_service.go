package service

import (
	"github.com/labora-api/ItemAPI/model"
	"github.com/labora-api/ItemAPI/repository/interfaces"
)

type ItemService struct {
    DbHandler interfaces.DBHandler
}

func (s *ItemService) GetAllItems(page int, itemsPerPage int) model.ItemsResponse {
	return s.DbHandler.GetAll(page, itemsPerPage)
}

func (s *ItemService) GetItemById(id int) *model.ItemResponse{
	return s.DbHandler.GetItemById(id)
}

func (s *ItemService) CreateItem(dto model.ItemDTO) int64{
	return s.DbHandler.Create(dto)
}

func (s *ItemService) UpdateItem(dto model.ItemDTO, id int) bool{
	return s.DbHandler.Update(dto, id) > 0
}

func (s *ItemService) DeleteItem(id int) bool{
	return s.DbHandler.Delete(id) > 0
}