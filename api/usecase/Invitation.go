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

type Invitation struct {
	restClient gosdk.RestClientInterface
}

func NewInvitation(restClient gosdk.RestClientInterface) *Invitation {
	return &Invitation{restClient}
}

// GET: Queries
// @param string StackMethodType public-registration|public-purchase|invite-to-register|invite-to-purchase|invite-to-rsvp|invite-to-register-fcfs|invite-to-purchase-fcfs|invite-to-rsvp-fcfs
func (t *Invitation) GetAllInvitationStatusTypesForStackMethodType(StackMethodType string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackMethodType`, StackMethodType)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetAllInvitationStatusTypesForStackMethodType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Invitation) GetCheckInCountsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetCheckInCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
func (t *Invitation) GetCheckInCountsForTicketBlock(TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetCheckInCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param array|null WithData Event|UserName|User|UserIdentifier|Stack|TicketType|QuestionResponse|Answer
func (t *Invitation) GetInvitation(InvitationId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *Invitation) GetInvitationCountsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
func (t *Invitation) GetInvitationCountsForTicketBlock(TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string|null PoolId
func (t *Invitation) GetInvitationCountsForUser(UserId string, PoolId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationCountsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Invitation) GetInvitationLastActionCountsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationLastActionCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Invitation) GetInvitationStatusTypeCountsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/Invitation/UseCase/GetInvitationStatusTypeCountsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
func (t *Invitation) GetInvitationStatusTypeCountsForTicketBlock(TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

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
func (t *Invitation) ListInvitationsForEvent(EventId string, WithData *[]string, WithUserAttributes *[]string, Query *string, StatusFilter *[]string, LastModifiedTimestamp *int, IsCheckedIn *bool, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if WithUserAttributes != nil {
		for i := range *WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*WithUserAttributes)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if StatusFilter != nil {
		for i := range *StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*StatusFilter)[i])
		}
	}
	if LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*LastModifiedTimestamp))
	}
	if IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*IsCheckedIn))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Invitation) ListInvitationsForStack(EventId string, StackId string, WithData *[]string, WithUserAttributes *[]string, Query *string, StatusFilter *[]string, LastModifiedTimestamp *int, IsCheckedIn *bool, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`stackId`, StackId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if WithUserAttributes != nil {
		for i := range *WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*WithUserAttributes)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if StatusFilter != nil {
		for i := range *StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*StatusFilter)[i])
		}
	}
	if LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*LastModifiedTimestamp))
	}
	if IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*IsCheckedIn))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Invitation) ListInvitationsForTicketBlock(TicketBlockId string, WithData *[]string, WithUserAttributes *[]string, Query *string, StatusFilter *[]string, LastModifiedTimestamp *int, IsCheckedIn *bool, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if WithUserAttributes != nil {
		for i := range *WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*WithUserAttributes)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if StatusFilter != nil {
		for i := range *StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*StatusFilter)[i])
		}
	}
	if LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*LastModifiedTimestamp))
	}
	if IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*IsCheckedIn))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string PoolId
// @param string EventId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-250
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null SortDirection
// @param array|null WithData Event|Stack
// @param array|null StatusFilter
func (t *Invitation) ListInvitationsForUser(UserId string, PoolId string, EventId string, Page *int, ItemsPerPage *int, EventDateFilterType *string, SortDirection *string, WithData *[]string, StatusFilter *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`eventId`, EventId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}
	if EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *EventDateFilterType)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if StatusFilter != nil {
		for i := range *StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*StatusFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Invitation/UseCase/ListInvitationsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string ParentEventId
// @param string PoolId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null SortDirection
// @param array|null WithData Event|Stack
// @param array|null StatusFilter
func (t *Invitation) ListInvitationsForUserForParent(UserId string, ParentEventId string, PoolId string, Page *int, ItemsPerPage *int, EventDateFilterType *string, SortDirection *string, WithData *[]string, StatusFilter *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`parentEventId`, ParentEventId)
	queryParameters.Add(`poolId`, PoolId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}
	if EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *EventDateFilterType)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if StatusFilter != nil {
		for i := range *StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*StatusFilter)[i])
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
func (t *Invitation) ListWaitlistForEvent(EventId string, WithData *[]string, WithUserAttributes *[]string, Query *string, LastModifiedTimestamp *int, IsCheckedIn *bool, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if WithUserAttributes != nil {
		for i := range *WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*WithUserAttributes)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*LastModifiedTimestamp))
	}
	if IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*IsCheckedIn))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Invitation) ListWaitlistForStack(EventId string, StackId string, WithData *[]string, WithUserAttributes *[]string, Query *string, StatusFilter *[]string, LastModifiedTimestamp *int, IsCheckedIn *bool, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`stackId`, StackId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if WithUserAttributes != nil {
		for i := range *WithUserAttributes {
			queryParameters.Add(`withUserAttributes[]`, (*WithUserAttributes)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if StatusFilter != nil {
		for i := range *StatusFilter {
			queryParameters.Add(`statusFilter[]`, (*StatusFilter)[i])
		}
	}
	if LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*LastModifiedTimestamp))
	}
	if IsCheckedIn != nil {
		queryParameters.Add(`isCheckedIn`, strconv.FormatBool(*IsCheckedIn))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Invitation) AddInvitationToWaitlist(InvitationId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/AddInvitationToWaitlist`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string InvitationStatus assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
func (t *Invitation) ChangeInvitationStatus(InvitationId string, InvitationStatus string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	queryParameters.Add(`invitationStatus`, InvitationStatus)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/ChangeInvitationStatus`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param int|null InviteCount
func (t *Invitation) ChangeInviteCount(InvitationId string, InviteCount *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	if InviteCount != nil {
		queryParameters.Add(`inviteCount`, strconv.Itoa(*InviteCount))
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
func (t *Invitation) CheckIn(InvitationId string, CheckInAt *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	if CheckInAt != nil {
		queryParameters.Add(`checkInAt`, strconv.Itoa(*CheckInAt))
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
func (t *Invitation) CreateInvitation(EventId string, StackId string, InvitationStatus string, InviteSource string, IsCheckedIn bool, InviteCount int, Email *string, FirstName *string, LastName *string, Company *string, Position *string, CheckInNotes *string, InvitationId *string, ShouldSendInvitation *bool, InvitationNotes *string, Title *string, Telephone *string, Other *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`invitationStatus`, InvitationStatus)
	queryParameters.Add(`inviteSource`, InviteSource)
	queryParameters.Add(`isCheckedIn`, strconv.FormatBool(IsCheckedIn))
	queryParameters.Add(`inviteCount`, strconv.Itoa(InviteCount))
	if Email != nil {
		queryParameters.Add(`email`, *Email)
	}
	if FirstName != nil {
		queryParameters.Add(`firstName`, *FirstName)
	}
	if LastName != nil {
		queryParameters.Add(`lastName`, *LastName)
	}
	if Company != nil {
		queryParameters.Add(`company`, *Company)
	}
	if Position != nil {
		queryParameters.Add(`position`, *Position)
	}
	if CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *CheckInNotes)
	}
	if InvitationId != nil {
		queryParameters.Add(`invitationId`, *InvitationId)
	}
	if ShouldSendInvitation != nil {
		queryParameters.Add(`shouldSendInvitation`, strconv.FormatBool(*ShouldSendInvitation))
	}
	if InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *InvitationNotes)
	}
	if Title != nil {
		queryParameters.Add(`title`, *Title)
	}
	if Telephone != nil {
		queryParameters.Add(`telephone`, *Telephone)
	}
	if Other != nil {
		queryParameters.Add(`other`, *Other)
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
func (t *Invitation) CreateInvitationForTicketBlock(EventId string, StackId string, TicketBlockId string, InvitationStatus string, InviteSource string, IsCheckedIn bool, InviteCount int, Email *string, FirstName *string, LastName *string, Company *string, Position *string, CheckInNotes *string, InvitationId *string, ShouldSendInvitation *bool, InvitationNotes *string, Title *string, Telephone *string, Other *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	queryParameters.Add(`invitationStatus`, InvitationStatus)
	queryParameters.Add(`inviteSource`, InviteSource)
	queryParameters.Add(`isCheckedIn`, strconv.FormatBool(IsCheckedIn))
	queryParameters.Add(`inviteCount`, strconv.Itoa(InviteCount))
	if Email != nil {
		queryParameters.Add(`email`, *Email)
	}
	if FirstName != nil {
		queryParameters.Add(`firstName`, *FirstName)
	}
	if LastName != nil {
		queryParameters.Add(`lastName`, *LastName)
	}
	if Company != nil {
		queryParameters.Add(`company`, *Company)
	}
	if Position != nil {
		queryParameters.Add(`position`, *Position)
	}
	if CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *CheckInNotes)
	}
	if InvitationId != nil {
		queryParameters.Add(`invitationId`, *InvitationId)
	}
	if ShouldSendInvitation != nil {
		queryParameters.Add(`shouldSendInvitation`, strconv.FormatBool(*ShouldSendInvitation))
	}
	if InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *InvitationNotes)
	}
	if Title != nil {
		queryParameters.Add(`title`, *Title)
	}
	if Telephone != nil {
		queryParameters.Add(`telephone`, *Telephone)
	}
	if Other != nil {
		queryParameters.Add(`other`, *Other)
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
func (t *Invitation) CreateInvitationsFromGroup(GroupId string, StackId string, GuestsPerInvitation int, InvitationCreationType string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(GuestsPerInvitation))
	queryParameters.Add(`invitationCreationType`, InvitationCreationType)

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
func (t *Invitation) CreateInvitationsFromGroupForCIOEvent(EventId string, GroupId string, GuestsPerInvitation int, StackId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`groupId`, GroupId)
	queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(GuestsPerInvitation))
	if StackId != nil {
		queryParameters.Add(`stackId`, *StackId)
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
func (t *Invitation) CreateInvitationsFromGroupForTicketBlock(GroupId string, TicketBlockId string, StackId string, GuestsPerInvitation int, InvitationCreationType string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(GuestsPerInvitation))
	queryParameters.Add(`invitationCreationType`, InvitationCreationType)

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
func (t *Invitation) CreateWebhook(EventId string, WebhookType string, WebhookMethod string, Url string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`webhookType`, WebhookType)
	queryParameters.Add(`webhookMethod`, WebhookMethod)
	queryParameters.Add(`url`, Url)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/CreateWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string WebhookId
func (t *Invitation) DeleteWebhook(WebhookId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`webhookId`, WebhookId)

	return t.restClient.Post(
		`/v2/Invitation/UseCase/DeleteWebhook`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
func (t *Invitation) DisableArrivalAlert(InvitationId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)

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
func (t *Invitation) PromoteInvitationsFromWaitlist(InvitationIds []string, NewInvitationStatus string, ShouldSendEmail *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range InvitationIds {
		queryParameters.Add(`invitationIds[]`, InvitationIds[i])
	}
	queryParameters.Add(`newInvitationStatus`, NewInvitationStatus)
	if ShouldSendEmail != nil {
		queryParameters.Add(`shouldSendEmail`, strconv.FormatBool(*ShouldSendEmail))
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
func (t *Invitation) SetAllQuestionResponses(InvitationId string, QuestionId string, QuestionResponseIdsWithAnswersJson string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	queryParameters.Add(`questionId`, QuestionId)
	queryParameters.Add(`questionResponseIdsWithAnswersJson`, QuestionResponseIdsWithAnswersJson)

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
func (t *Invitation) SetArrivalAlertEmail(InvitationId string, ToEmail string, CcEmails *[]string, ShouldSendArrivalAlert *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	queryParameters.Add(`toEmail`, ToEmail)
	if CcEmails != nil {
		for i := range *CcEmails {
			queryParameters.Add(`ccEmails[]`, (*CcEmails)[i])
		}
	}
	if ShouldSendArrivalAlert != nil {
		queryParameters.Add(`shouldSendArrivalAlert`, strconv.FormatBool(*ShouldSendArrivalAlert))
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
func (t *Invitation) SetCheckInNotes(InvitationId string, CheckInNotes *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	if CheckInNotes != nil {
		queryParameters.Add(`checkInNotes`, *CheckInNotes)
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
func (t *Invitation) SetInvitationNotes(InvitationId string, InvitationNotes *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	if InvitationNotes != nil {
		queryParameters.Add(`invitationNotes`, *InvitationNotes)
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
func (t *Invitation) SetQuestionResponse(InvitationId string, QuestionId string, AnswerIds *[]string, Text *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	queryParameters.Add(`questionId`, QuestionId)
	if AnswerIds != nil {
		for i := range *AnswerIds {
			queryParameters.Add(`answerIds[]`, (*AnswerIds)[i])
		}
	}
	if Text != nil {
		queryParameters.Add(`text`, *Text)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/SetQuestionResponse`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
func (t *Invitation) UndoCheckIn(InvitationId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)

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
func (t *Invitation) UpdateInvitation(InvitationId string, StackId string, InvitationStatus string, Company *string, Position *string, Email *string, FirstName *string, LastName *string, Other *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`invitationStatus`, InvitationStatus)
	if Company != nil {
		queryParameters.Add(`company`, *Company)
	}
	if Position != nil {
		queryParameters.Add(`position`, *Position)
	}
	if Email != nil {
		queryParameters.Add(`email`, *Email)
	}
	if FirstName != nil {
		queryParameters.Add(`firstName`, *FirstName)
	}
	if LastName != nil {
		queryParameters.Add(`lastName`, *LastName)
	}
	if Other != nil {
		queryParameters.Add(`other`, *Other)
	}

	return t.restClient.Post(
		`/v2/Invitation/UseCase/UpdateInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}
