package model

import (
	"testing"
)

func TestCreateCampaign(t *testing.T) {

}

func TestAddCampaign(t *testing.T) {
	cl := CampaignList{}
	beforeLen := len(cl)
	c := Campaign{}
	cl.AddCampaign(c)

	if len(cl) != (beforeLen + 1) {
		t.Errorf("Campaign list couldn't added to list. List Length: %d", len(cl))
	}
}

func TestCreateCampaignCommand(t *testing.T) {

}
func TestGetCampaignInfoCommand(t *testing.T) {

}

func TestGetCampaignInfo(t *testing.T) {

}

func TestGetCampaignStatus(t *testing.T) {

}
