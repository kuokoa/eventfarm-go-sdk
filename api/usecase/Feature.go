package usecase

import (
	"net/http"
	"net/url"

	"bitbucket.ef.network/go/sdk"
)

type Feature struct {
	restClient sdk.RestClientInterface
}

func NewFeature(restClient sdk.RestClientInterface) *Feature {
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
