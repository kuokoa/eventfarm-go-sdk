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

type TicketBlock struct {
	restClient gosdk.RestClientInterface
}

func NewTicketBlock(restClient gosdk.RestClientInterface) *TicketBlock {
	return &TicketBlock{restClient}
}

// GET: Queries
// @param string TicketBlockId
// @param array|null WithData Event|Allotments|AllotmentsAndStack|AllotmentCounts
func (t *TicketBlock) GetTicketBlock(TicketBlockId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *TicketBlock) ListTicketBlocksForEvent(EventId string, Query *string, WithData *[]string, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string, EventDateFilterType *string, ShouldHideDeleted *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *EventDateFilterType)
	}
	if ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*ShouldHideDeleted))
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
func (t *TicketBlock) ListTicketBlocksForUser(UserId string, Query *string, WithData *[]string, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string, EventDateFilterType *string, PoolId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *EventDateFilterType)
	}
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
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
func (t *TicketBlock) AddUserRoleToTicketBlock(TicketBlockId string, Email string, FirstName string, LastName string, AuthenticatedUserId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	queryParameters.Add(`email`, Email)
	queryParameters.Add(`firstName`, FirstName)
	queryParameters.Add(`lastName`, LastName)
	if AuthenticatedUserId != nil {
		queryParameters.Add(`authenticatedUserId`, *AuthenticatedUserId)
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
func (t *TicketBlock) CreateTicketBlock(EventId string, Name string, IsBlacklistEnabled *bool, EmailText *string, TicketBlockId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`name`, Name)
	if IsBlacklistEnabled != nil {
		queryParameters.Add(`isBlacklistEnabled`, strconv.FormatBool(*IsBlacklistEnabled))
	}
	if EmailText != nil {
		queryParameters.Add(`emailText`, *EmailText)
	}
	if TicketBlockId != nil {
		queryParameters.Add(`ticketBlockId`, *TicketBlockId)
	}

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/CreateTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
func (t *TicketBlock) DisableTicketBlockBlacklist(TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/DisableTicketBlockBlacklist`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
func (t *TicketBlock) EnableTicketBlockBlacklist(TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/EnableTicketBlockBlacklist`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string TicketBlockId
func (t *TicketBlock) RemoveAllUserRolesFromTicketBlock(UserId string, TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/RemoveAllUserRolesFromTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
func (t *TicketBlock) RemoveTicketBlock(TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/RemoveTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
// @param string EmailText
func (t *TicketBlock) SetTicketBlockEmailText(TicketBlockId string, EmailText string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	queryParameters.Add(`emailText`, EmailText)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/SetTicketBlockEmailText`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
// @param string Name
func (t *TicketBlock) SetTicketBlockName(TicketBlockId string, Name string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	queryParameters.Add(`name`, Name)

	return t.restClient.Post(
		`/v2/TicketBlock/UseCase/SetTicketBlockName`,
		&queryParameters,
		nil,
		nil,
	)
}
