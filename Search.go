package googleads

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type SearchResults struct {
	Results []map[string]map[string]string `json:"results"`
}

type SearchConfig struct {
	CustomerID              string             `json:"-"`
	Query                   string             `json:"query"`
	PageToken               *string            `json:"page_token,omitempty"`
	PageSize                *uint32            `json:"page_size,omitempty"`
	ValidateOnly            *bool              `json:"validate_only,omitempty"`
	ReturnTotalResultsCount *bool              `json:"return_total_results_count,omitempty"`
	SummaryRowSetting       *SummaryRowSetting `json:"summary_row_setting,omitempty"`
}

type SummaryRowSetting string

const (
	SummaryRowSettingUnspecified           SummaryRowSetting = "UNSPECIFIED"
	SummaryRowSettingUnknown               SummaryRowSetting = "UNKNOWN"
	SummaryRowSettingNoSummaryRow          SummaryRowSetting = "NO_SUMMARY_ROW"
	SummaryRowSettingSummaryRowWithResults SummaryRowSetting = "SUMMARY_ROW_WITH_RESULTS"
	SummaryRowSettingSummaryRowOnly        SummaryRowSetting = "SUMMARY_ROW_ONLY"
)

func (service *Service) Search(config *SearchConfig) (*SearchResults, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("SearchConfig is nil")
	}

	searchResults := SearchResults{}

	headers := make(http.Header)
	headers.Set("developer-token", service.developerToken)

	requestConfig := go_http.RequestConfig{
		Method:            http.MethodPost,
		URL:               service.url(fmt.Sprintf("customers/%s/googleAds:search", removeHyphens(config.CustomerID))),
		BodyModel:         config,
		ResponseModel:     &searchResults,
		NonDefaultHeaders: &headers,
	}
	_, _, e := service.googleService.HTTPRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &searchResults, nil
}
