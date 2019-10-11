/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk/rest"
)

type Withdrawal struct {
	restClient rest.RestClientInterface
}

func NewWithdrawal(restClient rest.RestClientInterface) *Withdrawal {
	return &Withdrawal{restClient}
}

// GET: Queries

type GetWithdrawalParameters struct {
	WithdrawalId string
	WithData     *[]string
}

func (t *Withdrawal) GetWithdrawal(p *GetWithdrawalParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`withdrawalId`, p.WithdrawalId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Withdrawal/UseCase/GetWithdrawal`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListWithdrawalsForEventParameters struct {
	EventId       string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	SortBy        *string
	SortDirection *string
}

func (t *Withdrawal) ListWithdrawalsForEvent(p *ListWithdrawalsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
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
		`/v2/Withdrawal/UseCase/ListWithdrawalsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateWithdrawalParameters struct {
	EventId      string
	Amount       float64
	WithdrawalId *string
}

func (t *Withdrawal) CreateWithdrawal(p *CreateWithdrawalParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`amount`, fmt.Sprintf("%f", p.Amount))
	if p.WithdrawalId != nil {
		queryParameters.Add(`withdrawalId`, *p.WithdrawalId)
	}

	return t.restClient.Post(
		`/v2/Withdrawal/UseCase/CreateWithdrawal`,
		&queryParameters,
		nil,
		nil,
	)
}

type DeleteWithdrawalParameters struct {
	WithdrawalId string
}

func (t *Withdrawal) DeleteWithdrawal(p *DeleteWithdrawalParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`withdrawalId`, p.WithdrawalId)

	return t.restClient.Post(
		`/v2/Withdrawal/UseCase/DeleteWithdrawal`,
		&queryParameters,
		nil,
		nil,
	)
}
