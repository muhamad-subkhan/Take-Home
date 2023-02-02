package router

import (
	"ShoppingCart/handler"
	"ShoppingCart/pkg/mysql"
	"ShoppingCart/repositories"

	"github.com/gorilla/mux"
)

func Router(r *mux.Router) {
	repositories := repositories.Repo(mysql.DB)
	h := handler.Handler(repositories)

	r.HandleFunc("/add-produk", h.AddProduk).Methods("POST") // Menambahkan Produk
	r.HandleFunc("/update-produk/{id}", h.UpdateProduk).Methods("PATCH") // mengupdate Kuantitas

	r.HandleFunc("/delete-produk/{id}", h.DeleteProduk).Methods("DELETE")


	// r.HandleFunc("/filter/{name}", h.FilterByName).Methods{"GET"}
	r.HandleFunc("/filter/{name}", h.FilterByName).Methods("GET")
	r.HandleFunc("/filter-qty/{qty}", h.FilterByQty).Methods("GET")
}