package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID              string          `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Order           Order
	OrderID         string          `gorm:"size:36;index"`
	Produk          Produk
	ProdukID        string          `gorm:"size:36;index"`
	Qty             int
	BaseTotalHarga  decimal.Decimal `gorm:"type:decimal(16,2)"`
	TotalBiayaAdmin decimal.Decimal `gorm:"type:decimal(16,2)"`
	JumlahPajak     decimal.Decimal `gorm:"type:decimal(16,2)"`
	PersenPajak     decimal.Decimal `gorm:"type:decimal(10,2)"`
	TotalDiskon     decimal.Decimal `gorm:"type:decimal(16,2)"`
	PersenDiskon    decimal.Decimal `gorm:"type:decimal(10,2)"`
	SubTotal        decimal.Decimal `gorm:"type:decimal(16,2)"`
	Sku             string          `gorm:"size:36;index"`
	Nama            string          `gorm:"size:255"`
	Berat           decimal.Decimal `gorm:"type:decimal(10,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}