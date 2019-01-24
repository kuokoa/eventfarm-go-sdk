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

type UserAddress struct {
	restClient rest.RestClientInterface
}

func NewUserAddress(restClient rest.RestClientInterface) *UserAddress {
	return &UserAddress{restClient}
}

// GET: Queries
// @param string PoolId
// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListAddressesForUserParameters struct {
	PoolId       string
	UserId       string
	Page         *int64
	ItemsPerPage *int64
}

func (t *UserAddress) ListAddressesForUser(p *ListAddressesForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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

type AddUserAddressParameters struct {
	PoolId        string
	UserId        string
	Address1      *string
	City          *string
	State         *string
	PostalCode    *string
	Country       *string
	Address2      *string
	UserAddressId *string
}

func (t *UserAddress) AddUserAddress(p *AddUserAddressParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.Address1 != nil {
		queryParameters.Add(`address1`, *p.Address1)
	}
	if p.City != nil {
		queryParameters.Add(`city`, *p.City)
	}
	if p.State != nil {
		queryParameters.Add(`state`, *p.State)
	}
	if p.PostalCode != nil {
		queryParameters.Add(`postalCode`, *p.PostalCode)
	}
	if p.Country != nil {
		queryParameters.Add(`country`, *p.Country)
	}
	if p.Address2 != nil {
		queryParameters.Add(`address2`, *p.Address2)
	}
	if p.UserAddressId != nil {
		queryParameters.Add(`userAddressId`, *p.UserAddressId)
	}

	return t.restClient.Post(
		`/v2/UserAddress/UseCase/AddUserAddress`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserAddressId

type RemoveUserAddressParameters struct {
	UserAddressId string
}

func (t *UserAddress) RemoveUserAddress(p *RemoveUserAddressParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userAddressId`, p.UserAddressId)

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

type SetUserAddressParameters struct {
	UserAddressId string
	Address1      *string
	City          *string
	State         *string
	PostalCode    *string
	Country       *string
	Address2      *string
}

func (t *UserAddress) SetUserAddress(p *SetUserAddressParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userAddressId`, p.UserAddressId)
	if p.Address1 != nil {
		queryParameters.Add(`address1`, *p.Address1)
	}
	if p.City != nil {
		queryParameters.Add(`city`, *p.City)
	}
	if p.State != nil {
		queryParameters.Add(`state`, *p.State)
	}
	if p.PostalCode != nil {
		queryParameters.Add(`postalCode`, *p.PostalCode)
	}
	if p.Country != nil {
		queryParameters.Add(`country`, *p.Country)
	}
	if p.Address2 != nil {
		queryParameters.Add(`address2`, *p.Address2)
	}

	return t.restClient.Post(
		`/v2/UserAddress/UseCase/SetUserAddress`,
		&queryParameters,
		nil,
		nil,
	)
}
