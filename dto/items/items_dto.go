package items

import (
	"time"

	userdto "reuse-api/dto/user"
)

type ItemRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Category    string   `json:"category" binding:"required"`
	Condition   string   `json:"condition" binding:"required"`
	Latitude    float64  `json:"latitude" binding:"required"`
	Longitude   float64  `json:"longitude" binding:"required"`
	Images      []string `json:"image_urls"`
}

type ItemResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Condition   string    `json:"condition"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Images      []string  `json:"image_urls"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	OwnerName   string    `json:"owner_name"`
	OwnerID     uint      `json:"owner_id"`
}

type ItemResponseWithUser struct {
	Item ItemResponse         `json:"item"`
	User userdto.UserResponse `json:"user"`
}
