/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk"
)

type Canvas struct {
	restClient sdk.RestClientInterface
}

func NewCanvas(restClient sdk.RestClientInterface) *Canvas {
	return &Canvas{restClient}
}

// GET: Queries

// POST: Commands
// @param string EventId

type EnableCanvasForEventParameters struct {
	EventId string
}

func (t *Canvas) EnableCanvasForEvent(p *EnableCanvasForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Canvas/UseCase/EnableCanvasForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}
