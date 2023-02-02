package handler

import (
	"ShoppingCart/Result"
	"ShoppingCart/models"
	"ShoppingCart/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type ProductRequest struct {
	KodeProduct uint   `json:"Kode"`
	NameProduct string `json:"name"`
	QtyProduct  uint   `json:"qty"`
}

type handler struct {
	repositories repositories.Repositories
}

func Handler(repositories repositories.Repositories) *handler {
	return &handler{repositories}
}

func (h *handler) AddProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := new(ProductRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validator := validator.New()
	if err := validator.Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	entity := models.Product{
		KodeProduct: req.KodeProduct,
		NameProduct: req.NameProduct,
		QtyProduct:  req.QtyProduct,
	}

	data, err := h.repositories.Create(entity)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result.ErrorResult{Code: http.StatusBadRequest, Message: "error while creating"}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Result.SuccesResult{Code: http.StatusOK, Message: "Succes", Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) GetProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	get, err := h.repositories.GetProduk(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := Result.ErrorResult{Code: http.StatusBadRequest, Message: "error Getting product"}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Result.SuccesResult{Code: http.StatusOK, Message: "Succes", Data: get}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) UpdateProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := new(ProductRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(respone)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "ID Tidak Ditemukan !!!"}
		json.NewEncoder(w).Encode(respone)
		return
	}

	product, err := h.repositories.GetProduk(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "Product Not Found"}
		json.NewEncoder(w).Encode(respone)
		return
	}

	if req.QtyProduct != 0 {
		product.QtyProduct = req.QtyProduct
	}

	data, err := h.repositories.UpdateProduk(product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "Update Failed"}
		json.NewEncoder(w).Encode(respone)
		return
	}

	response := Result.SuccesResult{Code: http.StatusOK, Message: "Succes", Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) DeleteProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.repositories.GetProduk(id)
	if err != nil {
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "ID Not Found"}
		json.NewEncoder(w).Encode(respone)
		return
	}

	data, err := h.repositories.DeleteProduk(product)
	if err != nil {
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "Failed to Delete Product"}
		json.NewEncoder(w).Encode(respone)
		return
	}

	response := Result.SuccesResult{Code: http.StatusOK, Message: "Succes", Data: data}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) FindProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	film, err := h.repositories.FindProduk()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "Can't Find Product"}
		json.NewEncoder(w).Encode(respone)
	}

	w.WriteHeader(http.StatusOK)
	response := Result.SuccesResult{Code: http.StatusOK, Data: film}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) FilterByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	name := vars["name"]

	data , err := h.repositories.FilterByName(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "Can't Find Product"}
		json.NewEncoder(w).Encode(respone)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := Result.SuccesResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}


func (h *handler) FilterByQty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	qty, _ := strconv.Atoi(vars["qty"])

	data , err := h.repositories.FilterByQty(qty)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respone := Result.ErrorResult{Code: http.StatusBadRequest, Message: "Can't Find Product"}
		json.NewEncoder(w).Encode(respone)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := Result.SuccesResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)

}
