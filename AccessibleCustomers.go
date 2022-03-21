package googleads

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type AccessibleCustomers struct {
	ResourceNames []string `json:"resourceNames"`
}

func (service *Service) listAccessibleCustomers() (*AccessibleCustomers, *errortools.Error) {
	accessibleCustomers := AccessibleCustomers{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url("customers:listAccessibleCustomers"),
		ResponseModel: &accessibleCustomers,
	}
	_, _, e := service.httpRequest(&requestConfig, nil)
	if e != nil {
		errorResponse := service.googleService.ErrorResponse()
		if errorResponse != nil {
			if errorResponse.Error.Code == 401 {
				for _, detail := range errorResponse.Error.Details {
					for _, err := range detail.Errors {
						if err.ErrorCode["authenticationError"] == "NOT_ADS_USER" {
							e.SetMessage("This account is not a GoogleAds user")
						}
					}
				}
			}
		}

		return nil, e
	}

	return &accessibleCustomers, nil
}
