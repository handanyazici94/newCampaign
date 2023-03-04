package model

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Product struct {
	code  string
	price float64
	stock int
}

type ProductList []Product

func (pl *ProductList) StartCreateProductCommand(code string, price float64, stock int) string {
	findProduct := pl.FindProductByCode(code)

	if findProduct == (Product{}) {
		findProduct.CreateProduct(code, price, stock)
		pl.AddProduct(findProduct)
	} else {
		log.Fatal("Error: Product has already been.Product Code: ", code)
	}
	return CreateProductOutput(findProduct)
}

func (p *Product) CreateProduct(code string, price float64, stock int) {
	p.code = code
	p.price = price
	p.stock = stock
}

func (pl *ProductList) AddProduct(p Product) {
	*pl = append(*pl, p)
}

func CreateProductOutput(p Product) string {
	sStock := strconv.Itoa(p.stock)
	sPrice := fmt.Sprint(p.price)

	return "Product created; code " + p.code + ", price " + sPrice + ", stock " + sStock
}

func StartProductInfoCommand(code string, pl ProductList, cl CampaignList, t time.Time) string {
	findProduct := pl.FindProductByCode(code)

	if findProduct == (Product{}) {
		log.Fatal("Error: Product couldn't be found. Product Code: ", code)
	}

	for _, campaign := range cl {
		if campaign.productCode == code {
			status := GetCampaignStatus(campaign, t)
			if status == CampaignStatus(0) {
				findProduct.CalculateProductPrice(t, campaign.duration, campaign.priceManipulationLimit)
			}
		}
	}
	return ProductInfoOutput(findProduct)
}

func ProductInfoOutput(p Product) string {
	sStock := strconv.Itoa(p.stock)
	sPrice := fmt.Sprint(p.price)

	return "Product " + p.code + " info; price " + sPrice + ", stock " + sStock
}

func (product *Product) CalculateProductPrice(t time.Time, campaignDuration int, priceManipulationLimit int) {
	timeHour := t.Hour()
	disCountAlgorithm := (priceManipulationLimit / campaignDuration) * timeHour
	product.price = product.price - ((product.price * float64(disCountAlgorithm)) / 100)
}

func (pl *ProductList) FindProductByCode(productCode string) Product {
	for _, product := range *pl {
		if product.code == productCode {
			return product
		}
	}
	return Product{}

}
