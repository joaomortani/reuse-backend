package item

import (
	dto "reuse-api/dto/items"
	"reuse-api/models/items"
	repositories "reuse-api/repositories/item"
	"time"
)

type ItemService interface {
	Create(req dto.ItemRequest, userID uint) (*items.Item, error)
	GetAll() ([]dto.ItemResponse, error)
	GetById(id uint) (*dto.ItemResponse, error)
}

type itemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) ItemService {
	return &itemService{repo: repo}

}

func (s *itemService) Create(req dto.ItemRequest, userID uint) (*items.Item, error) {
	item := &items.Item{
		Title:     req.Title,
		Category:  req.Category,
		Condition: req.Condition,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		OwnerID:   userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *itemService) GetAll() ([]dto.ItemResponse, error) {
	items, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.ItemResponse
	for _, item := range items {
		responses = append(responses, dto.ItemResponse{
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Category:    item.Category,
			Condition:   item.Condition,
			Latitude:    item.Latitude,
			Longitude:   item.Longitude,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			OwnerID:     item.OwnerID,
		})
	}

	return responses, nil

}

func (s *itemService) GetById(id uint) (*dto.ItemResponse, error) {
	item, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	response := &dto.ItemResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Category:    item.Category,
		Condition:   item.Condition,
		Latitude:    item.Latitude,
		Longitude:   item.Longitude,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
		OwnerID:     item.OwnerID,
	}
	return response, nil
}
