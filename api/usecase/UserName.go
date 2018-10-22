/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk"
)

type UserName struct {
	restClient sdk.RestClientInterface
}

func NewUserName(restClient sdk.RestClientInterface) *UserName {
	return &UserName{restClient}
}

// GET: Queries
// @param string PoolId
// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListNamesForUserParameters struct {
	PoolId       string
	UserId       string
	Page         *int
	ItemsPerPage *int
}

func (t *UserName) ListNamesForUser(p *ListNamesForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
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

type AddUserNameParameters struct {
	PoolId     string
	UserId     string
	FirstName  *string
	LastName   *string
	UserNameId *string
}

func (t *UserName) AddUserName(p *AddUserNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}
	if p.UserNameId != nil {
		queryParameters.Add(`userNameId`, *p.UserNameId)
	}

	return t.restClient.Post(
		`/v2/UserName/UseCase/AddUserName`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserNameId

type RemoveUserNameParameters struct {
	UserNameId string
}

func (t *UserName) RemoveUserName(p *RemoveUserNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userNameId`, p.UserNameId)

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

type SetUserNameParameters struct {
	UserNameId string
	FirstName  *string
	LastName   *string
}

func (t *UserName) SetUserName(p *SetUserNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userNameId`, p.UserNameId)
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}

	return t.restClient.Post(
		`/v2/UserName/UseCase/SetUserName`,
		&queryParameters,
		nil,
		nil,
	)
}
