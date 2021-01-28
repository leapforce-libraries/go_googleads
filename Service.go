package googleads

import (
	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
	bigquery "github.com/leapforce-libraries/go_google/bigquery"
)

const (
	APIName string = "GoogleAds"
	APIURL  string = "https://googleads.googleapis.com/v6"
)

// Service stores Service configuration
//
type Service struct {
	developerToken string
	googleService  *google.Service
}

type ServiceConfig struct {
	ClientID       string
	ClientSecret   string
	Scope          string
	DeveloperToken string
}

// methods
//
func NewService(serviceConfig *ServiceConfig, bigQueryService *bigquery.Service) *Service {
	if serviceConfig == nil {
		return nil
	}

	googleServiceConfig := google.ServiceConfig{
		APIName:      APIName,
		ClientID:     serviceConfig.ClientID,
		ClientSecret: serviceConfig.ClientSecret,
		Scope:        serviceConfig.Scope,
	}

	googleService := google.NewService(googleServiceConfig, bigQueryService)

	return &Service{
		serviceConfig.DeveloperToken,
		googleService,
	}
}

func (service *Service) InitToken() *errortools.Error {
	return service.googleService.InitToken()
}
