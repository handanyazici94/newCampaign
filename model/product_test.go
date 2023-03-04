package model

import (
	"fmt"
	"testing"
	"time"
)

const (
	ProductInfoCommand   = "Product %s info; price %v, stock %d"
	CreateProductCommand = "Product created; code %s, price %v, stock %d"
)

func TestAddProduct(t *testing.T) {
	pl := ProductList{}
	beforeLen := len(pl)
	p := Product{}

	pl.AddProduct(p)

	if len(pl) != (beforeLen + 1) {
		t.Errorf("Product list couldn't added to list. List Length: %d", len(pl))
	}
}
func TestCreateProduct(t *testing.T) {

}
func TestProductInfo(t *testing.T) {

}
func TestCreateProductCommand(t *testing.T) {
	p := Product{code: "P1", price: 50, stock: 100}
	output := CreateProductOutput(p)
	expected := fmt.Sprintf(CreateProductCommand, p.code, p.price, p.stock)

	if output != expected {
		t.Errorf("Expected Output: %s , Returned Output: %s", expected, output)
	}

}
func TestProductInfoCommand(t *testing.T) {
	p := Product{code: "P1", price: 50, stock: 100}
	output := ProductInfoOutput(p)
	expected := fmt.Sprintf(ProductInfoCommand, p.code, p.price, p.stock)

	if output != expected {
		t.Errorf("Expected Output: %s , Returned Output: %s", expected, output)
	}
}

func TestCalculateProductPrice(t *testing.T) {
	timee := time.Time{}
	timee = timee.Add(time.Hour * time.Duration(2))
	campaignDuration := 3
	priceManipulationLimit := 30
	p := Product{code: "P1", price: 50, stock: 100}
	expectedPrice := float64(40)

	p.CalculateProductPrice(timee, campaignDuration, priceManipulationLimit)
	if p.price != expectedPrice {
		t.Errorf("Product price couldn't be decreased properly. Expected Price: %f, But Returned Price: %f ", expectedPrice, p.price)
	}
}
