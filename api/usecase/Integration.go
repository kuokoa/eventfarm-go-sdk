package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type Integration struct {
	restClient sdk.RestClientInterface
}

func NewIntegration(restClient sdk.RestClientInterface) *Integration {
	return &Integration{restClient}
}

// GET: Queries

// POST: Commands
// @param string EventId

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
