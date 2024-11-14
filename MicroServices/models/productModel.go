package models

import (
	"encoding/json"
	"io"
	"time"
)


type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"created_on"`
	UpdatedOn   string  `json:"updated_on"`
	DeletedOn   string  `json:"deleted_on,omitempty"`
}

// Mock data for products
var productList = []*Product{
	{
		ID:          1,
		Name:        "Laptop",
		Description: "A high-performance laptop for professionals",
		Price:       1200.99,
		SKU:         "LAP-1234",
		CreatedOn:   time.Now().Format(time.RFC3339),
		UpdatedOn:   time.Now().Format(time.RFC3339),
	},
	{
		ID:          2,
		Name:        "Smartphone",
		Description: "A smartphone with advanced features",
		Price:       799.49,
		SKU:         "SMT-5678",
		CreatedOn:   time.Now().Format(time.RFC3339),
		UpdatedOn:   time.Now().Format(time.RFC3339),
	},
	{
		ID:          3,
		Name:        "Wireless Headphones",
		Description: "Noise-cancelling wireless headphones",
		Price:       299.99,
		SKU:         "HD-91011",
		CreatedOn:   time.Now().Format(time.RFC3339),
		UpdatedOn:   time.Now().Format(time.RFC3339),
	},
}


type Products []*Product


func GetProducts()Products{
	return productList
}

func CreateProduct(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)


}
func getNextID() int{
	return productList[len(productList) - 1].ID + 1
}

func (products *Products) ToJSON(w io.Writer) error{
	return json.NewEncoder(w).Encode(products)
}

func (product *Product) FromJson(w io.Reader) error{
	return json.NewDecoder(w).Decode(product)
}