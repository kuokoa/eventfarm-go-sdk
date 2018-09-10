package usecase

import (
	"gosdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type Integration struct {
	restClient gosdk.RestClientInterface
}

func NewIntegration(restClient gosdk.RestClientInterface) *Integration {
	return &Integration{restClient}
}

// GET: Queries

// POST: Commands
// @param string EventId
func (t *Integration) InitializeSalesforceIntegrationsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Integration/UseCase/InitializeSalesforceIntegrationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}
