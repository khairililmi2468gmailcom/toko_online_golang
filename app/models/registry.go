package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Alamat{}},
		{Model: Produk{}},
		{Model: GambarProduk{}},
		{Model: Section{}},
		{Model: Kategori{}},
		{Model: Order{}},
		{Model: OrderItem{}},
		{Model: OrderCustomer{}},
		{Model: Pembayaran{}},
		{Model: Pengiriman{}},
	}
}