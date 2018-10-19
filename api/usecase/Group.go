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

type Group struct {
	restClient sdk.RestClientInterface
}

func NewGroup(restClient sdk.RestClientInterface) *Group {
	return &Group{restClient}
}

// GET: Queries
// @param string GroupId
// @param array|null WithData totalUsersInGroup|creatorUser

type GetGroupParameters struct {
	GroupId  string
	WithData *[]string
}

func (t *Group) GetGroup(p *GetGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Group/UseCase/GetGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string UserId
// @param string|null GroupOwnerUserId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
// @param string|null SortBy
// @param string|null SortDirection ascending|descending

type ListGroupMembershipForUserParameters struct {
	PoolId           string
	UserId           string
	GroupOwnerUserId *string
	Page             *int
	ItemsPerPage     *int
	SortBy           *string
	SortDirection    *string
}

func (t *Group) ListGroupMembershipForUser(p *ListGroupMembershipForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	if p.GroupOwnerUserId != nil {
		queryParameters.Add(`groupOwnerUserId`, *p.GroupOwnerUserId)
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

	return t.restClient.Get(
		`/v2/Group/UseCase/ListGroupMembershipForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string|null Query
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
// @param string|null SortBy
// @param string|null SortDirection ascending|descending

type ListGroupsOwnedByUserParameters struct {
	UserId        string
	Query         *string
	Page          *int
	ItemsPerPage  *int
	SortBy        *string
	SortDirection *string
}

func (t *Group) ListGroupsOwnedByUser(p *ListGroupsOwnedByUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
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

	return t.restClient.Get(
		`/v2/Group/UseCase/ListGroupsOwnedByUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string GroupId
// @param array UserIds

type AddUsersToGroupParameters struct {
	GroupId string
	UserIds []string
}

func (t *Group) AddUsersToGroup(p *AddUsersToGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	for i := range p.UserIds {
		queryParameters.Add(`userIds[]`, p.UserIds[i])
	}

	return t.restClient.Post(
		`/v2/Group/UseCase/AddUsersToGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string GroupName
// @param string|null GroupId

type CreateGroupForUserParameters struct {
	UserId    string
	GroupName string
	GroupId   *string
}

func (t *Group) CreateGroupForUser(p *CreateGroupForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`groupName`, p.GroupName)
	if p.GroupId != nil {
		queryParameters.Add(`groupId`, *p.GroupId)
	}

	return t.restClient.Post(
		`/v2/Group/UseCase/CreateGroupForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId

type DeleteGroupParameters struct {
	GroupId string
}

func (t *Group) DeleteGroup(p *DeleteGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)

	return t.restClient.Post(
		`/v2/Group/UseCase/DeleteGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string DestinationGroupId
// @param array FromGroupIds

type MergeGroupsParameters struct {
	DestinationGroupId string
	FromGroupIds       []string
}

func (t *Group) MergeGroups(p *MergeGroupsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`destinationGroupId`, p.DestinationGroupId)
	for i := range p.FromGroupIds {
		queryParameters.Add(`fromGroupIds[]`, p.FromGroupIds[i])
	}

	return t.restClient.Post(
		`/v2/Group/UseCase/MergeGroups`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId
// @param array UserIds

type RemoveUsersFromGroupParameters struct {
	GroupId string
	UserIds []string
}

func (t *Group) RemoveUsersFromGroup(p *RemoveUsersFromGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	for i := range p.UserIds {
		queryParameters.Add(`userIds[]`, p.UserIds[i])
	}

	return t.restClient.Post(
		`/v2/Group/UseCase/RemoveUsersFromGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId
// @param string UserId
// @param string GroupName

type SetGroupNameParameters struct {
	GroupId   string
	UserId    string
	GroupName string
}

func (t *Group) SetGroupName(p *SetGroupNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`groupName`, p.GroupName)

	return t.restClient.Post(
		`/v2/Group/UseCase/SetGroupName`,
		&queryParameters,
		nil,
		nil,
	)
}
