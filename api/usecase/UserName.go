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

type UserName struct {
	restClient gosdk.RestClientInterface
}

func NewUserName(restClient gosdk.RestClientInterface) *UserName {
	return &UserName{restClient}
}

// GET: Queries
// @param string PoolId
// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *UserName) ListNamesForUser(PoolId string, UserId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/UserName/UseCase/ListNamesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PoolId
// @param string UserId
// @param string|null FirstName
// @param string|null LastName
// @param string|null UserNameId
func (t *UserName) AddUserName(PoolId string, UserId string, FirstName *string, LastName *string, UserNameId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	if FirstName != nil {
		queryParameters.Add(`firstName`, *FirstName)
	}
	if LastName != nil {
		queryParameters.Add(`lastName`, *LastName)
	}
	if UserNameId != nil {
		queryParameters.Add(`userNameId`, *UserNameId)
	}

	return t.restClient.Post(
		`/v2/UserName/UseCase/AddUserName`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserNameId
func (t *UserName) RemoveUserName(UserNameId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userNameId`, UserNameId)

	return t.restClient.Post(
		`/v2/UserName/UseCase/RemoveUserName`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserNameId
// @param string|null FirstName
// @param string|null LastName
func (t *UserName) SetUserName(UserNameId string, FirstName *string, LastName *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userNameId`, UserNameId)
	if FirstName != nil {
		queryParameters.Add(`firstName`, *FirstName)
	}
	if LastName != nil {
		queryParameters.Add(`lastName`, *LastName)
	}

	return t.restClient.Post(
		`/v2/UserName/UseCase/SetUserName`,
		&queryParameters,
		nil,
		nil,
	)
}
