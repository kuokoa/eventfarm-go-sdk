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

type ListAccessiblePoolsForUserParameters struct {
	UserId        string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
	SortBy        *string
	SortDirection *string
}

func (t *Pool) ListAccessiblePoolsForUser(p *ListAccessiblePoolsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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

type ListPoolAllotmentsForPoolParameters struct {
	PoolId        string
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolAllotmentsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolContactsByPoolIdParameters struct {
	PoolId       string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-500
}

func (t *Pool) ListPoolContactsByPoolId(p *ListPoolContactsByPoolIdParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPoolContactsByPoolId`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListPoolsParameters struct {
	Name          *string
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListPools`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTagsForPoolParameters struct {
	PoolId        string
	SortBy        *string
	SortDirection *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListTagsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListUniqueTagNamesForPoolParameters struct {
	PoolId string
}

func (t *Pool) ListUniqueTagNamesForPool(p *ListUniqueTagNamesForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Pool/UseCase/ListUniqueTagNamesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePoolContractParameters struct {
	PoolId           string
	PoolContractType string
	StartDate        int64
	EndDate          int64
	PoolContractId   *string
}

func (t *Pool) CreatePoolContract(p *CreatePoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`poolContractType`, p.PoolContractType)
	queryParameters.Add(`startDate`, strconv.FormatInt(p.StartDate, 10))
	queryParameters.Add(`endDate`, strconv.FormatInt(p.EndDate, 10))
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

type UpdatePoolContractParameters struct {
	PoolContractId   string
	PoolContractType *string
	StartDate        *int64
	EndDate          *int64
}

func (t *Pool) UpdatePoolContract(p *UpdatePoolContractParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolContractId`, p.PoolContractId)
	if p.PoolContractType != nil {
		queryParameters.Add(`poolContractType`, *p.PoolContractType)
	}
	if p.StartDate != nil {
		queryParameters.Add(`startDate`, strconv.FormatInt(*p.StartDate, 10))
	}
	if p.EndDate != nil {
		queryParameters.Add(`endDate`, strconv.FormatInt(*p.EndDate, 10))
	}

	return t.restClient.Post(
		`/v2/Pool/UseCase/UpdatePoolContract`,
		&queryParameters,
		nil,
		nil,
	)
}
