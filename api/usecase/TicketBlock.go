package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type TicketBlock struct {
	restClient sdk.RestClientInterface
}

func NewTicketBlock(restClient sdk.RestClientInterface) *TicketBlock {
	return &TicketBlock{restClient}
}

// GET: Queries
// @param string TicketBlockId
// @param array|null WithData Event|Allotments|AllotmentsAndStack|AllotmentCounts

type GetTicketBlockParameters struct {
	TicketBlockId string
	WithData      *[]string
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

// @param string EventId
// @param string|null Query
// @param array|null WithData Event|Allotments|AllotmentsAndStack|AllotmentCounts
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-50
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param bool|null ShouldHideDeleted true|false

type ListTicketBlocksForEventParameters struct {
	EventId             string
	Query               *string
	WithData            *[]string
	Page                *int
	ItemsPerPage        *int
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

// @param string UserId
// @param string|null Query
// @param array|null WithData Event
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
// @param string|null SortBy event-start|event-end
// @param string|null SortDirection ascending|descending
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null PoolId

type ListTicketBlocksForUserParameters struct {
	UserId              string
	Query               *string
	WithData            *[]string
	Page                *int
	ItemsPerPage        *int
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
// @param string TicketBlockId
// @param string Email
// @param string FirstName
// @param string LastName
// @param string|null AuthenticatedUserId

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

// @param string EventId
// @param string Name
// @param bool|null IsBlacklistEnabled true|false
// @param string|null EmailText
// @param string|null TicketBlockId

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

// @param string TicketBlockId

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

// @param string TicketBlockId

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

// @param string UserId
// @param string TicketBlockId

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

// @param string TicketBlockId

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

// @param string TicketBlockId
// @param string EmailText

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

// @param string TicketBlockId
// @param string Name

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
