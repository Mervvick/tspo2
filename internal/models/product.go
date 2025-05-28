// // internal/models/product.go
// package models

// import (
// 	"time"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// 	"gorm.io/datatypes"
// )

// type Product struct {
// 	ID             uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
// 	Name           string         `gorm:"not null" json:"name"`
// 	Slug           string         `gorm:"uniqueIndex;not null" json:"slug"`
// 	Description    string         `json:"description"`
// 	CategoryID     uuid.UUID      `gorm:"type:uuid;index" json:"category_id"`
// 	Category       Category       `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
// 	Price          float64        `gorm:"not null;index" json:"price"`
// 	Stock          int            `gorm:"not null;default:0" json:"stock"`
// 	ImageURLs      datatypes.JSON `gorm:"type:jsonb" json:"image_urls"`
// 	Specifications datatypes.JSON `gorm:"type:jsonb" json:"specifications"`
// 	CreatedAt      time.Time      `gorm:"autoCreateTime" json:"created_at"`
// 	UpdatedAt      time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
// }

// func (p *Product) BeforeCreate(tx *gorm.DB) error {
// 	if p.ID == uuid.Nil {
// 		p.ID = uuid.New()
// 	}
// 	return nil
// }

// Файл: internal/models/product.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Slug        string    `gorm:"uniqueIndex;not null" json:"slug"`
	Description string    `json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null;default:0" json:"stock"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
