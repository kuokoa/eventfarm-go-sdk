/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"github.com/eventfarm/go-sdk/rest"
	"net/http"
	"net/url"
	"strconv"
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
