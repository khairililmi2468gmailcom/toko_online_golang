package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)


type Pengiriman struct{
    ID              string          `gorm:"size:36;not null;uniqueIndex;primary_key"`
    User            User
    UserID          string          `gorm:"size:36;index"`
    Order           Order
    OrderID         string          `gorm:"size:36;index"`
    NoPengiriman    string          `gorm:"size:255;index"`
    Status          string          `gorm:"size:36;index"`
    JumlahBarang    int
    TotalBerat      decimal.Decimal `gorm:"type:decimal(10,2);"`
    NamaDepan       string          `gorm:"size:100;not null"`
    NamaBelakang    string          `gorm:"size:100;not null"`
    IdKota          string          `gorm:"size:100"`    
    IdProvinsi      string          `gorm:"size:100"`
    Alamat1         string          `gorm:"size:100"`
    Alamat2         string          `gorm:"size:100"`
    NoHp            string          `gorm:"size:50"`
    Email           string          `gorm:"size:100"`
    KodePos         string          `gorm:"size:100"`
    Pengirim        string          `gorm:"size:36"`
    PengirimanAt    time.Time
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       gorm.DeletedAt
}