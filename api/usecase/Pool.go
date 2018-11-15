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

type Pool struct {
	restClient rest.RestClientInterface
}

func NewPool(restClient rest.RestClientInterface) *Pool {
	return &Pool{restClient}
}

// GET: Queries
// @param string PoolId

type GetPoolParameters struct {
	PoolId string
}

func (t *Pool) GetPool(p *GetPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/GetPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type GetPoolContractParameters struct {
	PoolId string
}

func (t *Pool) GetPoolContract(p *GetPoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/GetPoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
// @param string|null SortBy
// @param string|null SortDirection ascending|descending

type ListAccessiblePoolsForUserParameters struct {
	UserId        string
	Page          *int
	ItemsPerPage  *int
	SortBy        *string
	SortDirection *string
}

func (t *Pool) ListAccessiblePoolsForUser(p *ListAccessiblePoolsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListAccessiblePoolsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListPoolAllotmentsForPoolParameters struct {
	PoolId        string
	SortBy        *string
	SortDirection *string
	Page          *int
	ItemsPerPage  *int
}

func (t *Pool) ListPoolAllotmentsForPool(p *ListPoolAllotmentsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolAllotmentsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500

type ListPoolContactsByPoolIdParameters struct {
	PoolId       string
	Page         *int
	ItemsPerPage *int
}

func (t *Pool) ListPoolContactsByPoolId(p *ListPoolContactsByPoolIdParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolContactsByPoolId`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string|null Name
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListPoolsParameters struct {
	Name          *string
	SortBy        *string
	SortDirection *string
	Page          *int
	ItemsPerPage  *int
}

func (t *Pool) ListPools(p *ListPoolsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPools`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string|null SortBy name|slug
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500

type ListTagsForPoolParameters struct {
	PoolId        string
	SortBy        *string
	SortDirection *string
	Page          *int
	ItemsPerPage  *int
}

func (t *Pool) ListTagsForPool(p *ListTagsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListTagsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PoolId
// @param string PoolContractType cio|intro|trial|pro|premier|premierPlus|custom
// @param int StartDate
// @param int EndDate
// @param string|null PoolContractId

type CreatePoolContractParameters struct {
	PoolId           string
	PoolContractType string
	StartDate        int
	EndDate          int
	PoolContractId   *string
}

func (t *Pool) CreatePoolContract(p *CreatePoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`poolContractType`, p.PoolContractType)
	queryParameters.Add(`startDate`, strconv.Itoa(p.StartDate))
	queryParameters.Add(`endDate`, strconv.Itoa(p.EndDate))
	if p.PoolContractId != nil {
		queryParameters.Add(`poolContractId`, *p.PoolContractId)
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/CreatePoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string PoolWebhookType
// @param string Url

type CreatePoolWebhookParameters struct {
	PoolId          string
	PoolWebhookType string
	Url             string
}

func (t *Pool) CreatePoolWebhook(p *CreatePoolWebhookParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`poolWebhookType`, p.PoolWebhookType)
	queryParameters.Add(`url`, p.Url)

	return t.restClient.Post(
		`/v2/Pool/UseCase/CreatePoolWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type RemoveCustomerDisplayNameForPoolParameters struct {
	PoolId string
}

func (t *Pool) RemoveCustomerDisplayNameForPool(p *RemoveCustomerDisplayNameForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/RemoveCustomerDisplayNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type RemovePrivacyPolicyLinkForPoolParameters struct {
	PoolId string
}

func (t *Pool) RemovePrivacyPolicyLinkForPool(p *RemovePrivacyPolicyLinkForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/Pool/UseCase/RemovePrivacyPolicyLinkForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string UserId
// @param string|null SlackUserId
// @param array|null RequestedFeatureSlugs
// @param string|null Other

type SendUpgradeRequestToCsmParameters struct {
	PoolId                string
	UserId                string
	SlackUserId           *string
	RequestedFeatureSlugs *[]string
	Other                 *string
}

func (t *Pool) SendUpgradeRequestToCsm(p *SendUpgradeRequestToCsmParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.SlackUserId != nil {
		queryParameters.Add(`slackUserId`, *p.SlackUserId)
	}
	if p.RequestedFeatureSlugs != nil {
		for i := range *p.RequestedFeatureSlugs {
			queryParameters.Add(`requestedFeatureSlugs[]`, (*p.RequestedFeatureSlugs)[i])
		}
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/SendUpgradeRequestToCsm`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string CustomerDisplayName

type SetCustomerDisplayNameForPoolParameters struct {
	PoolId              string
	CustomerDisplayName string
}

func (t *Pool) SetCustomerDisplayNameForPool(p *SetCustomerDisplayNameForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`customerDisplayName`, p.CustomerDisplayName)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetCustomerDisplayNameForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string PrivacyPolicyLink

type SetPrivacyPolicyLinkForPoolParameters struct {
	PoolId            string
	PrivacyPolicyLink string
}

func (t *Pool) SetPrivacyPolicyLinkForPool(p *SetPrivacyPolicyLinkForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`privacyPolicyLink`, p.PrivacyPolicyLink)

	return t.restClient.Post(
		`/v2/Pool/UseCase/SetPrivacyPolicyLinkForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolContractId
// @param string|null PoolContractType cio|intro|trial|pro|premier|premierPlus|custom
// @param int|null StartDate
// @param int|null EndDate

type UpdatePoolContractParameters struct {
	PoolContractId   string
	PoolContractType *string
	StartDate        *int
	EndDate          *int
}

func (t *Pool) UpdatePoolContract(p *UpdatePoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolContractId`, p.PoolContractId)
	if p.PoolContractType != nil {
		queryParameters.Add(`poolContractType`, *p.PoolContractType)
	}
	if p.StartDate != nil {
		queryParameters.Add(`startDate`, strconv.Itoa(*p.StartDate))
	}
	if p.EndDate != nil {
		queryParameters.Add(`endDate`, strconv.Itoa(*p.EndDate))
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/UpdatePoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}
