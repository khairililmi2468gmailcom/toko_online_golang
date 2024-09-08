package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Produk struct {
	ID               string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	ParentID         string `gorm:"size:36;index"`
	User             User
	UserID           string `gorm:"size:36;index"`
	ProductImages    []GambarProduk
	Kategori         []Kategori      `gorm:"many2many:product_categories;"`
	Sku              string          `gorm:"size:100;index"`
	Nama             string          `gorm:"size:255"`
	Slug             string          `gorm:"size:255"`
	Harga            decimal.Decimal `gorm:"type:decimal(16,2);"`
	Stock            int
	Berat            decimal.Decimal `gorm:"type:decimal(10,2);"`
	DeskripsiSingkat string          `gorm:"size:255"`
	Deskripsi        string          `gorm:"type:text"`
	Status           int             `gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (p *Produk) GetProducts(db *gorm.DB, perPage int, page int) (*[]Produk, int64, error) {
	var err error
	var products []Produk
	var count int64

	err = db.Debug().Model(&Produk{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * perPage

	err = db.Debug().Model(&Produk{}).Order("created_at desc").Limit(perPage).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return &products, count, nil
}