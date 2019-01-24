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

type PoolContract struct {
	restClient rest.RestClientInterface
}

func NewPoolContract(restClient rest.RestClientInterface) *PoolContract {
	return &PoolContract{restClient}
}

// GET: Queries
// @param string PoolId

type CountPoolContractUsersParameters struct {
	PoolId string
}

func (t *PoolContract) CountPoolContractUsers(p *CountPoolContractUsersParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/PoolContract/UseCase/CountPoolContractUsers`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type GetEmailCountsForPoolContractParameters struct {
	PoolId string
}

func (t *PoolContract) GetEmailCountsForPoolContract(p *GetEmailCountsForPoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/PoolContract/UseCase/GetEmailCountsForPoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
