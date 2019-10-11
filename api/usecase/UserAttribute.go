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

type UserAttribute struct {
	restClient rest.RestClientInterface
}

func NewUserAttribute(restClient rest.RestClientInterface) *UserAttribute {
	return &UserAttribute{restClient}
}

// GET: Queries

type ListCustomAttributesForUserParameters struct {
	PoolId       string
	UserId       string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *UserAttribute) ListCustomAttributesForUser(p *ListCustomAttributesForUserParameters) (r *http.Response, err error) {
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
		`/v2/UserAttribute/UseCase/ListCustomAttributesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInfoAttributesForUserParameters struct {
	PoolId       string
	UserId       string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *UserAttribute) ListInfoAttributesForUser(p *ListInfoAttributesForUserParameters) (r *http.Response, err error) {
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
		`/v2/UserAttribute/UseCase/ListInfoAttributesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type RemoveUserAttributeParameters struct {
	UserAttributeId string
}

func (t *UserAttribute) RemoveUserAttribute(p *RemoveUserAttributeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userAttributeId`, p.UserAttributeId)

	return t.restClient.Post(
		`/v2/UserAttribute/UseCase/RemoveUserAttribute`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetCustomUserAttributeParameters struct {
	PoolId         string
	UserId         string
	AttributeKey   string
	AttributeValue string
}

func (t *UserAttribute) SetCustomUserAttribute(p *SetCustomUserAttributeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`attributeKey`, p.AttributeKey)
	queryParameters.Add(`attributeValue`, p.AttributeValue)

	return t.restClient.Post(
		`/v2/UserAttribute/UseCase/SetCustomUserAttribute`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetInfoUserAttributeParameters struct {
	PoolId         string
	UserId         string
	AttributeKey   string
	AttributeValue string
}

func (t *UserAttribute) SetInfoUserAttribute(p *SetInfoUserAttributeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`attributeKey`, p.AttributeKey)
	queryParameters.Add(`attributeValue`, p.AttributeValue)

	return t.restClient.Post(
		`/v2/UserAttribute/UseCase/SetInfoUserAttribute`,
		&queryParameters,
		nil,
		nil,
	)
}
