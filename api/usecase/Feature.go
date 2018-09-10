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

type Feature struct {
	restClient gosdk.RestClientInterface
}

func NewFeature(restClient gosdk.RestClientInterface) *Feature {
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
