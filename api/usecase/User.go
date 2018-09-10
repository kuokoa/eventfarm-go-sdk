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

type User struct {
	restClient gosdk.RestClientInterface
}

func NewUser(restClient gosdk.RestClientInterface) *User {
	return &User{restClient}
}

// GET: Queries
// @param string RemoveUserId
// @param string RequestUserId
// @param string PoolId
func (t *User) CheckIfUserCanBeRemovedFromPool(RemoveUserId string, RequestUserId string, PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`removeUserId`, RemoveUserId)
	queryParameters.Add(`requestUserId`, RequestUserId)
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/User/UseCase/CheckIfUserCanBeRemovedFromPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param array|null WithData UserName|UserAddress|UserToken|UserIdentifier|isEFAdmin|internalUserName
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
func (t *User) GetUser(UserId string, WithData *[]string, WithUserAttributes *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
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

	return t.restClient.Get(
		`/v2/User/UseCase/GetUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string UserId
// @param array|null WithData UserName|UserAddress|UserToken|UserIdentifier|isEFAdmin|internalUserName
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
func (t *User) GetUserInPool(PoolId string, UserId string, WithData *[]string, WithUserAttributes *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
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

	return t.restClient.Get(
		`/v2/User/UseCase/GetUserInPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string EventId
func (t *User) GetUserRolesForEvent(UserId string, EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/User/UseCase/GetUserRolesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string TicketBlockId
func (t *User) GetUserRolesForTicketBlock(UserId string, TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Get(
		`/v2/User/UseCase/GetUserRolesForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param array PoolIds
// @param array|null WithData UserIdentifiers|UserNames|UserAttributes
// @param string|null Query
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *User) ListUsersForPools(PoolIds []string, WithData *[]string, Query *string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range PoolIds {
		queryParameters.Add(`poolIds[]`, PoolIds[i])
	}
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
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
		`/v2/User/UseCase/ListUsersForPools`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketBlockId
// @param array|null WithData UserIdentifiers|UserNames|UserAttributes
// @param string|null Query
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *User) ListUsersForTicketBlock(TicketBlockId string, WithData *[]string, Query *string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
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
		`/v2/User/UseCase/ListUsersForTicketBlock`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId
// @param string PoolId
// @param array|null WithData UserIdentifiers|UserNames|UserAttributes
// @param string|null Query
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *User) ListUsersInGroup(GroupId string, PoolId string, WithData *[]string, Query *string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	queryParameters.Add(`poolId`, PoolId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
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
		`/v2/User/UseCase/ListUsersInGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param array|null WithData eventRoles|PoolContacts|UserIdentifiers|UserNames|UserAttributes
// @param string|null Query
// @param string|null SortBy name
// @param string|null SortDirection ascending|descending
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *User) ListUsersWithRolesForEvent(EventId string, WithData *[]string, Query *string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
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
		`/v2/User/UseCase/ListUsersWithRolesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string UserId
// @param string Token
func (t *User) AccessUserForgotPasswordToken(UserId string, Token string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`token`, Token)

	return t.restClient.Post(
		`/v2/User/UseCase/AccessUserForgotPasswordToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string GroupId
func (t *User) AddUserAccessToGroup(UserId string, GroupId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`groupId`, GroupId)

	return t.restClient.Post(
		`/v2/User/UseCase/AddUserAccessToGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string|null Email
// @param string|null FirstName
// @param string|null LastName
// @param string|null Company
// @param string|null Position
// @param string|null Phone
// @param string|null PoolId
// @param string|null Title
// @param string|null Other
func (t *User) CreateAuthUser(Email *string, FirstName *string, LastName *string, Company *string, Position *string, Phone *string, PoolId *string, Title *string, Other *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
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
	if Phone != nil {
		queryParameters.Add(`phone`, *Phone)
	}
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if Title != nil {
		queryParameters.Add(`title`, *Title)
	}
	if Other != nil {
		queryParameters.Add(`other`, *Other)
	}

	return t.restClient.Post(
		`/v2/User/UseCase/CreateAuthUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string|null Email
// @param string|null FirstName
// @param string|null LastName
// @param string|null Company
// @param string|null Position
// @param string|null Phone
// @param string|null Title
// @param string|null Other
func (t *User) CreateCIOAccount(Email *string, FirstName *string, LastName *string, Company *string, Position *string, Phone *string, Title *string, Other *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
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
	if Phone != nil {
		queryParameters.Add(`phone`, *Phone)
	}
	if Title != nil {
		queryParameters.Add(`title`, *Title)
	}
	if Other != nil {
		queryParameters.Add(`other`, *Other)
	}

	return t.restClient.Post(
		`/v2/User/UseCase/CreateCIOAccount`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
func (t *User) CreateForgotPasswordToken(UserId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)

	return t.restClient.Post(
		`/v2/User/UseCase/CreateForgotPasswordToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string|null Email
// @param string|null FirstName
// @param string|null LastName
// @param string|null Company
// @param string|null Position
// @param string|null Phone
// @param string|null PoolId
// @param string|null Title
// @param string|null Other
func (t *User) CreateUser(Email *string, FirstName *string, LastName *string, Company *string, Position *string, Phone *string, PoolId *string, Title *string, Other *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
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
	if Phone != nil {
		queryParameters.Add(`phone`, *Phone)
	}
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if Title != nil {
		queryParameters.Add(`title`, *Title)
	}
	if Other != nil {
		queryParameters.Add(`other`, *Other)
	}

	return t.restClient.Post(
		`/v2/User/UseCase/CreateUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string PoolId
// @param string Email
// @param string|null UserContactAgentId
func (t *User) CreateUserContactAgent(UserId string, PoolId string, Email string, UserContactAgentId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`email`, Email)
	if UserContactAgentId != nil {
		queryParameters.Add(`userContactAgentId`, *UserContactAgentId)
	}

	return t.restClient.Post(
		`/v2/User/UseCase/CreateUserContactAgent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string EventId
func (t *User) RemoveEventRoleForUser(UserId string, EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/User/UseCase/RemoveEventRoleForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string GroupId
func (t *User) RemoveUserAccessToGroup(UserId string, GroupId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`groupId`, GroupId)

	return t.restClient.Post(
		`/v2/User/UseCase/RemoveUserAccessToGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserContactAgentId
func (t *User) RemoveUserContactAgent(UserContactAgentId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userContactAgentId`, UserContactAgentId)

	return t.restClient.Post(
		`/v2/User/UseCase/RemoveUserContactAgent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param array RemoveUserIds
// @param string RequestUserId
// @param string PoolId
func (t *User) RemoveUsersFromPool(RemoveUserIds []string, RequestUserId string, PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range RemoveUserIds {
		queryParameters.Add(`removeUserIds[]`, RemoveUserIds[i])
	}
	queryParameters.Add(`requestUserId`, RequestUserId)
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Post(
		`/v2/User/UseCase/RemoveUsersFromPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Email
func (t *User) SendForgotPasswordEmail(Email string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, Email)

	return t.restClient.Post(
		`/v2/User/UseCase/SendForgotPasswordEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Email
func (t *User) SendVerificationEmail(Email string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, Email)

	return t.restClient.Post(
		`/v2/User/UseCase/SendVerificationEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
// @param string AuthUserId
// @param string ChangeUserId
// @param string Email
func (t *User) SetEmailForInvitation(InvitationId string, AuthUserId string, ChangeUserId string, Email string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	queryParameters.Add(`authUserId`, AuthUserId)
	queryParameters.Add(`changeUserId`, ChangeUserId)
	queryParameters.Add(`email`, Email)

	return t.restClient.Post(
		`/v2/User/UseCase/SetEmailForInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string EventId
// @param string EventRole organizer|assistant|support|check-in-staff|read-only
// @param string AuthenticatedUserId
func (t *User) SetEventRoleForUser(UserId string, EventId string, EventRole string, AuthenticatedUserId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`eventRole`, EventRole)
	queryParameters.Add(`authenticatedUserId`, AuthenticatedUserId)

	return t.restClient.Post(
		`/v2/User/UseCase/SetEventRoleForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserContactAgentId
// @param string Email
func (t *User) UpdateUserContactAgent(UserContactAgentId string, Email string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userContactAgentId`, UserContactAgentId)
	queryParameters.Add(`email`, Email)

	return t.restClient.Post(
		`/v2/User/UseCase/UpdateUserContactAgent`,
		&queryParameters,
		nil,
		nil,
	)
}
