package model

import (
	"log"
	"strconv"
	"time"
)

type Order struct {
	quantity    int
	productCode string
}

func StartCreateOrderCommand(productCode string, quantity int, pl ProductList, cl CampaignList, t time.Time, csMap map[Campaign]CampaignSalesList) string {
	findProduct := Product{}
	order := Order{}

	for i, product := range pl {
		if product.code == productCode {
			pl[i].stock = product.stock - quantity
			findProduct = product
		}
	}
	if findProduct == (Product{}) {
		log.Fatal("Error: Product couldn't be found. Product Code: ", productCode)
	}
	order.CreateOrder(productCode, quantity)
	AddOrderToCampaignSalesList(order, findProduct, cl, t, csMap)
	return CreateOrderOutput(order)
}

func (order *Order) CreateOrder(productCode string, quantity int) {
	order.quantity = quantity
	order.productCode = productCode
}

func AddOrderToCampaignSalesList(order Order, p Product, cl CampaignList, t time.Time, csMap map[Campaign]CampaignSalesList) {

	newCampaignSales := CampaignSales{}
	campaignSalesList := CampaignSalesList{}

	for _, campaign := range cl {
		if campaign.productCode == order.productCode {
			status := GetCampaignStatus(campaign, t)
			if status == CampaignStatus(0) {
				p.CalculateProductPrice(t, campaign.duration, campaign.priceManipulationLimit)
				newCampaignSales.CreateCampaignSales(order.quantity, p.price)

				if csl, found := csMap[campaign]; found {
					csl.Add(newCampaignSales)
					csMap[campaign] = csl
				} else {
					campaignSalesList.Add(newCampaignSales)
					csMap[campaign] = campaignSalesList
				}
			}
		}
	}
}

func CreateOrderOutput(o Order) string {
	sQuantity := strconv.Itoa(o.quantity)

	return "Order created; product " + o.productCode + ", quantity " + sQuantity
}
