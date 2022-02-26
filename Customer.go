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
	Id                                        string   `json:"id"`
	DescriptiveName                           string   `json:"descriptiveName"`
	CurrencyCode                              string   `json:"currencyCode"`
	TimeZone                                  string   `json:"timeZone"`
	AutoTaggingEnabled                        bool     `json:"autoTaggingEnabled"`
	HasPartnersBadge                          bool     `json:"hasPartnersBadge"`
	Manager                                   bool     `json:"manager"`
	TestAccount                               bool     `json:"testAccount"`
}

func (service *Service) GetCustomer(customerId string) (*Customer, *errortools.Error) {
	customer := Customer{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("customers/%s", removeHyphens(customerId))),
		ResponseModel: &customer,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &customer, nil
}

type AccessibleCustomers struct {
	ResourceNames []string `json:"resourceNames"`
}

func (service *Service) ListAccessibleCustomers() (*AccessibleCustomers, *errortools.Error) {
	accessibleCustomers := AccessibleCustomers{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("customers:listAccessibleCustomers"),
		ResponseModel: &accessibleCustomers,
	}
	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &accessibleCustomers, nil
}
