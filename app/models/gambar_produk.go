package models

import (
	"time"
)

type GambarProduk struct {
	ID              string      `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Produk          Produk
	ProdukID        string      `gorm:"size:36;index"`
	Path            string      `gorm:"type:text"`
	SangatBesar     string      `gorm:"type:text"`
	Besar           string      `gorm:"type:text"`
	Sedang          string      `gorm:"type:text"`
	Kecil           string      `gorm:"type:text"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}