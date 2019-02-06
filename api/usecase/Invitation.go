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

type GetInvitationParameters struct {
	InvitationId string
	WithData     *[]string // Event | UserName | User | UserIdentifier | Stack | TicketType | QuestionResponse | Answer
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

type ListInvitationsForEventParameters struct {
	EventId               string
	WithData              *[]string // UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-250
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
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForStackParameters struct {
	EventId               string
	StackId               string
	WithData              *[]string // UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-250
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
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForTicketBlockParameters struct {
	TicketBlockId         string
	WithData              *[]string // UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-250
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
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListInvitationsForUserParameters struct {
	UserId              string
	PoolId              *string
	EventId             *string
	Page                *int64 // >= 1
	ItemsPerPage        *int64 // 1-250
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string // Event | Stack | StackAndTicketType
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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

type ListInvitationsForUserByEmailParameters struct {
	Email               string
	PoolId              *string
	EventId             *string
	Page                *int64 // >= 1
	ItemsPerPage        *int64 // 1-250
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string // Event | Stack | StackAndTicketType
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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

type ListInvitationsForUserForParentParameters struct {
	UserId              string
	ParentEventId       string
	PoolId              *string
	Page                *int64 // >= 1
	ItemsPerPage        *int64 // 1-100
	EventDateFilterType *string
	SortDirection       *string
	WithData            *[]string // Event | Stack
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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

type ListWaitlistForEventParameters struct {
	EventId               string
	WithData              *[]string // UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom
	Query                 *string
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-100
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
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListWaitlistForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListWaitlistForStackParameters struct {
	EventId               string
	StackId               string
	WithData              *[]string // UserIdentifiers | StackAndTicketType | QuestionResponses | maxLastModifiedAt
	WithUserAttributes    *[]string // internal | info | hover | facebook | linked-in | salesforce | twitter | convio | google | custom
	Query                 *string
	StatusFilter          *[]string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	LastModifiedTimestamp *int64
	IsCheckedIn           *bool
	SortBy                *string
	SortDirection         *string
	Page                  *int64 // >= 1
	ItemsPerPage          *int64 // 1-100
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
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListWaitlistForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

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

type ChangeInviteCountParameters struct {
	InvitationId string
	InviteCount  *int64
}

func (t *Invitation) ChangeInviteCount(p *ChangeInviteCountParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.InviteCount != nil {
		queryParameters.Add(`inviteCount`, strconv.FormatInt(*p.InviteCount, 10))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ChangeInviteCount`,
		&queryParameters,
		nil,
		nil,
	)
}

type CheckInParameters struct {
	InvitationId string
	CheckInAt    *int64
}

func (t *Invitation) CheckIn(p *CheckInParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.CheckInAt != nil {
		queryParameters.Add(`checkInAt`, strconv.FormatInt(*p.CheckInAt, 10))
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CheckIn`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateInvitationParameters struct {
	EventId              string
	StackId              string
	InvitationStatus     string
	InviteSource         string
	IsCheckedIn          bool
	InviteCount          int64 // >= 1
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
	queryParameters.Add(`inviteCount`, strconv.FormatInt(p.InviteCount, 10))
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

type CreateInvitationForTicketBlockParameters struct {
	EventId              string
	StackId              string
	TicketBlockId        string
	InvitationStatus     string
	InviteSource         string
	IsCheckedIn          bool
	InviteCount          int64 // >= 1
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
	queryParameters.Add(`inviteCount`, strconv.FormatInt(p.InviteCount, 10))
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

type CreateInvitationsFromGroupParameters struct {
	GroupId                string
	StackId                string
	GuestsPerInvitation    int64 // >= 1
	InvitationCreationType string
}

func (t *Invitation) CreateInvitationsFromGroup(p *CreateInvitationsFromGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(p.GuestsPerInvitation, 10))
	queryParameters.Add(`invitationCreationType`, p.InvitationCreationType)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateInvitationsFromGroupForCIOEventParameters struct {
	EventId             string
	GroupId             string
	GuestsPerInvitation int64 // >= 1
	StackId             *string
}

func (t *Invitation) CreateInvitationsFromGroupForCIOEvent(p *CreateInvitationsFromGroupForCIOEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(p.GuestsPerInvitation, 10))
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

type CreateInvitationsFromGroupForTicketBlockParameters struct {
	GroupId                string
	TicketBlockId          string
	StackId                string
	GuestsPerInvitation    int64 // >= 1
	InvitationCreationType string
}

func (t *Invitation) CreateInvitationsFromGroupForTicketBlock(p *CreateInvitationsFromGroupForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(p.GuestsPerInvitation, 10))
	queryParameters.Add(`invitationCreationType`, p.InvitationCreationType)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateInvitationsFromGroupForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

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
