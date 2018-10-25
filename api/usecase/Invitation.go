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

type Invitation struct {
	restClient rest.RestClientInterface
}

func NewInvitation(restClient rest.RestClientInterface) *Invitation {
	return &Invitation{restClient}
}

// GET: Queries
// @param string StackMethodType public-registration|public-purchase|invite-to-register|invite-to-purchase|invite-to-rsvp|invite-to-register-fcfs|invite-to-purchase-fcfs|invite-to-rsvp-fcfs

type GetAllInvitationStatusTypesForStackMethodTypeParameters struct {
	StackMethodType string
}

func (t *Invitation) GetAllInvitationStatusTypesForStackMethodType(p *GetAllInvitationStatusTypesForStackMethodTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackMethodType`, p.StackMethodType)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetAllInvitationStatusTypesForStackMethodType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId

type GetCheckInCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetCheckInCountsForEvent(p *GetCheckInCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetCheckInCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId

type GetCheckInCountsForTicketBlockParameters struct {
	TicketBlockId string
}

func (t *Invitation) GetCheckInCountsForTicketBlock(p *GetCheckInCountsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetCheckInCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param array|null WithData Event|UserName|User|UserIdentifier|Stack|TicketType|QuestionResponse|Answer

type GetInvitationParameters struct {
	InvitationId string
	WithData     *[]string
}

func (t *Invitation) GetInvitation(p *GetInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId

type GetInvitationCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetInvitationCountsForEvent(p *GetInvitationCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId

type GetInvitationCountsForTicketBlockParameters struct {
	TicketBlockId string
}

func (t *Invitation) GetInvitationCountsForTicketBlock(p *GetInvitationCountsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string|null PoolId

type GetInvitationCountsForUserParameters struct {
	UserId string
	PoolId *string
}

func (t *Invitation) GetInvitationCountsForUser(p *GetInvitationCountsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId

type GetInvitationLastActionCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetInvitationLastActionCountsForEvent(p *GetInvitationLastActionCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationLastActionCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId

type GetInvitationStatusTypeCountsForEventParameters struct {
	EventId string
}

func (t *Invitation) GetInvitationStatusTypeCountsForEvent(p *GetInvitationStatusTypeCountsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationStatusTypeCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId

type GetInvitationStatusTypeCountsForTicketBlockParameters struct {
	TicketBlockId string
}

func (t *Invitation) GetInvitationStatusTypeCountsForTicketBlock(p *GetInvitationStatusTypeCountsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationStatusTypeCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param array|null WithData UserIdentifiers|StackAndTicketType|QuestionResponses|maxLastModifiedAt
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
// @param string|null Query
// @param array|null StatusFilter assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param int|null LastModifiedTimestamp
// @param bool|null IsCheckedIn true|false
// @param string|null SortBy name|last-action|last-notified|created-at|modified-at|checked-in-at
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-250

type ListInvitationsForEventParameters struct {
	EventId               string
	WithData              *[]string
	WithUserAttributes    *[]string
	Query                 *string
	StatusFilter          *[]string
	LastModifiedTimestamp *int
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int
	ItemsPerPage          *int
}

func (t *Invitation) ListInvitationsForEvent(p *ListInvitationsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*p.LastModifiedTimestamp))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListInvitationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string StackId
// @param array|null WithData UserIdentifiers|StackAndTicketType|QuestionResponses|maxLastModifiedAt
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
// @param string|null Query
// @param array|null StatusFilter assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param int|null LastModifiedTimestamp
// @param bool|null IsCheckedIn true|false
// @param string|null SortBy name|last-action|last-notified|created-at|modified-at|checked-in-at
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-250

type ListInvitationsForStackParameters struct {
	EventId               string
	StackId               string
	WithData              *[]string
	WithUserAttributes    *[]string
	Query                 *string
	StatusFilter          *[]string
	LastModifiedTimestamp *int
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int
	ItemsPerPage          *int
}

func (t *Invitation) ListInvitationsForStack(p *ListInvitationsForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*p.LastModifiedTimestamp))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListInvitationsForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
// @param array|null WithData UserIdentifiers|StackAndTicketType|QuestionResponses|maxLastModifiedAt
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
// @param string|null Query
// @param array|null StatusFilter assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param int|null LastModifiedTimestamp
// @param bool|null IsCheckedIn true|false
// @param string|null SortBy name|last-action|last-notified|created-at|modified-at|checked-in-at
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-250

type ListInvitationsForTicketBlockParameters struct {
	TicketBlockId         string
	WithData              *[]string
	WithUserAttributes    *[]string
	Query                 *string
	StatusFilter          *[]string
	LastModifiedTimestamp *int
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int
	ItemsPerPage          *int
}

func (t *Invitation) ListInvitationsForTicketBlock(p *ListInvitationsForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*p.LastModifiedTimestamp))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListInvitationsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string|null PoolId
// @param string|null EventId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-250
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null SortDirection
// @param array|null WithData Event|Stack
// @param array|null StatusFilter

type ListInvitationsForUserParameters struct {
	UserId              string
	PoolId              *string
	EventId             *string
	Page                *int
	ItemsPerPage        *int
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string
	StatusFilter        *[]string
}

func (t *Invitation) ListInvitationsForUser(p *ListInvitationsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Email
// @param string|null PoolId
// @param string|null EventId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-250
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null SortDirection
// @param array|null WithData Event|Stack
// @param array|null StatusFilter

type ListInvitationsForUserByEmailParameters struct {
	Email               string
	PoolId              *string
	EventId             *string
	Page                *int
	ItemsPerPage        *int
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string
	StatusFilter        *[]string
}

func (t *Invitation) ListInvitationsForUserByEmail(p *ListInvitationsForUserByEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUserByEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string ParentEventId
// @param string|null PoolId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null SortDirection
// @param array|null WithData Event|Stack
// @param array|null StatusFilter

type ListInvitationsForUserForParentParameters struct {
	UserId              string
	ParentEventId       string
	PoolId              *string
	Page                *int
	ItemsPerPage        *int
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string
	StatusFilter        *[]string
}

func (t *Invitation) ListInvitationsForUserForParent(p *ListInvitationsForUserForParentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`parentEventId`, p.ParentEventId)
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUserForParent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param array|null WithData UserIdentifiers|StackAndTicketType|QuestionResponses|maxLastModifiedAt
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
// @param string|null Query
// @param int|null LastModifiedTimestamp
// @param bool|null IsCheckedIn true|false
// @param string|null SortBy name|last-action|last-notified|created-at|modified-at|checked-in-at
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListWaitlistForEventParameters struct {
	EventId               string
	WithData              *[]string
	WithUserAttributes    *[]string
	Query                 *string
	LastModifiedTimestamp *int
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int
	ItemsPerPage          *int
}

func (t *Invitation) ListWaitlistForEvent(p *ListWaitlistForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*p.LastModifiedTimestamp))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListWaitlistForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string StackId
// @param array|null WithData UserIdentifiers|StackAndTicketType|QuestionResponses|maxLastModifiedAt
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
// @param string|null Query
// @param array|null StatusFilter assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param int|null LastModifiedTimestamp
// @param bool|null IsCheckedIn true|false
// @param string|null SortBy name|last-action|last-notified|created-at|modified-at|checked-in-at
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListWaitlistForStackParameters struct {
	EventId               string
	StackId               string
	WithData              *[]string
	WithUserAttributes    *[]string
	Query                 *string
	StatusFilter          *[]string
	LastModifiedTimestamp *int
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int
	ItemsPerPage          *int
}

func (t *Invitation) ListWaitlistForStack(p *ListWaitlistForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.WithUserAttributes != nil {
		for i := range *p.WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*p.WithUserAttributes)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.StatusFilter != nil {
		for i := range *p.StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*p.StatusFilter)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*p.LastModifiedTimestamp))
	}
	if p.IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*p.IsCheckedIn))
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
		`/v2/Invitation/UseCase/ListWaitlistForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string InvitationId

type AddInvitationToWaitlistParameters struct {
	InvitationId string
}

func (t *Invitation) AddInvitationToWaitlist(p *AddInvitationToWaitlistParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/AddInvitationToWaitlist`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string InvitationStatus assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted

type ChangeInvitationStatusParameters struct {
	InvitationId     string
	InvitationStatus string
}

func (t *Invitation) ChangeInvitationStatus(p *ChangeInvitationStatusParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ChangeInvitationStatus`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param int|null InviteCount

type ChangeInviteCountParameters struct {
	InvitationId string
	InviteCount  *int
}

func (t *Invitation) ChangeInviteCount(p *ChangeInviteCountParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.InviteCount != nil {
		queryParameters.Add(`inviteCount`, strconv.Itoa(*p.InviteCount))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ChangeInviteCount`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param int|null CheckInAt

type CheckInParameters struct {
	InvitationId string
	CheckInAt    *int
}

func (t *Invitation) CheckIn(p *CheckInParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.CheckInAt != nil {
		queryParameters.Add(`checkInAt`, strconv.Itoa(*p.CheckInAt))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CheckIn`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string StackId
// @param string InvitationStatus assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param string InviteSource direct-invite|api-addition|distribution|event-invite|group-invite|import|ios-leave-behind|leave-behind|mobile-leave-behind|mobile-purchase|outside-purchase|public-interface|transferred|import-salesforce|import-marketo
// @param bool IsCheckedIn true|false
// @param int InviteCount >= 1
// @param string|null Email
// @param string|null FirstName
// @param string|null LastName
// @param string|null Company
// @param string|null Position
// @param string|null CheckInNotes
// @param string|null InvitationId
// @param bool|null ShouldSendInvitation true|false
// @param string|null InvitationNotes
// @param string|null Title
// @param string|null Telephone
// @param string|null Other

type CreateInvitationParameters struct {
	EventId              string
	StackId              string
	InvitationStatus     string
	InviteSource         string
	IsCheckedIn          bool
	InviteCount          int
	Email                *string
	FirstName            *string
	LastName             *string
	Company              *string
	Position             *string
	CheckInNotes         *string
	InvitationId         *string
	ShouldSendInvitation *bool
	InvitationNotes      *string
	Title                *string
	Telephone            *string
	Other                *string
}

func (t *Invitation) CreateInvitation(p *CreateInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)
	queryParameters.Add(`inviteSource`, p.InviteSource)
	queryParameters.Add(`isCheckedIn`, strconv.FormatBool(p.IsCheckedIn))
	queryParameters.Add(`inviteCount`, strconv.Itoa(p.InviteCount))
	if p.Email != nil {
		queryParameters.Add(`email`, *p.Email)
	}
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}
	if p.Company != nil {
		queryParameters.Add(`company`, *p.Company)
	}
	if p.Position != nil {
		queryParameters.Add(`position`, *p.Position)
	}
	if p.CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *p.CheckInNotes)
	}
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}
	if p.ShouldSendInvitation != nil {
		queryParameters.Add(`shouldSendInvitation`, strconv.FormatBool(*p.ShouldSendInvitation))
	}
	if p.InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *p.InvitationNotes)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Telephone != nil {
		queryParameters.Add(`telephone`, *p.Telephone)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string StackId
// @param string TicketBlockId
// @param string InvitationStatus assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param string InviteSource direct-invite|api-addition|distribution|event-invite|group-invite|import|ios-leave-behind|leave-behind|mobile-leave-behind|mobile-purchase|outside-purchase|public-interface|transferred|import-salesforce|import-marketo
// @param bool IsCheckedIn true|false
// @param int InviteCount >= 1
// @param string|null Email
// @param string|null FirstName
// @param string|null LastName
// @param string|null Company
// @param string|null Position
// @param string|null CheckInNotes
// @param string|null InvitationId
// @param bool|null ShouldSendInvitation true|false
// @param string|null InvitationNotes
// @param string|null Title
// @param string|null Telephone
// @param string|null Other

type CreateInvitationForTicketBlockParameters struct {
	EventId              string
	StackId              string
	TicketBlockId        string
	InvitationStatus     string
	InviteSource         string
	IsCheckedIn          bool
	InviteCount          int
	Email                *string
	FirstName            *string
	LastName             *string
	Company              *string
	Position             *string
	CheckInNotes         *string
	InvitationId         *string
	ShouldSendInvitation *bool
	InvitationNotes      *string
	Title                *string
	Telephone            *string
	Other                *string
}

func (t *Invitation) CreateInvitationForTicketBlock(p *CreateInvitationForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)
	queryParameters.Add(`inviteSource`, p.InviteSource)
	queryParameters.Add(`isCheckedIn`, strconv.FormatBool(p.IsCheckedIn))
	queryParameters.Add(`inviteCount`, strconv.Itoa(p.InviteCount))
	if p.Email != nil {
		queryParameters.Add(`email`, *p.Email)
	}
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}
	if p.Company != nil {
		queryParameters.Add(`company`, *p.Company)
	}
	if p.Position != nil {
		queryParameters.Add(`position`, *p.Position)
	}
	if p.CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *p.CheckInNotes)
	}
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}
	if p.ShouldSendInvitation != nil {
		queryParameters.Add(`shouldSendInvitation`, strconv.FormatBool(*p.ShouldSendInvitation))
	}
	if p.InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *p.InvitationNotes)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Telephone != nil {
		queryParameters.Add(`telephone`, *p.Telephone)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId
// @param string StackId
// @param int GuestsPerInvitation >= 1
// @param string InvitationCreationType unconfirmed-no-email|confirmed-no-email|send-email

type CreateInvitationsFromGroupParameters struct {
	GroupId                string
	StackId                string
	GuestsPerInvitation    int
	InvitationCreationType string
}

func (t *Invitation) CreateInvitationsFromGroup(p *CreateInvitationsFromGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(p.GuestsPerInvitation))
	queryParameters.Add(`invitationCreationType`, p.InvitationCreationType)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string GroupId
// @param int GuestsPerInvitation >= 1
// @param string|null StackId

type CreateInvitationsFromGroupForCIOEventParameters struct {
	EventId             string
	GroupId             string
	GuestsPerInvitation int
	StackId             *string
}

func (t *Invitation) CreateInvitationsFromGroupForCIOEvent(p *CreateInvitationsFromGroupForCIOEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(p.GuestsPerInvitation))
	if p.StackId != nil {
		queryParameters.Add(`stackId`, *p.StackId)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroupForCIOEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId
// @param string TicketBlockId
// @param string StackId
// @param int GuestsPerInvitation >= 1
// @param string InvitationCreationType unconfirmed-no-email|confirmed-no-email|send-email

type CreateInvitationsFromGroupForTicketBlockParameters struct {
	GroupId                string
	TicketBlockId          string
	StackId                string
	GuestsPerInvitation    int
	InvitationCreationType string
}

func (t *Invitation) CreateInvitationsFromGroupForTicketBlock(p *CreateInvitationsFromGroupForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(p.GuestsPerInvitation))
	queryParameters.Add(`invitationCreationType`, p.InvitationCreationType)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroupForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string WebhookType
// @param string WebhookMethod
// @param string Url

type CreateWebhookParameters struct {
	EventId       string
	WebhookType   string
	WebhookMethod string
	Url           string
}

func (t *Invitation) CreateWebhook(p *CreateWebhookParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`webhookType`, p.WebhookType)
	queryParameters.Add(`webhookMethod`, p.WebhookMethod)
	queryParameters.Add(`url`, p.Url)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string WebhookId

type DeleteWebhookParameters struct {
	WebhookId string
}

func (t *Invitation) DeleteWebhook(p *DeleteWebhookParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`webhookId`, p.WebhookId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/DeleteWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId

type DisableArrivalAlertParameters struct {
	InvitationId string
}

func (t *Invitation) DisableArrivalAlert(p *DisableArrivalAlertParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/DisableArrivalAlert`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param array InvitationIds
// @param string NewInvitationStatus
// @param bool|null ShouldSendEmail true|false

type PromoteInvitationsFromWaitlistParameters struct {
	InvitationIds       []string
	NewInvitationStatus string
	ShouldSendEmail     *bool
}

func (t *Invitation) PromoteInvitationsFromWaitlist(p *PromoteInvitationsFromWaitlistParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.InvitationIds {
		queryParameters.Add(`invitationIds[]`, p.InvitationIds[i])
	}
	queryParameters.Add(`newInvitationStatus`, p.NewInvitationStatus)
	if p.ShouldSendEmail != nil {
		queryParameters.Add(`shouldSendEmail`, strconv.FormatBool(*p.ShouldSendEmail))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/PromoteInvitationsFromWaitlist`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string QuestionId
// @param string QuestionResponseIdsWithAnswersJson

type SetAllQuestionResponsesParameters struct {
	InvitationId                       string
	QuestionId                         string
	QuestionResponseIdsWithAnswersJson string
}

func (t *Invitation) SetAllQuestionResponses(p *SetAllQuestionResponsesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`questionResponseIdsWithAnswersJson`, p.QuestionResponseIdsWithAnswersJson)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetAllQuestionResponses`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string ToEmail
// @param array|null CcEmails
// @param bool|null ShouldSendArrivalAlert true|false

type SetArrivalAlertEmailParameters struct {
	InvitationId           string
	ToEmail                string
	CcEmails               *[]string
	ShouldSendArrivalAlert *bool
}

func (t *Invitation) SetArrivalAlertEmail(p *SetArrivalAlertEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`toEmail`, p.ToEmail)
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.ShouldSendArrivalAlert != nil {
		queryParameters.Add(`shouldSendArrivalAlert`, strconv.FormatBool(*p.ShouldSendArrivalAlert))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetArrivalAlertEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string|null CheckInNotes

type SetCheckInNotesParameters struct {
	InvitationId string
	CheckInNotes *string
}

func (t *Invitation) SetCheckInNotes(p *SetCheckInNotesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *p.CheckInNotes)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetCheckInNotes`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string|null InvitationNotes

type SetInvitationNotesParameters struct {
	InvitationId    string
	InvitationNotes *string
}

func (t *Invitation) SetInvitationNotes(p *SetInvitationNotesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *p.InvitationNotes)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetInvitationNotes`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string QuestionId
// @param array|null AnswerIds
// @param string|null Text

type SetQuestionResponseParameters struct {
	InvitationId string
	QuestionId   string
	AnswerIds    *[]string
	Text         *string
}

func (t *Invitation) SetQuestionResponse(p *SetQuestionResponseParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`questionId`, p.QuestionId)
	if p.AnswerIds != nil {
		for i := range *p.AnswerIds {
			queryParameters.Add(`answerIds[]`, (*p.AnswerIds)[i])
		}
	}
	if p.Text != nil {
		queryParameters.Add(`text`, *p.Text)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetQuestionResponse`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId

type UndoCheckInParameters struct {
	InvitationId string
}

func (t *Invitation) UndoCheckIn(p *UndoCheckInParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/UndoCheckIn`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string StackId
// @param string InvitationStatus assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param string|null Company
// @param string|null Position
// @param string|null Email
// @param string|null FirstName
// @param string|null LastName
// @param string|null Other
// @param string|null Telephone

type UpdateInvitationParameters struct {
	InvitationId     string
	StackId          string
	InvitationStatus string
	Company          *string
	Position         *string
	Email            *string
	FirstName        *string
	LastName         *string
	Other            *string
	Telephone        *string
}

func (t *Invitation) UpdateInvitation(p *UpdateInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`invitationStatus`, p.InvitationStatus)
	if p.Company != nil {
		queryParameters.Add(`company`, *p.Company)
	}
	if p.Position != nil {
		queryParameters.Add(`position`, *p.Position)
	}
	if p.Email != nil {
		queryParameters.Add(`email`, *p.Email)
	}
	if p.FirstName != nil {
		queryParameters.Add(`firstName`, *p.FirstName)
	}
	if p.LastName != nil {
		queryParameters.Add(`lastName`, *p.LastName)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}
	if p.Telephone != nil {
		queryParameters.Add(`telephone`, *p.Telephone)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/UpdateInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}
