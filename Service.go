package googleads

import (
	"fmt"
	"net/http"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName string = "GoogleAds"
	apiUrl  string = "https://googleads.googleapis.com/v16"
)

var _developerToken string
var _loginCustomerId *string = nil

type Service struct {
	googleService       *google.Service
	accessibleCustomers map[string]bool
}

func NewServiceWithOAuth2(cfg *google.ServiceWithOAuth2Config, developerToken string, loginCustomerId *string) (*Service, *errortools.Error) {
	_developerToken = developerToken
	_loginCustomerId = loginCustomerId

	googleService, e := google.NewServiceWithOAuth2(cfg)
	if e != nil {
		return nil, e
	}

	service := Service{
		googleService:       googleService,
		accessibleCustomers: make(map[string]bool),
	}

	accessibleCustomers, e := service.listAccessibleCustomers()
	if e != nil {
		return nil, e
	}

	if accessibleCustomers != nil {
		for _, customer := range accessibleCustomers.ResourceNames {
			service.accessibleCustomers[customer] = true
		}
	}

	return &service, nil
}

func (service *Service) AccessibleCustomers() []string {
	accessibleCustomers := []string{}

	for c := range service.accessibleCustomers {
		accessibleCustomers = append(accessibleCustomers, c)
	}
	return accessibleCustomers
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig, customerId *string) (*http.Request, *http.Response, *errortools.Error) {
	header := http.Header{}
	header.Set("developer-token", _developerToken)

	isAccessibleCustomer := true
	if customerId != nil {
		_, isAccessibleCustomer = service.accessibleCustomers[fmt.Sprintf("customers/%s", removeHyphens(*customerId))]
	}

	if !isAccessibleCustomer {
		if _loginCustomerId == nil {
			return nil, nil, errortools.ErrorMessage("customer not accessible without login-customer-id")
		}
		header.Set("login-customer-id", removeHyphens(*_loginCustomerId))
	}

	requestConfig.NonDefaultHeaders = &header

	return service.googleService.HttpRequest(requestConfig)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.googleService.ApiKey()
}

func (service *Service) ApiCallCount() int64 {
	return service.googleService.ApiCallCount()
}

func (service *Service) ApiReset() {
	service.googleService.ApiReset()
}

func removeHyphens(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

func (service *Service) ErrorResponse() *google.ErrorResponse {
	return service.googleService.ErrorResponse()
}
