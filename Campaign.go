package googleads

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Campaign struct {
	ResourceName                string `json:"resourceName"`
	Status                      string `json:"status"`
	AdServingOptimizationStatus string `json:"adServingOptimizationStatus"`
	AdvertisingChannelType      string `json:"advertisingChannelType"`
	NetworkSettings             struct {
		TargetGoogleSearch         bool `json:"targetGoogleSearch"`
		TargetSearchNetwork        bool `json:"targetSearchNetwork"`
		TargetContentNetwork       bool `json:"targetContentNetwork"`
		TargetPartnerSearchNetwork bool `json:"targetPartnerSearchNetwork"`
	} `json:"networkSettings"`
	ExperimentType      string `json:"experimentType"`
	ServingStatus       string `json:"servingStatus"`
	BiddingStrategyType string `json:"biddingStrategyType"`
	TargetingSetting    struct {
		TargetRestrictions []struct {
			TargetingDimension string `json:"targetingDimension"`
			BidOnly            bool   `json:"bidOnly"`
		} `json:"targetRestrictions"`
	} `json:"targetingSetting"`
	SelectiveOptimization struct {
		ConversionActions []string `json:"conversionActions"`
	} `json:"selectiveOptimization"`
	GeoTargetTypeSetting struct {
		PositiveGeoTargetType string `json:"positiveGeoTargetType"`
		NegativeGeoTargetType string `json:"negativeGeoTargetType"`
	} `json:"geoTargetTypeSetting"`
	PaymentMode       string   `json:"paymentMode"`
	BaseCampaign      string   `json:"baseCampaign"`
	Name              string   `json:"name"`
	Id                string   `json:"id"`
	CampaignBudget    string   `json:"campaignBudget"`
	StartDate         string   `json:"startDate"`
	EndDate           string   `json:"endDate"`
	OptimizationScore *float64 `json:"optimizationScore"`
}

func (service *Service) GetCampaign(customerId string, campaignId string) (*Campaign, *errortools.Error) {
	campaign := Campaign{}
	customerId = removeHyphens(customerId)

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("customers/%s/campaigns/%s", customerId, campaignId)),
		ResponseModel: &campaign,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &campaign, nil
}
