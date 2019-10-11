/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type PoolContract struct {
	restClient rest.RestClientInterface
}

func NewPoolContract(restClient rest.RestClientInterface) *PoolContract {
	return &PoolContract{restClient}
}

// GET: Queries

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
