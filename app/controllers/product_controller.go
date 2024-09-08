package controllers

import (
	"net/http"
	"strconv"

	"github.com/khairililmi2468gmailcom/toko_online_golang/app/models"

	"github.com/unrolled/render"
)

func (server *Server) Produk(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := models.Produk{}
	products, totalRows, err := productModel.GetProducts(server.DB, perPage, page)
	if err != nil {
		return
	}

	pagination, _ := GetPaginationLinks(server.AppConfig, PaginationParams{
		Path:        "products",
		TotalRows:   int32(totalRows),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products":   products,
		"pagination": pagination,
	})
}