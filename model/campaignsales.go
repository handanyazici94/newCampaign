package model

type CampaignSales struct {
	count int
	price float64
}
type CampaignSalesList []CampaignSales

func (cs *CampaignSales) CreateCampaignSales(count int, price float64) {
	cs.count = count
	cs.price = price
}

func (csl *CampaignSalesList) Add(cs CampaignSales) {
	*csl = append(*csl, cs)
}

func CalculateCampaignSalesInfo(csMap map[Campaign]CampaignSalesList, campaign Campaign) (int, int, int) {
	totalSales := 0
	totalPrice := 0
	averageItemPrice := 0
	turnOver := 0
	if campaignSalesList, found := csMap[campaign]; found {
		for _, campaignSales := range campaignSalesList {
			totalSales = totalSales + campaignSales.count
			totalPrice = totalPrice + (campaignSales.count * int(campaignSales.price))
		}
		averageItemPrice = (totalPrice / totalSales)
		turnOver = totalSales * averageItemPrice
	}
	return totalSales, turnOver, averageItemPrice
}
