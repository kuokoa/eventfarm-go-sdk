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

type Transaction struct {
	restClient rest.RestClientInterface
}

func NewTransaction(restClient rest.RestClientInterface) *Transaction {
	return &Transaction{restClient}
}

// GET: Queries

type GetTransactionParameters struct {
	TransactionId string
	WithData      *[]string
}

func (t *Transaction) GetTransaction(p *GetTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transactionId`, p.TransactionId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Transaction/UseCase/GetTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetTransactionTotalsForEventParameters struct {
	EventId string
}

func (t *Transaction) GetTransactionTotalsForEvent(p *GetTransactionTotalsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Transaction/UseCase/GetTransactionTotalsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTransactionsForEventParameters struct {
	EventId       string
	WithData      *[]string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	SortBy        *string
	SortDirection *string
}

func (t *Transaction) ListTransactionsForEvent(p *ListTransactionsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
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
		`/v2/Transaction/UseCase/ListTransactionsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateTransactionParameters struct {
	EventId              string
	Type                 string
	TransType            string
	BillingFirstName     *string
	BillingLastName      *string
	BillingAddress1      *string
	BillingAddress2      *string
	BillingCity          *string
	BillingState         *string
	BillingPostalCode    *string
	BillingCountry       *string
	BillingEmail         *string
	BillingPhone         *string
	ShippingFirstName    *string
	ShippingLastName     *string
	ShippingAddress1     *string
	ShippingAddress2     *string
	ShippingCity         *string
	ShippingState        *string
	ShippingPostalCode   *string
	ShippingCountry      *string
	CardType             *string
	CardNumber           *string
	CardExpMon           *string
	CardExpYear          *string
	ProcessingCurrency   *string
	Amount               *float64
	DonationAmount       *float64
	ShippingAmount       *float64
	PromotionAmount      *float64
	TransId              *string
	TransResult          *string
	TransOrderRef        *string
	TransMerchantProfile *string
	Memo                 *string
	InvoiceNo            *string
	IsOutside            *bool
	TransactionId        *string
}

func (t *Transaction) CreateTransaction(p *CreateTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`type`, p.Type)
	queryParameters.Add(`transType`, p.TransType)
	if p.BillingFirstName != nil {
		queryParameters.Add(`billingFirstName`, *p.BillingFirstName)
	}
	if p.BillingLastName != nil {
		queryParameters.Add(`billingLastName`, *p.BillingLastName)
	}
	if p.BillingAddress1 != nil {
		queryParameters.Add(`billingAddress1`, *p.BillingAddress1)
	}
	if p.BillingAddress2 != nil {
		queryParameters.Add(`billingAddress2`, *p.BillingAddress2)
	}
	if p.BillingCity != nil {
		queryParameters.Add(`billingCity`, *p.BillingCity)
	}
	if p.BillingState != nil {
		queryParameters.Add(`billingState`, *p.BillingState)
	}
	if p.BillingPostalCode != nil {
		queryParameters.Add(`billingPostalCode`, *p.BillingPostalCode)
	}
	if p.BillingCountry != nil {
		queryParameters.Add(`billingCountry`, *p.BillingCountry)
	}
	if p.BillingEmail != nil {
		queryParameters.Add(`billingEmail`, *p.BillingEmail)
	}
	if p.BillingPhone != nil {
		queryParameters.Add(`billingPhone`, *p.BillingPhone)
	}
	if p.ShippingFirstName != nil {
		queryParameters.Add(`shippingFirstName`, *p.ShippingFirstName)
	}
	if p.ShippingLastName != nil {
		queryParameters.Add(`shippingLastName`, *p.ShippingLastName)
	}
	if p.ShippingAddress1 != nil {
		queryParameters.Add(`shippingAddress1`, *p.ShippingAddress1)
	}
	if p.ShippingAddress2 != nil {
		queryParameters.Add(`shippingAddress2`, *p.ShippingAddress2)
	}
	if p.ShippingCity != nil {
		queryParameters.Add(`shippingCity`, *p.ShippingCity)
	}
	if p.ShippingState != nil {
		queryParameters.Add(`shippingState`, *p.ShippingState)
	}
	if p.ShippingPostalCode != nil {
		queryParameters.Add(`shippingPostalCode`, *p.ShippingPostalCode)
	}
	if p.ShippingCountry != nil {
		queryParameters.Add(`shippingCountry`, *p.ShippingCountry)
	}
	if p.CardType != nil {
		queryParameters.Add(`cardType`, *p.CardType)
	}
	if p.CardNumber != nil {
		queryParameters.Add(`cardNumber`, *p.CardNumber)
	}
	if p.CardExpMon != nil {
		queryParameters.Add(`cardExpMon`, *p.CardExpMon)
	}
	if p.CardExpYear != nil {
		queryParameters.Add(`cardExpYear`, *p.CardExpYear)
	}
	if p.ProcessingCurrency != nil {
		queryParameters.Add(`processingCurrency`, *p.ProcessingCurrency)
	}
	if p.Amount != nil {
		queryParameters.Add(`amount`, fmt.Sprintf("%f", *p.Amount))
	}
	if p.DonationAmount != nil {
		queryParameters.Add(`donationAmount`, fmt.Sprintf("%f", *p.DonationAmount))
	}
	if p.ShippingAmount != nil {
		queryParameters.Add(`shippingAmount`, fmt.Sprintf("%f", *p.ShippingAmount))
	}
	if p.PromotionAmount != nil {
		queryParameters.Add(`promotionAmount`, fmt.Sprintf("%f", *p.PromotionAmount))
	}
	if p.TransId != nil {
		queryParameters.Add(`transId`, *p.TransId)
	}
	if p.TransResult != nil {
		queryParameters.Add(`transResult`, *p.TransResult)
	}
	if p.TransOrderRef != nil {
		queryParameters.Add(`transOrderRef`, *p.TransOrderRef)
	}
	if p.TransMerchantProfile != nil {
		queryParameters.Add(`transMerchantProfile`, *p.TransMerchantProfile)
	}
	if p.Memo != nil {
		queryParameters.Add(`memo`, *p.Memo)
	}
	if p.InvoiceNo != nil {
		queryParameters.Add(`invoiceNo`, *p.InvoiceNo)
	}
	if p.IsOutside != nil {
		queryParameters.Add(`isOutside`, strconv.FormatBool(*p.IsOutside))
	}
	if p.TransactionId != nil {
		queryParameters.Add(`transactionId`, *p.TransactionId)
	}

	return t.restClient.Post(
		`/v2/Transaction/UseCase/CreateTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}

type DeleteTransactionParameters struct {
	TransactionId string
}

func (t *Transaction) DeleteTransaction(p *DeleteTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transactionId`, p.TransactionId)

	return t.restClient.Post(
		`/v2/Transaction/UseCase/DeleteTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}

type ProcessTransactionParameters struct {
	EventId     string
	GuestsJson  []string
	PaymentJson []string
}

func (t *Transaction) ProcessTransaction(p *ProcessTransactionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	for i := range p.GuestsJson {
		queryParameters.Add(`guestsJson[]`, p.GuestsJson[i])
	}
	for i := range p.PaymentJson {
		queryParameters.Add(`paymentJson[]`, p.PaymentJson[i])
	}

	return t.restClient.Post(
		`/v2/Transaction/UseCase/ProcessTransaction`,
		&queryParameters,
		nil,
		nil,
	)
}
