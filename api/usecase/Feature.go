/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type Feature struct {
	restClient rest.RestClientInterface
}

func NewFeature(restClient rest.RestClientInterface) *Feature {
	return &Feature{restClient}
}

// GET: Queries

func (t *Feature) ListFeatures() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/Feature/UseCase/ListFeatures`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
