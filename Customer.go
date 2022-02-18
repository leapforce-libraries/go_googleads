package googleads

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Customer struct {
	ResourceName         string `json:"resourceName"`
	CallReportingSetting struct {
		CallReportingEnabled           bool   `json:"callReportingEnabled"`
		CallConversionReportingEnabled bool   `json:"callConversionReportingEnabled"`
		CallConversionAction           string `json:"callConversionAction"`
	} `json:"callReportingSetting"`
	ConversionTrackingSetting struct {
		ConversionTrackingId string `json:"conversionTrackingId"`
	} `json:"conversionTrackingSetting"`
	RemarketingSetting struct {
		GoogleGlobalSiteTag string `json:"googleGlobalSiteTag"`
	} `json:"remarketingSetting"`
	PayPerConversionEligibilityFailureReasons []string `json:"payPerConversionEligibilityFailureReasons"`
	ID                                        string   `json:"id"`
	DescriptiveName                           string   `json:"descriptiveName"`
	CurrencyCode                              string   `json:"currencyCode"`
	TimeZone                                  string   `json:"timeZone"`
	AutoTaggingEnabled                        bool     `json:"autoTaggingEnabled"`
	HasPartnersBadge                          bool     `json:"hasPartnersBadge"`
	Manager                                   bool     `json:"manager"`
	TestAccount                               bool     `json:"testAccount"`
}

func (service *Service) GetCustomer(customerID string) (*Customer, *errortools.Error) {
	customer := Customer{}

	headers := make(http.Header)
	headers.Set("developer-token", _developerToken)

	requestConfig := go_http.RequestConfig{
		Method:            http.MethodGet,
		URL:               service.url(fmt.Sprintf("customers/%s", removeHyphens(customerID))),
		ResponseModel:     &customer,
		NonDefaultHeaders: &headers,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &customer, nil
}
