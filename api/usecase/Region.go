package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"bitbucket.ef.network/go/sdk"
)

type Region struct {
	restClient sdk.RestClientInterface
}

func NewRegion(restClient sdk.RestClientInterface) *Region {
	return &Region{restClient}
}

// GET: Queries
// @param string Query
// @param int|null Page
// @param int|null ItemsPerPage 1-200

type ListTimezonesForRegionParameters struct {
	Query        string
	Page         *int
	ItemsPerPage *int
}

func (t *Region) ListTimezonesForRegion(p *ListTimezonesForRegionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`query`, p.Query)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Region/UseCase/ListTimezonesForRegion`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
