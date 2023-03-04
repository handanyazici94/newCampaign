package model

import (
	"log"
	"strconv"
	"time"
)

type Campaign struct {
	name                   string
	duration               int
	priceManipulationLimit int
	targetSales            int
	productCode            string
}
type CampaignList []Campaign
type CampaignStatus int

const (
	Active CampaignStatus = iota
	Ended
)

func (cs CampaignStatus) String() string {
	return [...]string{"Active", "Ended"}[cs]
}

func (cl *CampaignList) StartCreateCampaignCommand(name string, productCode string, duration int, priceManipulationLimit int, targetSalesCount int, pl ProductList) string {
	campaign := cl.FindCampaignByName(name)

	if campaign == (Campaign{}) {
		findProduct := pl.FindProductByCode(productCode)
		if findProduct == (Product{}) {
			log.Fatal("Error: Product couldn't be found. Product Code: ", productCode)
		}

		campaign.name = name
		campaign.priceManipulationLimit = priceManipulationLimit
		campaign.targetSales = targetSalesCount
		campaign.duration = duration
		campaign.productCode = productCode
		cl.AddCampaign(campaign)
	} else {
		log.Fatal("Error: Campaign has already been.Campaign Name: ", name)
	}
	return CreateCampaignOutput(campaign)
}

func (cl *CampaignList) AddCampaign(c Campaign) {
	*cl = append(*cl, c)
}

func CreateCampaignOutput(c Campaign) string {
	sTargetSales := strconv.Itoa(c.targetSales)
	sDuration := strconv.Itoa(c.duration)
	sLimit := strconv.Itoa(c.priceManipulationLimit)

	return "Campaign created; name " + c.name + ", product " + c.productCode + ", duration " + sDuration + ", limit " +
		sLimit + ", target sales count " + sTargetSales
}

func StartCampaignInfoCommand(name string, cl CampaignList, t time.Time, csMap map[Campaign]CampaignSalesList) string {

	findCampaign := cl.FindCampaignByName(name)
	if findCampaign == (Campaign{}) {
		log.Fatal("Error: Campaign couldn't be found. Campaign Name: ", name)
	}
	campaignStatus := GetCampaignStatus(findCampaign, t)
	totalSales, turnOver, averageItemPrice := CalculateCampaignSalesInfo(csMap, findCampaign)

	return GetCampaignInfoOutput(findCampaign, campaignStatus, totalSales, turnOver, averageItemPrice)
}

func GetCampaignInfoOutput(c Campaign, status CampaignStatus, totalSales int, turnover int, averageItemPrice int) string {
	sTargetSales := strconv.Itoa(c.targetSales)
	sTotalSales := strconv.Itoa(totalSales)
	sTurnover := strconv.Itoa(turnover)
	sAverageItemPrice := strconv.Itoa(averageItemPrice)

	return "Campaign C1 info; Status " + status.String() + ", Target Sales " + sTargetSales + ", Total Sales " + sTotalSales +
		", Turnover " + sTurnover + ", Average Item Price " + sAverageItemPrice
}

func GetCampaignStatus(c Campaign, t time.Time) CampaignStatus {
	timeHour := t.Hour()

	if c.duration > timeHour {
		return Active
	}
	return Ended
}

func (cl *CampaignList) FindCampaignByName(campaignName string) Campaign {
	for _, campaign := range *cl {
		if campaign.name == campaignName {
			return campaign
		}
	}
	return Campaign{}
}
