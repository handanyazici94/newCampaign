package main

import (
	"new-campaign/io"
	"new-campaign/model"
	"new-campaign/util"
	"os"
	"strconv"
)

const (
	createProductCommand   = "create_product"
	createCampaignCommand  = "create_campaign"
	createOrderCommand     = "create_order"
	getProductInfoCommand  = "get_product_info"
	getCampaignInfoCommand = "get_campaign_info"
	increaseTimeCommand    = "increase_time"
)

func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	util.GetSystemTime()
	commands := io.ReadCommandsFromFile(inputFile)
	processCommands(outputFile, commands)

}

func processCommands(outputFile string, commands [][]string) {

	campaignSalesMap := make(map[model.Campaign]model.CampaignSalesList)
	productList := model.ProductList{}
	campaignList := model.CampaignList{}
	output := "null"
	time := util.GetSystemTime()

	for _, command := range commands {
		if command[0] == createProductCommand {
			stock, _ := strconv.Atoi(command[3])
			price, _ := strconv.ParseFloat(command[2], 64)
			output := productList.StartCreateProductCommand(command[1], price, stock)
			io.CreateAndWriteOutputFile(outputFile, output)
		} else if command[0] == createCampaignCommand {
			targetSales, _ := strconv.Atoi(command[5])
			duration, _ := strconv.Atoi(command[3])
			limit, _ := strconv.Atoi(command[4])
			output = campaignList.StartCreateCampaignCommand(command[1], command[2], duration, limit, targetSales, productList)
			io.CreateAndWriteOutputFile(outputFile, output)
		} else if command[0] == createOrderCommand {
			productCode := command[1]
			quantity, _ := strconv.Atoi(command[2])
			output = model.StartCreateOrderCommand(productCode, quantity, productList, campaignList, time, campaignSalesMap)
			io.CreateAndWriteOutputFile(outputFile, output)
		} else if command[0] == increaseTimeCommand {
			count, _ := strconv.Atoi(command[1])
			time = util.IncreaseTime(time, count)
			output = util.IncreaseTimeCommand(time)
			io.CreateAndWriteOutputFile(outputFile, output)
		} else if command[0] == getProductInfoCommand {
			output = model.StartProductInfoCommand(command[1], productList, campaignList, time)
			io.CreateAndWriteOutputFile(outputFile, output)
		} else if command[0] == getCampaignInfoCommand {
			output = model.StartCampaignInfoCommand(command[1], campaignList, time, campaignSalesMap)
			io.CreateAndWriteOutputFile(outputFile, output)
		}
	}

}
