/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type Canvas struct {
	restClient rest.RestClientInterface
}

func NewCanvas(restClient rest.RestClientInterface) *Canvas {
	return &Canvas{restClient}
}

// GET: Queries

// POST: Commands

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

func (t *Canvas) EnableCanvasForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Canvas/UseCase/EnableCanvasForEvent`,
		data,
		nil,
		nil,
	)
}
