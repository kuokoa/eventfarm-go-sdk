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

type Group struct {
	restClient gosdk.RestClientInterface
}

func NewGroup(restClient gosdk.RestClientInterface) *Group {
	return &Group{restClient}
}

// GET: Queries
// @param string GroupId
// @param array|null WithData totalUsersInGroup|creatorUser
func (t *Group) GetGroup(GroupId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *Group) ListGroupMembershipForUser(PoolId string, UserId string, GroupOwnerUserId *string, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	if GroupOwnerUserId != nil {
		queryParameters.Add(`groupOwnerUserId`, *GroupOwnerUserId)
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
func (t *Group) ListGroupsOwnedByUser(UserId string, Query *string, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	if Query != nil {
		queryParameters.Add(`query`, *Query)
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
func (t *Group) AddUsersToGroup(GroupId string, UserIds []string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	for i := range UserIds {
		queryParameters.Add(`userIds[]`, UserIds[i])
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
func (t *Group) CreateGroupForUser(UserId string, GroupName string, GroupId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`groupName`, GroupName)
	if GroupId != nil {
		queryParameters.Add(`groupId`, *GroupId)
	}

	return t.restClient.Post(
		`/v2/Group/UseCase/CreateGroupForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId
func (t *Group) DeleteGroup(GroupId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)

	return t.restClient.Post(
		`/v2/Group/UseCase/DeleteGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string DestinationGroupId
// @param array FromGroupIds
func (t *Group) MergeGroups(DestinationGroupId string, FromGroupIds []string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`destinationGroupId`, DestinationGroupId)
	for i := range FromGroupIds {
		queryParameters.Add(`fromGroupIds[]`, FromGroupIds[i])
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
func (t *Group) RemoveUsersFromGroup(GroupId string, UserIds []string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	for i := range UserIds {
		queryParameters.Add(`userIds[]`, UserIds[i])
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
func (t *Group) SetGroupName(GroupId string, UserId string, GroupName string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`groupName`, GroupName)

	return t.restClient.Post(
		`/v2/Group/UseCase/SetGroupName`,
		&queryParameters,
		nil,
		nil,
	)
}
