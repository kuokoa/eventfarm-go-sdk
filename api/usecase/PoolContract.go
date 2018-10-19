package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = fmt.Sprint("Foo")
var _ = http.NoBody

type PoolContract struct {
	restClient sdk.RestClientInterface
}

func NewPoolContract(restClient sdk.RestClientInterface) *PoolContract {
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
