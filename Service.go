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
	apiUrl  string = "https://googleads.googleapis.com/v10"
)

var _developerToken string
var _loginCustomerId *string = nil

type Service google.Service

func NewServiceWithOAuth2(cfg *google.ServiceWithOAuth2Config, developerToken string, loginCustomerId *string) (*Service, *errortools.Error) {
	_developerToken = developerToken
	_loginCustomerId = loginCustomerId

	googleService, e := google.NewServiceWithOAuth2(cfg)
	if e != nil {
		return nil, e
	}
	service := Service(*googleService)
	return &service, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	header := http.Header{}
	header.Set("developer-token", _developerToken)

	if _loginCustomerId != nil {
		header.Set("login-customer-id", removeHyphens(*_loginCustomerId))
	}

	requestConfig.NonDefaultHeaders = &header

	return service.googleService().HttpRequest(requestConfig)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.googleService().ApiKey()
}

func (service *Service) ApiCallCount() int64 {
	return service.googleService().ApiCallCount()
}

func (service *Service) ApiReset() {
	service.googleService().ApiReset()
}

func (service *Service) googleService() *google.Service {
	googleService := google.Service(*service)
	return &googleService
}

func removeHyphens(s string) string {
	return strings.ReplaceAll(s, "-", "")
}
