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

type Region struct {
	restClient gosdk.RestClientInterface
}

func NewRegion(restClient gosdk.RestClientInterface) *Region {
	return &Region{restClient}
}

// GET: Queries
// @param string Query
// @param int|null Page
// @param int|null ItemsPerPage 1-200
func (t *Region) ListTimezonesForRegion(Query string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`query`, Query)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Region/UseCase/ListTimezonesForRegion`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
