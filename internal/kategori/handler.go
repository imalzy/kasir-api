package kategori

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func GetKategoriByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID - "+err.Error(), http.StatusBadRequest)
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

func UpdateKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID - "+err.Error(), http.StatusBadRequest)
		return
	}

	var updated Kategori
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

func DeleteKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/kategori/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID - "+err.Error(), http.StatusBadRequest)
		return
	}

	errDelete := Delete(id)
	if errDelete == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

func ListKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := List()
	json.NewEncoder(w).Encode(p)
}

func CreateKategori(w http.ResponseWriter, r *http.Request) {
	var newKategori Kategori
	if err := json.NewDecoder(r.Body).Decode(&newKategori); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	p := Create(newKategori)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
