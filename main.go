package main

// @title Product API
// @version 1.0
// @description API for product management
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

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type APIResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

var products = []Product{
	{ID: 1, Name: "Indomie", Price: 3500, Stock: 10},
	{ID: 2, Name: "Vit 1000ml", Price: 3000, Stock: 40},
	{ID: 3, Name: "Sweet Soy Sauce", Price: 12000, Stock: 20},
}

var categories = []Category{
	{ID: 1, Name: "Food"},
	{ID: 2, Name: "Stationery"},
	{ID: 3, Name: "Tools"},
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

// PRODUCTS API

// @Summary Get all products
// @Description Retrieve all product data
// @Tags products
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/products [get]
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, products, "Successfully retrieved data")
}

// @Summary Add new product
// @Description Add a new product
// @Tags products
// @Param product body Product true "Product Data"
// @Produce json
// @Success 201 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/products [post]
func addProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		respondError(w, "Invalid request", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	respondJSON(w, http.StatusCreated, newProduct, "Successfully added data")
}

// @Summary Get product by ID
// @Description Retrieve product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/products/{id} [get]
func getProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for _, p := range products {
		if p.ID == id {
			respondJSON(w, http.StatusOK, p, "Successfully retrieved product data")
			return
		}
	}

	respondError(w, "Product not found", http.StatusNotFound)
}

// @Summary Update product
// @Description Update product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Param product body Product true "Product Data"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/products/{id} [put]
func updateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondError(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		respondError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range products {
		if products[i].ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct
			respondJSON(w, http.StatusOK, updatedProduct, "Update successful")
			return
		}
	}

	respondError(w, "Product not found", http.StatusNotFound)
}

// @Summary Delete product
// @Description Delete product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/products/{id} [delete]
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondError(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			respondJSON(w, http.StatusOK, nil, "Delete successful")
			return
		}
	}

	respondError(w, "Product not found", http.StatusNotFound)
}

// CATEGORIES API

// @Summary Get all categories
// @Description Retrieve all category data
// @Tags categories
// @Produce json
// @Success 200 {object} APIResponse
// @Router /api/categories [get]
func getAllCategories(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, categories, "Successfully retrieved data")
}

// @Summary Add new category
// @Description Add a new category
// @Tags categories
// @Param category body Category true "Category Data"
// @Produce json
// @Success 201 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Router /api/categories [post]
func addCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		respondError(w, "Invalid request", http.StatusBadRequest)
		return
	}
	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)
	respondJSON(w, http.StatusCreated, newCategory, "Successfully added data")
}

// @Summary Get category by ID
// @Description Retrieve category by ID
// @Tags categories
// @Param id path int true "Category ID"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/categories/{id} [get]
func getCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, c := range categories {
		if c.ID == id {
			respondJSON(w, http.StatusOK, c, "Successfully retrieved category data")
			return
		}
	}

	respondError(w, "Category not found", http.StatusNotFound)
}

// @Summary Update category
// @Description Update category by ID
// @Tags categories
// @Param id path int true "Category ID"
// @Param category body Category true "Category Data"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/categories/{id} [put]
func updateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondError(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updatedCategory Category
	err = json.NewDecoder(r.Body).Decode(&updatedCategory)
	if err != nil {
		respondError(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].ID == id {
			updatedCategory.ID = id
			categories[i] = updatedCategory
			respondJSON(w, http.StatusOK, updatedCategory, "Update successful")
			return
		}
	}

	respondError(w, "Category not found", http.StatusNotFound)
}

// @Summary Delete category
// @Description Delete category by ID
// @Tags categories
// @Param id path int true "Category ID"
// @Produce json
// @Success 200 {object} APIResponse
// @Failure 400 {object} APIResponse
// @Failure 404 {object} APIResponse
// @Router /api/categories/{id} [delete]
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		respondError(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, c := range categories {
		if c.ID == id {
			categories = append(categories[:i], categories[i+1:]...)
			respondJSON(w, http.StatusOK, nil, "Delete successful")
			return
		}
	}

	respondError(w, "Category not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductById(w, r)
		} else if r.Method == "PUT" {
			updateProduct(w, r)
		} else if r.Method == "DELETE" {
			deleteProduct(w, r)
		}
	})

	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getAllProducts(w, r)
		} else if r.Method == "POST" {
			addProduct(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryById(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getAllCategories(w, r)
		} else if r.Method == "POST" {
			addCategory(w, r)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		respondJSON(w, http.StatusOK, nil, "API is running")
	})

	fmt.Println("Server running on :8080")

	http.Handle("/swagger/", httpSwagger.WrapHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to run server")
	}
}
