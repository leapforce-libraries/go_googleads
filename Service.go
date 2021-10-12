package googleads

import (
	"fmt"
	"strings"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
)

const (
	apiName string = "GoogleAds"
	apiURL  string = "https://googleads.googleapis.com/v8"
)

// Service stores Service configuration
//
type Service struct {
	clientID       string
	developerToken string
	googleService  *google.Service
}

type ServiceConfig struct {
	ClientID       string
	ClientSecret   string
	DeveloperToken string
}

// methods
//
func NewService(serviceConfig *ServiceConfig, bigQueryService *bigquery.Service) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.ClientID == "" {
		return nil, errortools.ErrorMessage("ClientID not provided")
	}

	if serviceConfig.ClientSecret == "" {
		return nil, errortools.ErrorMessage("ClientSecret not provided")
	}

	googleServiceConfig := google.ServiceConfig{
		APIName:      apiName,
		ClientID:     serviceConfig.ClientID,
		ClientSecret: serviceConfig.ClientSecret,
	}

	googleService, e := google.NewService(&googleServiceConfig, bigQueryService)
	if e != nil {
		return nil, e
	}

	return &Service{
		serviceConfig.ClientID,
		serviceConfig.DeveloperToken,
		googleService,
	}, nil
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func (service *Service) InitToken(scope string, accessType *string, prompt *string, state *string) *errortools.Error {
	return service.googleService.InitToken(scope, accessType, prompt, state)
}

func (service *Service) APIName() string {
	return apiName
}

func (service *Service) APIKey() string {
	return service.clientID
}

func (service *Service) APICallCount() int64 {
	return service.googleService.APICallCount()
}

func (service *Service) APIReset() {
	service.googleService.APIReset()
}

func removeHyphens(s string) string {
	return strings.ReplaceAll(s, "-", "")
}
