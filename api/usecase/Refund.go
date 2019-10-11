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

type Refund struct {
	restClient rest.RestClientInterface
}

func NewRefund(restClient rest.RestClientInterface) *Refund {
	return &Refund{restClient}
}

// GET: Queries

type GetRefundParameters struct {
	RefundId string
	WithData *[]string
}

func (t *Refund) GetRefund(p *GetRefundParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`refundId`, p.RefundId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Refund/UseCase/GetRefund`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListRefundsForPaymentParameters struct {
	PaymentId     string
	WithData      *[]string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	SortBy        *string
	SortDirection *string
}

func (t *Refund) ListRefundsForPayment(p *ListRefundsForPaymentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`paymentId`, p.PaymentId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
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
		`/v2/Refund/UseCase/ListRefundsForPayment`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateRefundParameters struct {
	PaymentId     string
	RefundTransId string
	Amount        float64
	RefundId      *string
}

func (t *Refund) CreateRefund(p *CreateRefundParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`paymentId`, p.PaymentId)
	queryParameters.Add(`refundTransId`, p.RefundTransId)
	queryParameters.Add(`amount`, fmt.Sprintf("%f", p.Amount))
	if p.RefundId != nil {
		queryParameters.Add(`refundId`, *p.RefundId)
	}

	return t.restClient.Post(
		`/v2/Refund/UseCase/CreateRefund`,
		&queryParameters,
		nil,
		nil,
	)
}

type DeleteRefundParameters struct {
	RefundId string
}

func (t *Refund) DeleteRefund(p *DeleteRefundParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`refundId`, p.RefundId)

	return t.restClient.Post(
		`/v2/Refund/UseCase/DeleteRefund`,
		&queryParameters,
		nil,
		nil,
	)
}
