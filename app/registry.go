package app

import "github.com/khairililmi2468gmailcom/toko_online_golang/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Alamat{}},
		{Model: models.Produk{}},
		{Model: models.GambarProduk{}},
		{Model: models.Section{}},
		{Model: models.Kategori{}},
		{Model: models.Order{}},
		{Model: models.OrderItem{}},
		{Model: models.OrderCustomer{}},
		{Model: models.Pembayaran{}},
		{Model: models.Pengiriman{}},
	}
}