package fakers

import (
	"time"

	"github.com/bxcodec/faker/v3"

	"github.com/google/uuid"
	"github.com/khairililmi2468gmailcom/toko_online_golang/app/models"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		NamaDepan:     faker.FirstName(),
		NamaBelakang:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "$2y$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi", // password
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}