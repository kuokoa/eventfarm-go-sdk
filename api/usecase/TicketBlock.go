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

type TicketBlock struct {
	restClient rest.RestClientInterface
}

func NewTicketBlock(restClient rest.RestClientInterface) *TicketBlock {
	return &TicketBlock{restClient}
}

// GET: Queries

type GetTicketBlockParameters struct {
	TicketBlockId string
	WithData      *[]string // Event | Allotments | AllotmentsAndStack | AllotmentCounts
}

func (t *TicketBlock) GetTicketBlock(p *GetTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/TicketBlock/UseCase/GetTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTicketBlocksForEventParameters struct {
	EventId             string
	Query               *string
	WithData            *[]string // Event | Allotments | AllotmentsAndStack | AllotmentCounts
	Page                *int64    // >= 1
	ItemsPerPage        *int64    // 1-50
	SortBy              *string
	SortDirection       *string
	EventDateFilterType *string
	ShouldHideDeleted   *bool
}

func (t *TicketBlock) ListTicketBlocksForEvent(p *ListTicketBlocksForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
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
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*p.ShouldHideDeleted))
	}

	return t.restClient.Get(
		`/v2/TicketBlock/UseCase/ListTicketBlocksForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListTicketBlocksForUserParameters struct {
	UserId              string
	Query               *string
	WithData            *[]string // Event
	Page                *int64    // >= 1
	ItemsPerPage        *int64    // 1-500
	SortBy              *string
	SortDirection       *string
	EventDateFilterType *string
	PoolId              *string
}

func (t *TicketBlock) ListTicketBlocksForUser(p *ListTicketBlocksForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
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
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Get(
		`/v2/TicketBlock/UseCase/ListTicketBlocksForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddUserRoleToTicketBlockParameters struct {
	TicketBlockId       string
	Email               string
	FirstName           string
	LastName            string
	AuthenticatedUserId *string
}

func (t *TicketBlock) AddUserRoleToTicketBlock(p *AddUserRoleToTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`email`, p.Email)
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	if p.AuthenticatedUserId != nil {
		queryParameters.Add(`authenticatedUserId`, *p.AuthenticatedUserId)
	}

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/AddUserRoleToTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateTicketBlockParameters struct {
	EventId            string
	Name               string
	IsBlacklistEnabled *bool
	EmailText          *string
	TicketBlockId      *string
}

func (t *TicketBlock) CreateTicketBlock(p *CreateTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`name`, p.Name)
	if p.IsBlacklistEnabled != nil {
		queryParameters.Add(`isBlacklistEnabled`, strconv.FormatBool(*p.IsBlacklistEnabled))
	}
	if p.EmailText != nil {
		queryParameters.Add(`emailText`, *p.EmailText)
	}
	if p.TicketBlockId != nil {
		queryParameters.Add(`ticketBlockId`, *p.TicketBlockId)
	}

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/CreateTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type DisableTicketBlockBlacklistParameters struct {
	TicketBlockId string
}

func (t *TicketBlock) DisableTicketBlockBlacklist(p *DisableTicketBlockBlacklistParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/DisableTicketBlockBlacklist`,
		&queryParameters,
		nil,
		nil,
	)
}

type EnableTicketBlockBlacklistParameters struct {
	TicketBlockId string
}

func (t *TicketBlock) EnableTicketBlockBlacklist(p *EnableTicketBlockBlacklistParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/EnableTicketBlockBlacklist`,
		&queryParameters,
		nil,
		nil,
	)
}

type RemoveAllUserRolesFromTicketBlockParameters struct {
	UserId        string
	TicketBlockId string
}

func (t *TicketBlock) RemoveAllUserRolesFromTicketBlock(p *RemoveAllUserRolesFromTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/RemoveAllUserRolesFromTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type RemoveTicketBlockParameters struct {
	TicketBlockId string
}

func (t *TicketBlock) RemoveTicketBlock(p *RemoveTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/RemoveTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetTicketBlockEmailTextParameters struct {
	TicketBlockId string
	EmailText     string
}

func (t *TicketBlock) SetTicketBlockEmailText(p *SetTicketBlockEmailTextParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`emailText`, p.EmailText)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/SetTicketBlockEmailText`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetTicketBlockNameParameters struct {
	TicketBlockId string
	Name          string
}

func (t *TicketBlock) SetTicketBlockName(p *SetTicketBlockNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/SetTicketBlockName`,
		&queryParameters,
		nil,
		nil,
	)
}
