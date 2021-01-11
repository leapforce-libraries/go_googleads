package googleads

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	google "github.com/leapforce-libraries/go_google"
)

const (
	APIName string = "GoogleAds"
	APIURL  string = "https://googleads.googleapis.com/v6"
)

// Service stores Service configuration
//
type Service struct {
	googleService *google.Service
}

type ServiceConfig struct {
	ClientID       string
	ClientSecret   string
	Scope          string
	DeveloperToken string
}

// methods
//
func NewService(serviceConfig *ServiceConfig, bigQuery *google.BigQuery) *Service {
	if serviceConfig == nil {
		return nil
	}

	headers := make(http.Header)
	headers.Set("developer-token", serviceConfig.DeveloperToken)

	googleServiceConfig := google.ServiceConfig{
		APIName:           APIName,
		ClientID:          serviceConfig.ClientID,
		ClientSecret:      serviceConfig.ClientSecret,
		Scope:             serviceConfig.Scope,
		NonDefaultHeaders: &headers,
	}

	googleService := google.NewService(googleServiceConfig, bigQuery)

	return &Service{googleService}
}

func (service *Service) InitToken() *errortools.Error {
	return service.googleService.InitToken()
}
