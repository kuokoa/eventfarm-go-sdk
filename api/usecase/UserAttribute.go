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

type UserAttribute struct {
	restClient gosdk.RestClientInterface
}

func NewUserAttribute(restClient gosdk.RestClientInterface) *UserAttribute {
	return &UserAttribute{restClient}
}

// GET: Queries
// @param string PoolId
// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *UserAttribute) ListCustomAttributesForUser(PoolId string, UserId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
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
func (t *UserAttribute) ListInfoAttributesForUser(PoolId string, UserId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
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
		`/v2/UserAttribute/UseCase/ListInfoAttributesForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string UserAttributeId
func (t *UserAttribute) RemoveUserAttribute(UserAttributeId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userAttributeId`, UserAttributeId)

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
func (t *UserAttribute) SetCustomUserAttribute(PoolId string, UserId string, AttributeKey string, AttributeValue string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`attributeKey`, AttributeKey)
	queryParameters.Add(`attributeValue`, AttributeValue)

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
func (t *UserAttribute) SetInfoUserAttribute(PoolId string, UserId string, AttributeKey string, AttributeValue string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`attributeKey`, AttributeKey)
	queryParameters.Add(`attributeValue`, AttributeValue)

	return t.restClient.Post(
		`/v2/UserAttribute/UseCase/SetInfoUserAttribute`,
		&queryParameters,
		nil,
		nil,
	)
}
