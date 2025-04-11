package items

import (
	"time"

	"github.com/lib/pq"
)

type Item struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	OwnerID     uint           `json:"owner_id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Condition   string         `json:"condition"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Images      pq.StringArray `gorm:"type:text[]" json:"image_urls"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
