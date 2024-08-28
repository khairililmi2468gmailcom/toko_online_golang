package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
    ID                  string             `gorm:"size:36;not null;uniqueIndex;primary_key"`
    UserID              string             `gorm:"size:36;index"`
    User                User   
    OrderItems          []OrderItem
    OrderCustomer       *OrderCustomer
    Code                string `            gorm:"size:50;index"` 
    Status              int
    OrderDate           time.Time
    PaymentDue          time.Time
    PaymentStatus       string              `gorm:"size:50;index"`
    PaymentToken        string              `gorm:"size:100;index"`
    BaseTotalHarga      decimal.Decimal     `gorm:"type:decimal(16,2)"`
    TotalBiayaAdmin     decimal.Decimal     `gorm:"type:decimal(16,2)"`
    PersenPajak         decimal.Decimal     `gorm:"type:decimal(10,2)"`
    TotalDiskon         decimal.Decimal     `gorm:"type:decimal(16,2)"`
    PersenDiskon        decimal.Decimal     `gorm:"type:decimal(10,2)"`
    BiayaPengiriman     decimal.Decimal     `gorm:"type:decimal(16,2)"`
    TotalKeseluruhan    decimal.Decimal     `gorm:"type:decimal(16,2)"`
    Catatan             string              `gorm:"type:text"`
    KurirPengirim       string              `gorm:"size:100"`
    ServicePengirim     string              `gorm:"size:100"`
    DisetujuiOleh       string              `gorm:"size:36"`
    ApprovedAt          time.Time
    DibatalkanOleh      string              `gorm:"size:36"`
    CancelledAt         time.Time
    CancellationNote    string `gorm:"size:255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}