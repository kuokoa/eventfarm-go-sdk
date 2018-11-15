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

type User struct {
	restClient rest.RestClientInterface
}

func NewUser(restClient rest.RestClientInterface) *User {
	return &User{restClient}
}

// GET: Queries
// @param string RemoveUserId
// @param string RequestUserId
// @param string PoolId

type CheckIfUserCanBeRemovedFromPoolParameters struct {
	RemoveUserId  string
	RequestUserId string
	PoolId        string
}

func (t *User) CheckIfUserCanBeRemovedFromPool(p *CheckIfUserCanBeRemovedFromPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`removeUserId`, p.RemoveUserId)
	queryParameters.Add(`requestUserId`, p.RequestUserId)
	queryParameters.Add(`poolId`, p.PoolId)

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
// @param string|null PoolId

type GetUserParameters struct {
	UserId             string
	WithData           *[]string
	WithUserAttributes *[]string
	PoolId             *string
}

func (t *User) GetUser(p *GetUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
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
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Get(
		`/v2/User/UseCase/GetUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Email
// @param array|null WithData UserName|UserAddress|UserToken|isEFAdmin|internalUserName
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom
// @param string|null PoolId

type GetUserByEmailParameters struct {
	Email              string
	WithData           *[]string
	WithUserAttributes *[]string
	PoolId             *string
}

func (t *User) GetUserByEmail(p *GetUserByEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)
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
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}

	return t.restClient.Get(
		`/v2/User/UseCase/GetUserByEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string UserId
// @param array|null WithData UserName|UserAddress|UserToken|UserIdentifier|isEFAdmin|internalUserName
// @param array|null WithUserAttributes internal|info|hover|facebook|linked-in|salesforce|twitter|convio|google|custom

type GetUserInPoolParameters struct {
	PoolId             string
	UserId             string
	WithData           *[]string
	WithUserAttributes *[]string
}

func (t *User) GetUserInPool(p *GetUserInPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
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

	return t.restClient.Get(
		`/v2/User/UseCase/GetUserInPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string EventId

type GetUserRolesForEventParameters struct {
	UserId  string
	EventId string
}

func (t *User) GetUserRolesForEvent(p *GetUserRolesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/User/UseCase/GetUserRolesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string TicketBlockId

type GetUserRolesForTicketBlockParameters struct {
	UserId        string
	TicketBlockId string
}

func (t *User) GetUserRolesForTicketBlock(p *GetUserRolesForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

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

type ListUsersForPoolsParameters struct {
	PoolIds       []string
	WithData      *[]string
	Query         *string
	SortBy        *string
	SortDirection *string
	Page          *int
	ItemsPerPage  *int
}

func (t *User) ListUsersForPools(p *ListUsersForPoolsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.PoolIds {
		queryParameters.Add(`poolIds[]`, p.PoolIds[i])
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
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

type ListUsersForTicketBlockParameters struct {
	TicketBlockId string
	WithData      *[]string
	Query         *string
	SortBy        *string
	SortDirection *string
	Page          *int
	ItemsPerPage  *int
}

func (t *User) ListUsersForTicketBlock(p *ListUsersForTicketBlockParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
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

type ListUsersInGroupParameters struct {
	GroupId       string
	PoolId        string
	WithData      *[]string
	Query         *string
	SortBy        *string
	SortDirection *string
	Page          *int
	ItemsPerPage  *int
}

func (t *User) ListUsersInGroup(p *ListUsersInGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`poolId`, p.PoolId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
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

type ListUsersWithRolesForEventParameters struct {
	EventId       string
	WithData      *[]string
	Query         *string
	SortBy        *string
	SortDirection *string
	Page          *int
	ItemsPerPage  *int
}

func (t *User) ListUsersWithRolesForEvent(p *ListUsersWithRolesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
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
		`/v2/User/UseCase/ListUsersWithRolesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string UserId
// @param string Token

type AccessUserForgotPasswordTokenParameters struct {
	UserId string
	Token  string
}

func (t *User) AccessUserForgotPasswordToken(p *AccessUserForgotPasswordTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`token`, p.Token)

	return t.restClient.Post(
		`/v2/User/UseCase/AccessUserForgotPasswordToken`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string GroupId

type AddUserAccessToGroupParameters struct {
	UserId  string
	GroupId string
}

func (t *User) AddUserAccessToGroup(p *AddUserAccessToGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`groupId`, p.GroupId)

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

type CreateAuthUserParameters struct {
	Email     *string
	FirstName *string
	LastName  *string
	Company   *string
	Position  *string
	Phone     *string
	PoolId    *string
	Title     *string
	Other     *string
}

func (t *User) CreateAuthUser(p *CreateAuthUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
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
	if p.Phone != nil {
		queryParameters.Add(`phone`, *p.Phone)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
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

type CreateCIOAccountParameters struct {
	Email     *string
	FirstName *string
	LastName  *string
	Company   *string
	Position  *string
	Phone     *string
	Title     *string
	Other     *string
}

func (t *User) CreateCIOAccount(p *CreateCIOAccountParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
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
	if p.Phone != nil {
		queryParameters.Add(`phone`, *p.Phone)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
	}

	return t.restClient.Post(
		`/v2/User/UseCase/CreateCIOAccount`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId

type CreateForgotPasswordTokenParameters struct {
	UserId string
}

func (t *User) CreateForgotPasswordToken(p *CreateForgotPasswordTokenParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)

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

type CreateUserParameters struct {
	Email     *string
	FirstName *string
	LastName  *string
	Company   *string
	Position  *string
	Phone     *string
	PoolId    *string
	Title     *string
	Other     *string
}

func (t *User) CreateUser(p *CreateUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
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
	if p.Phone != nil {
		queryParameters.Add(`phone`, *p.Phone)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Title != nil {
		queryParameters.Add(`title`, *p.Title)
	}
	if p.Other != nil {
		queryParameters.Add(`other`, *p.Other)
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

type CreateUserContactAgentParameters struct {
	UserId             string
	PoolId             string
	Email              string
	UserContactAgentId *string
}

func (t *User) CreateUserContactAgent(p *CreateUserContactAgentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`email`, p.Email)
	if p.UserContactAgentId != nil {
		queryParameters.Add(`userContactAgentId`, *p.UserContactAgentId)
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

type RemoveEventRoleForUserParameters struct {
	UserId  string
	EventId string
}

func (t *User) RemoveEventRoleForUser(p *RemoveEventRoleForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/User/UseCase/RemoveEventRoleForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string GroupId

type RemoveUserAccessToGroupParameters struct {
	UserId  string
	GroupId string
}

func (t *User) RemoveUserAccessToGroup(p *RemoveUserAccessToGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`groupId`, p.GroupId)

	return t.restClient.Post(
		`/v2/User/UseCase/RemoveUserAccessToGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserContactAgentId

type RemoveUserContactAgentParameters struct {
	UserContactAgentId string
}

func (t *User) RemoveUserContactAgent(p *RemoveUserContactAgentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userContactAgentId`, p.UserContactAgentId)

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

type RemoveUsersFromPoolParameters struct {
	RemoveUserIds []string
	RequestUserId string
	PoolId        string
}

func (t *User) RemoveUsersFromPool(p *RemoveUsersFromPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	for i := range p.RemoveUserIds {
		queryParameters.Add(`removeUserIds[]`, p.RemoveUserIds[i])
	}
	queryParameters.Add(`requestUserId`, p.RequestUserId)
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/User/UseCase/RemoveUsersFromPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Email
// @param string|null AppName

type SendForgotPasswordEmailParameters struct {
	Email   string
	AppName *string
}

func (t *User) SendForgotPasswordEmail(p *SendForgotPasswordEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)
	if p.AppName != nil {
		queryParameters.Add(`appName`, *p.AppName)
	}

	return t.restClient.Post(
		`/v2/User/UseCase/SendForgotPasswordEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Email
// @param string|null AppName

type SendVerificationEmailParameters struct {
	Email   string
	AppName *string
}

func (t *User) SendVerificationEmail(p *SendVerificationEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`email`, p.Email)
	if p.AppName != nil {
		queryParameters.Add(`appName`, *p.AppName)
	}

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

type SetEmailForInvitationParameters struct {
	InvitationId string
	AuthUserId   string
	ChangeUserId string
	Email        string
}

func (t *User) SetEmailForInvitation(p *SetEmailForInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`authUserId`, p.AuthUserId)
	queryParameters.Add(`changeUserId`, p.ChangeUserId)
	queryParameters.Add(`email`, p.Email)

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

type SetEventRoleForUserParameters struct {
	UserId              string
	EventId             string
	EventRole           string
	AuthenticatedUserId string
}

func (t *User) SetEventRoleForUser(p *SetEventRoleForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`eventRole`, p.EventRole)
	queryParameters.Add(`authenticatedUserId`, p.AuthenticatedUserId)

	return t.restClient.Post(
		`/v2/User/UseCase/SetEventRoleForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserContactAgentId
// @param string Email

type UpdateUserContactAgentParameters struct {
	UserContactAgentId string
	Email              string
}

func (t *User) UpdateUserContactAgent(p *UpdateUserContactAgentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userContactAgentId`, p.UserContactAgentId)
	queryParameters.Add(`email`, p.Email)

	return t.restClient.Post(
		`/v2/User/UseCase/UpdateUserContactAgent`,
		&queryParameters,
		nil,
		nil,
	)
}
