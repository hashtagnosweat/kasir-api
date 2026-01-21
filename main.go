package main

// @title Produk API
// @version 1.0
// @description API untuk manajemen produk
// @host localhost:8080
// @BasePath /

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	_ "kasir-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
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

// @Summary Get all produk
// @Description Mengambil semua data produk
// @Tags produk
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/produk [get]
func getAllProduk(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, produk, "Berhasil mendapatkan data")
}

// @Summary Add new produk
// @Description Menambahkan data produk baru
// @Tags produk
// @Param produk body Produk true "Data Produk"
// @Produce json
// @Success 201 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/produk [post]
func addProduk(w http.ResponseWriter, r *http.Request) {
	var produkBaru Produk
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		respondError(w, "Invalid request", http.StatusBadRequest)
		return
	}
	produkBaru.ID = len(produk) + 1
	produk = append(produk, produkBaru)
	respondJSON(w, http.StatusCreated, produkBaru, "Berhasil menambahkan data")
}

// @Summary Get produk by ID
// @Description Mengambil produk berdasarkan ID
// @Tags produk
// @Param id path int true "Produk ID"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/produk/{id} [get]
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

// @Summary Update produk
// @Description Mengubah data produk berdasarkan ID
// @Tags produk
// @Param id path int true "Produk ID"
// @Param produk body Produk true "Data Produk"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/produk/{id} [put]
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

// @Summary Delete produk
// @Description Menghapus produk berdasarkan ID
// @Tags produk
// @Param id path int true "Produk ID"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/produk/{id} [delete]
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
	// Routes for ID-based operations
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukById(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}

	})

	// Routes for collection-based operations
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getAllProduk(w, r)
		} else if r.Method == "POST" {
			addProduk(w, r)
		}
	})

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, nil, "API is running")
	})
	fmt.Println("Server running di :8080")

	// Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
