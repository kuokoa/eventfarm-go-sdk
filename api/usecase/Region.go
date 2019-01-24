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

type Region struct {
	restClient rest.RestClientInterface
}

func NewRegion(restClient rest.RestClientInterface) *Region {
	return &Region{restClient}
}

// GET: Queries
// @param string Query
// @param int|null Page
// @param int|null ItemsPerPage 1-200

type ListTimezonesForRegionParameters struct {
	Query        string
	Page         *int64
	ItemsPerPage *int64
}

func (t *Region) ListTimezonesForRegion(p *ListTimezonesForRegionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`query`, p.Query)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Region/UseCase/ListTimezonesForRegion`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
