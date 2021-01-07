package googleads

import (
	"net/http"

	google "github.com/leapforce-libraries/go_google"
)

const (
	APIName string = "GoogleAds"
	APIURL  string = "https://googleads.googleapis.com/v6"
)

// GoogleAds stores GoogleAds configuration
//
type GoogleAds struct {
	Client *google.GoogleClient
}

type GoogleAdsConfig struct {
	ClientID       string
	ClientSecret   string
	Scope          string
	DeveloperToken string
}

// methods
//
func NewGoogleAds(googleAdsConfig *GoogleAdsConfig, bigQuery *google.BigQuery) *GoogleAds {
	if googleAdsConfig == nil {
		return nil
	}

	headers := make(http.Header)
	headers.Set("developer-token", googleAdsConfig.DeveloperToken)

	googleClientConfig := google.GoogleClientConfig{
		APIName:           APIName,
		ClientID:          googleAdsConfig.ClientID,
		ClientSecret:      googleAdsConfig.ClientSecret,
		Scope:             googleAdsConfig.Scope,
		NonDefaultHeaders: &headers,
	}

	googleClient := google.NewGoogleClient(googleClientConfig, bigQuery)

	return &GoogleAds{googleClient}
}
