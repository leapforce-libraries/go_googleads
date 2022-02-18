package googleads

import (
	"fmt"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	apiName string = "GoogleAds"
	apiURL  string = "https://googleads.googleapis.com/v9"
)

var _developerToken string

type Service google.Service

func NewServiceWithAccessToken(cfg *google.ServiceWithAccessTokenConfig, developerToken string) (*Service, *errortools.Error) {
	_developerToken = developerToken

	googleService, e := google.NewServiceWithAccessToken(cfg)
	if e != nil {
		return nil, e
	}
	service := Service(*googleService)
	return &service, nil
}

func NewServiceWithApiKey(cfg *google.ServiceWithApiKeyConfig, developerToken string) (*Service, *errortools.Error) {
	_developerToken = developerToken

	googleService, e := google.NewServiceWithApiKey(cfg)
	if e != nil {
		return nil, e
	}
	service := Service(*googleService)
	return &service, nil
}

func NewServiceWithOAuth2(cfg *google.ServiceWithOAuth2Config, developerToken string) (*Service, *errortools.Error) {
	_developerToken = developerToken

	googleService, e := google.NewServiceWithOAuth2(cfg)
	if e != nil {
		return nil, e
	}
	service := Service(*googleService)
	return &service, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
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
