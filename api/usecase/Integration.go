/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type Integration struct {
	restClient rest.RestClientInterface
}

func NewIntegration(restClient rest.RestClientInterface) *Integration {
	return &Integration{restClient}
}

// GET: Queries

// POST: Commands

type InitializeSalesforceIntegrationsForEventParameters struct {
	EventId string
}

func (t *Integration) InitializeSalesforceIntegrationsForEvent(p *InitializeSalesforceIntegrationsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Integration/UseCase/InitializeSalesforceIntegrationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Integration) InitializeSalesforceIntegrationsForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Integration/UseCase/InitializeSalesforceIntegrationsForEvent`,
		data,
		nil,
		nil,
	)
}
