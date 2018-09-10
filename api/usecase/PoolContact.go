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

type PoolContact struct {
	restClient gosdk.RestClientInterface
}

func NewPoolContact(restClient gosdk.RestClientInterface) *PoolContact {
	return &PoolContact{restClient}
}

// GET: Queries
// @param string UserId
// @param string|null PoolId
// @param string|null PoolContactType full|create
func (t *PoolContact) ListPoolContactsForUser(UserId string, PoolId *string, PoolContactType *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if PoolContactType != nil {
		queryParameters.Add(`poolContactType`, *PoolContactType)
	}

	return t.restClient.Get(
		`/v2/PoolContact/UseCase/ListPoolContactsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
