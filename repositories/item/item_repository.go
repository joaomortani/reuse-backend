package repositories

import (
	models "reuse-api/models/items"

	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *models.Item) error
	GetAll() ([]models.Item, error)
	GetById(id uint) (*models.Item, error)
}

type itemRepository struct {
	database *gorm.DB
}

func NewItemRepository(database *gorm.DB) ItemRepository {
	return &itemRepository{database: database}
}

func (r *itemRepository) Create(item *models.Item) error {
	return r.database.Create(item).Error
}

func (r *itemRepository) GetAll() ([]models.Item, error) {
	var items []models.Item
	err := r.database.Find(&items).Error
	return items, err
}

func (r *itemRepository) GetById(id uint) (*models.Item, error) {
	var item models.Item
	err := r.database.First(&item, id).Error
	return &item, err
}
