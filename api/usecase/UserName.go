/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk/rest"
)

type UserName struct {
	restClient rest.RestClientInterface
}

func NewUserName(restClient rest.RestClientInterface) *UserName {
	return &UserName{restClient}
}

// GET: Queries

type GetUserNameParameters struct {
	UserNameId string
}

func (t *UserName) GetUserName(p *GetUserNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userNameId`, p.UserNameId)

	return t.restClient.Get(
		`/v2/UserName/UseCase/GetUserName`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListNamesForUserParameters struct {
	PoolId       string
	UserId       string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *UserName) ListNamesForUser(p *ListNamesForUserParameters) (r *http.Response, err error) {
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
		`/v2/UserName/UseCase/ListNamesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

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
