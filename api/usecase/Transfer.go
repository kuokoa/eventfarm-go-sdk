/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"github.com/eventfarm/go-sdk/rest"
	"net/http"
	"net/url"
	"strconv"
)

type Transfer struct {
	restClient rest.RestClientInterface
}

func NewTransfer(restClient rest.RestClientInterface) *Transfer {
	return &Transfer{restClient}
}

// GET: Queries

type GetTransferParameters struct {
	TransferId string
	WithData   *[]interface{}
}

func (t *Transfer) GetTransfer(p *GetTransferParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transferId`, p.TransferId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Transfer/UseCase/GetTransfer`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTransfersForEventParameters struct {
	EventId       string
	WithData      *[]interface{}
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	SortBy        *string
	SortDirection *string
}

func (t *Transfer) ListTransfersForEvent(p *ListTransfersForEventParameters) (r *http.Response, err error) {
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
		`/v2/Transfer/UseCase/ListTransfersForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTransfersForInvitationParameters struct {
	InvitationId  string
	WithData      *[]interface{}
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	SortBy        *string
	SortDirection *string
}

func (t *Transfer) ListTransfersForInvitation(p *ListTransfersForInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
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
		`/v2/Transfer/UseCase/ListTransfersForInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type ConfirmTransferParameters struct {
	TransferId      string
	Code            string
	ShouldSendEmail *bool
}

func (t *Transfer) ConfirmTransfer(p *ConfirmTransferParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transferId`, p.TransferId)
	queryParameters.Add(`code`, p.Code)
	if p.ShouldSendEmail != nil {
		queryParameters.Add(`shouldSendEmail`, strconv.FormatBool(*p.ShouldSendEmail))
	}

	return t.restClient.Post(
		`/v2/Transfer/UseCase/ConfirmTransfer`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateTransferParameters struct {
	EventId         string
	SenderId        string
	FirstName       string
	LastName        string
	Email           string
	TransferQty     int64
	ShouldSendEmail *bool
	ForceTransfer   *bool
	TransferId      *string
}

func (t *Transfer) CreateTransfer(p *CreateTransferParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`senderId`, p.SenderId)
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	queryParameters.Add(`email`, p.Email)
	queryParameters.Add(`transferQty`, strconv.FormatInt(p.TransferQty, 10))
	if p.ShouldSendEmail != nil {
		queryParameters.Add(`shouldSendEmail`, strconv.FormatBool(*p.ShouldSendEmail))
	}
	if p.ForceTransfer != nil {
		queryParameters.Add(`forceTransfer`, strconv.FormatBool(*p.ForceTransfer))
	}
	if p.TransferId != nil {
		queryParameters.Add(`transferId`, *p.TransferId)
	}

	return t.restClient.Post(
		`/v2/Transfer/UseCase/CreateTransfer`,
		&queryParameters,
		nil,
		nil,
	)
}

type DeleteTransferParameters struct {
	TransferId string
}

func (t *Transfer) DeleteTransfer(p *DeleteTransferParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transferId`, p.TransferId)

	return t.restClient.Post(
		`/v2/Transfer/UseCase/DeleteTransfer`,
		&queryParameters,
		nil,
		nil,
	)
}

type ForceTransferParameters struct {
	TransferId      string
	ShouldSendEmail *bool
}

func (t *Transfer) ForceTransfer(p *ForceTransferParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transferId`, p.TransferId)
	if p.ShouldSendEmail != nil {
		queryParameters.Add(`shouldSendEmail`, strconv.FormatBool(*p.ShouldSendEmail))
	}

	return t.restClient.Post(
		`/v2/Transfer/UseCase/ForceTransfer`,
		&queryParameters,
		nil,
		nil,
	)
}

type ResendAuthorizationRequestEmailParameters struct {
	TransferId string
}

func (t *Transfer) ResendAuthorizationRequestEmail(p *ResendAuthorizationRequestEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`transferId`, p.TransferId)

	return t.restClient.Post(
		`/v2/Transfer/UseCase/ResendAuthorizationRequestEmail`,
		&queryParameters,
		nil,
		nil,
	)
}
