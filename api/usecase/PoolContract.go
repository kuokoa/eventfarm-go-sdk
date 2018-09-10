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

type PoolContract struct {
	restClient gosdk.RestClientInterface
}

func NewPoolContract(restClient gosdk.RestClientInterface) *PoolContract {
	return &PoolContract{restClient}
}

// GET: Queries
// @param string PoolId
func (t *PoolContract) CountPoolContractUsers(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/PoolContract/UseCase/CountPoolContractUsers`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *PoolContract) GetEmailCountsForPoolContract(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/PoolContract/UseCase/GetEmailCountsForPoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
