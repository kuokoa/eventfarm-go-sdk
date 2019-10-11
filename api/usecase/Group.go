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

type Group struct {
	restClient rest.RestClientInterface
}

func NewGroup(restClient rest.RestClientInterface) *Group {
	return &Group{restClient}
}

// GET: Queries

type GetGroupParameters struct {
	GroupId  string
	WithData *[]string // totalUsersInGroup | creatorUser
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

type ListGroupMembershipForUserParameters struct {
	PoolId           string
	UserId           string
	GroupOwnerUserId *string
	Page             *int64 // >= 1
	ItemsPerPage     *int64 // 1-500
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

	return t.restClient.Get(
		`/v2/Group/UseCase/ListGroupMembershipForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListGroupsOwnedByUserParameters struct {
	UserId        string
	Query         *string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-500
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

	return t.restClient.Get(
		`/v2/Group/UseCase/ListGroupsOwnedByUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

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
