package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"bitbucket.ef.network/go/sdk"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type UserAttribute struct {
	restClient sdk.RestClientInterface
}

func NewUserAttribute(restClient sdk.RestClientInterface) *UserAttribute {
	return &UserAttribute{restClient}
}

// GET: Queries
// @param string PoolId
// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListCustomAttributesForUserParameters struct {
	PoolId       string
	UserId       string
	Page         *int
	ItemsPerPage *int
}

func (t *UserAttribute) ListCustomAttributesForUser(p *ListCustomAttributesForUserParameters) (r *http.Response, err error) {
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
		`/v2/UserAttribute/UseCase/ListCustomAttributesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListInfoAttributesForUserParameters struct {
	PoolId       string
	UserId       string
	Page         *int
	ItemsPerPage *int
}

func (t *UserAttribute) ListInfoAttributesForUser(p *ListInfoAttributesForUserParameters) (r *http.Response, err error) {
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
		`/v2/UserAttribute/UseCase/ListInfoAttributesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string UserAttributeId

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

// @param string PoolId
// @param string UserId
// @param string AttributeKey
// @param string AttributeValue

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

// @param string PoolId
// @param string UserId
// @param string AttributeKey company|position|title|telephone
// @param string AttributeValue

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
