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

type ActivityLog struct {
	restClient rest.RestClientInterface
}

func NewActivityLog(restClient rest.RestClientInterface) *ActivityLog {
	return &ActivityLog{restClient}
}

// GET: Queries

type ListEntriesForEventParameters struct {
	EventId      string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *ActivityLog) ListEntriesForEvent(p *ListEntriesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/ActivityLog/UseCase/ListEntriesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListEntriesForInvitationParameters struct {
	InvitationId string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *ActivityLog) ListEntriesForInvitation(p *ListEntriesForInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/ActivityLog/UseCase/ListEntriesForInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateActivityLogParameters struct {
	EventId       string
	Action        string
	ActionValue   string
	ActionTime    int64
	InvitationId  *string
	UserId        *string
	ActionUserId  *string
	ActivityLogId *string
}

func (t *ActivityLog) CreateActivityLog(p *CreateActivityLogParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`action`, p.Action)
	queryParameters.Add(`actionValue`, p.ActionValue)
	queryParameters.Add(`actionTime`, strconv.FormatInt(p.ActionTime, 10))
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}
	if p.UserId != nil {
		queryParameters.Add(`userId`, *p.UserId)
	}
	if p.ActionUserId != nil {
		queryParameters.Add(`actionUserId`, *p.ActionUserId)
	}
	if p.ActivityLogId != nil {
		queryParameters.Add(`activityLogId`, *p.ActivityLogId)
	}

	return t.restClient.Post(
		`/v2/ActivityLog/UseCase/CreateActivityLog`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateActivityLogAndDetailParameters struct {
	EventId          string
	Action           string
	ActionValue      string
	ActionTime       int64
	Content          string
	InvitationId     *string
	UserId           *string
	ActionUserId     *string
	ActivityLogId    *string
	ActivityDetailId *string
}

func (t *ActivityLog) CreateActivityLogAndDetail(p *CreateActivityLogAndDetailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`action`, p.Action)
	queryParameters.Add(`actionValue`, p.ActionValue)
	queryParameters.Add(`actionTime`, strconv.FormatInt(p.ActionTime, 10))
	queryParameters.Add(`content`, p.Content)
	if p.InvitationId != nil {
		queryParameters.Add(`invitationId`, *p.InvitationId)
	}
	if p.UserId != nil {
		queryParameters.Add(`userId`, *p.UserId)
	}
	if p.ActionUserId != nil {
		queryParameters.Add(`actionUserId`, *p.ActionUserId)
	}
	if p.ActivityLogId != nil {
		queryParameters.Add(`activityLogId`, *p.ActivityLogId)
	}
	if p.ActivityDetailId != nil {
		queryParameters.Add(`activityDetailId`, *p.ActivityDetailId)
	}

	return t.restClient.Post(
		`/v2/ActivityLog/UseCase/CreateActivityLogAndDetail`,
		&queryParameters,
		nil,
		nil,
	)
}
