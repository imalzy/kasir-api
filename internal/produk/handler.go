package produk

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID - "+err.Error(), http.StatusBadRequest)
		return
	}

	p, err := GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID - "+err.Error(), http.StatusBadRequest)
		return
	}

	var updated Product
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid request body - "+err.Error(), http.StatusBadRequest)
		return
	}

	p, err := Update(id, updated)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
		return
	}
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID - "+err.Error(), http.StatusBadRequest)
		return
	}

	errDelete := Delete(id)
	if errDelete == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

func ListProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := List()
	json.NewEncoder(w).Encode(p)
}

func CreateProduk(w http.ResponseWriter, r *http.Request) {
	var newProduk Product
	if err := json.NewDecoder(r.Body).Decode(&newProduk); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	p := Create(newProduk)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
