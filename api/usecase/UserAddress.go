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

type UserAddress struct {
	restClient gosdk.RestClientInterface
}

func NewUserAddress(restClient gosdk.RestClientInterface) *UserAddress {
	return &UserAddress{restClient}
}

// GET: Queries
// @param string PoolId
// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *UserAddress) ListAddressesForUser(PoolId string, UserId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
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
		`/v2/UserAddress/UseCase/ListAddressesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PoolId
// @param string UserId
// @param string|null Address1
// @param string|null City
// @param string|null State
// @param string|null PostalCode
// @param string|null Country
// @param string|null Address2
// @param string|null UserAddressId
func (t *UserAddress) AddUserAddress(PoolId string, UserId string, Address1 *string, City *string, State *string, PostalCode *string, Country *string, Address2 *string, UserAddressId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	if Address1 != nil {
		queryParameters.Add(`address1`, *Address1)
	}
	if City != nil {
		queryParameters.Add(`city`, *City)
	}
	if State != nil {
		queryParameters.Add(`state`, *State)
	}
	if PostalCode != nil {
		queryParameters.Add(`postalCode`, *PostalCode)
	}
	if Country != nil {
		queryParameters.Add(`country`, *Country)
	}
	if Address2 != nil {
		queryParameters.Add(`address2`, *Address2)
	}
	if UserAddressId != nil {
		queryParameters.Add(`userAddressId`, *UserAddressId)
	}

	return t.restClient.Post(
		`/v2/UserAddress/UseCase/AddUserAddress`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserAddressId
func (t *UserAddress) RemoveUserAddress(UserAddressId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userAddressId`, UserAddressId)

	return t.restClient.Post(
		`/v2/UserAddress/UseCase/RemoveUserAddress`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserAddressId
// @param string|null Address1
// @param string|null City
// @param string|null State
// @param string|null PostalCode
// @param string|null Country
// @param string|null Address2
func (t *UserAddress) SetUserAddress(UserAddressId string, Address1 *string, City *string, State *string, PostalCode *string, Country *string, Address2 *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userAddressId`, UserAddressId)
	if Address1 != nil {
		queryParameters.Add(`address1`, *Address1)
	}
	if City != nil {
		queryParameters.Add(`city`, *City)
	}
	if State != nil {
		queryParameters.Add(`state`, *State)
	}
	if PostalCode != nil {
		queryParameters.Add(`postalCode`, *PostalCode)
	}
	if Country != nil {
		queryParameters.Add(`country`, *Country)
	}
	if Address2 != nil {
		queryParameters.Add(`address2`, *Address2)
	}

	return t.restClient.Post(
		`/v2/UserAddress/UseCase/SetUserAddress`,
		&queryParameters,
		nil,
		nil,
	)
}
