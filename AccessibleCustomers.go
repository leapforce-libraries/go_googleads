package googleads

import (
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

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
	_, _, e := service.httpRequest(&requestConfig, nil)
	if e != nil {
		return nil, e
	}

	return &accessibleCustomers, nil
}
