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

type UserIdentifier struct {
	restClient gosdk.RestClientInterface
}

func NewUserIdentifier(restClient gosdk.RestClientInterface) *UserIdentifier {
	return &UserIdentifier{restClient}
}

// GET: Queries

// POST: Commands
// @param string UserIdentifierId
// @param string Identifier
// @param string|null PoolId
func (t *UserIdentifier) SetUserIdentifier(UserIdentifierId string, Identifier string, PoolId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userIdentifierId`, UserIdentifierId)
	queryParameters.Add(`identifier`, Identifier)
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}

	return t.restClient.Post(
		`/v2/UserIdentifier/UseCase/SetUserIdentifier`,
		&queryParameters,
		nil,
		nil,
	)
}
