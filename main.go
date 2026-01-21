package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

type APIResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "Kecap bang aw", Harga: 12000, Stok: 20},
}

// Helper

func respondJSON(w http.ResponseWriter, status int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(APIResponse{
		Status:  status,
		Data:    data,
		Message: message,
	})
}

func respondError(w http.ResponseWriter, message string, status int) {
	respondJSON(w, status, nil, message)
}

// Func

func getProdukById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			respondJSON(w, http.StatusOK, p, "Berhasil mendapatkan data produk dengan ID")
			return
		}
	}

	respondError(w, "Produk belum ada", http.StatusNotFound)
}

// PUT localhost:8080/api/produk/{id}
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	// ganti jadi int
	// loop data produk, cari yang sesuai id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondError(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		respondError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			respondJSON(w, http.StatusOK, updateProduk, "Update berhasil")
			return
		}
	}

	respondError(w, "Produk belum ada", http.StatusNotFound)
}

func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	// ganti id int
	// loop produk cari ID, dapet index yang mau dihapus
	// bikin slice baru dengan data sebelum dan sesudah index

	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondError(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			produk = append(produk[:i], produk[i+1:]...)
			respondJSON(w, http.StatusOK, nil, "Delete berhasil")
			return
		}
	}

	respondError(w, "Produk belum ada", http.StatusNotFound)
}

func main() {
	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukById(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}

	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			respondJSON(w, http.StatusOK, produk, "Berhasil mendapatkan data")
		} else if r.Method == "POST" {
			// baca data dari request
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				respondError(w, "Invalid request", http.StatusBadRequest)
			}

			// masukkin data ke dalam variable produk
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			respondJSON(w, http.StatusCreated, produkBaru, "Berhasil menambahkan data")
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, nil, "API is running")
	})
	fmt.Println("Server running di :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
