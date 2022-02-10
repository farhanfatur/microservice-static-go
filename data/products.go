package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	SKU         string `json:"sku"`
	CreatedOn   string `json:"-"`
	UpdatedOn   string `json:"-"`
	DeletedOn   string `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// ============= convert json =============

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// ============= Abstract function product =============

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = GetNextID()
	productList = append(productList, p)
}

var ProductNotFoundMsg = fmt.Errorf("Product Not Found")

func UpdateProduct(id int, p *Product) error {
	_, index, err := FindProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[index] = p
	return nil
}

func FindProduct(id int) (*Product, int, error) {
	for i, each := range productList {
		if each.ID == id {
			return each, i, nil
		}
	}

	return nil, -1, ProductNotFoundMsg
}

func GetNextID() int {
	prod := productList[len(productList)-1]
	return prod.ID + 1
}

// ============= Static Data =============
var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Latte Milk",
		Price:       10000,
		SKU:         "absd123",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Black Coffee",
		Description: "Would make you off",
		Price:       10000,
		SKU:         "asd5553",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}
