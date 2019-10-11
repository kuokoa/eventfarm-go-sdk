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

type EmailMessage struct {
	restClient rest.RestClientInterface
}

func NewEmailMessage(restClient rest.RestClientInterface) *EmailMessage {
	return &EmailMessage{restClient}
}

// GET: Queries

type GetEmailMessageParameters struct {
	EmailMessageId string
}

func (t *EmailMessage) GetEmailMessage(p *GetEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailMessageId`, p.EmailMessageId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetEmailMessageStatsForEventParameters struct {
	EventId string
}

func (t *EmailMessage) GetEmailMessageStatsForEvent(p *GetEmailMessageStatsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetEmailMessageStatsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetScheduledEmailMessageParameters struct {
	MongoQueueId string
}

func (t *EmailMessage) GetScheduledEmailMessage(p *GetScheduledEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, p.MongoQueueId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetScheduledEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListOutboxEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *EmailMessage) ListOutboxEmailMessagesByEvent(p *ListOutboxEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListOutboxEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListScheduledBatchEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *EmailMessage) ListScheduledBatchEmailMessagesByEvent(p *ListScheduledBatchEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListScheduledBatchEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListScheduledEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *EmailMessage) ListScheduledEmailMessagesByEvent(p *ListScheduledEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListScheduledEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListSentEmailMessagesByEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *EmailMessage) ListSentEmailMessagesByEvent(p *ListSentEmailMessagesByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListSentEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreatePreviewEmailMessageParameters struct {
	EventId         string
	OwnerUserId     string
	MessageContent  string
	ToEmails        []string
	FromName        string
	Subject         string
	LayoutType      string
	MessageType     string
	BackgroundColor string
	FromEmail       *string
	ReplyEmail      *string
	DomainMaskEmail *string
}

func (t *EmailMessage) CreatePreviewEmailMessage(p *CreatePreviewEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContent`, p.MessageContent)
	for i := range p.ToEmails {
		queryParameters.Add(`toEmails[]`, p.ToEmails[i])
	}
	queryParameters.Add(`fromName`, p.FromName)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`layoutType`, p.LayoutType)
	queryParameters.Add(`messageType`, p.MessageType)
	queryParameters.Add(`backgroundColor`, p.BackgroundColor)
	if p.FromEmail != nil {
		queryParameters.Add(`fromEmail`, *p.FromEmail)
	}
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/CreatePreviewEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateScheduledBatchEmailMessageParameters struct {
	EventId         string
	OwnerUserId     string
	EmailDesignId   string
	Type            string
	Targets         []string // confirmed | purchased | assigned | unconfirmed | attended
	MessageSendTime string
	Timezone        string
}

func (t *EmailMessage) CreateScheduledBatchEmailMessage(p *CreateScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	queryParameters.Add(`type`, p.Type)
	for i := range p.Targets {
		queryParameters.Add(`targets[]`, p.Targets[i])
	}
	queryParameters.Add(`messageSendTime`, p.MessageSendTime)
	queryParameters.Add(`timezone`, p.Timezone)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/CreateScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type RemoveScheduledBatchEmailMessageParameters struct {
	MongoQueueId string
}

func (t *EmailMessage) RemoveScheduledBatchEmailMessage(p *RemoveScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, p.MongoQueueId)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/RemoveScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

type SendAMessagePreviewParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessagePreview(p *SendAMessagePreviewParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessagePreview`,
		&queryParameters,
		nil,
		nil,
	)
}

type SendAMessageToAllCheckedInGuestsParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToAllCheckedInGuests(p *SendAMessageToAllCheckedInGuestsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToAllCheckedInGuests`,
		&queryParameters,
		nil,
		nil,
	)
}

type SendAMessageToEventWaitListParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToEventWaitList(p *SendAMessageToEventWaitListParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToEventWaitList`,
		&queryParameters,
		nil,
		nil,
	)
}

type SendAMessageToGroupParameters struct {
	GroupId                string
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToGroup(p *SendAMessageToGroupParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, p.GroupId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

type SendAMessageToNoShowsParameters struct {
	EventId                string
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToNoShows(p *SendAMessageToNoShowsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToNoShows`,
		&queryParameters,
		nil,
		nil,
	)
}

type SendAMessageToTicketTypesParameters struct {
	EventId                string
	TicketTypeIds          []string
	InvitationStatusTypes  []string // assigned | purchased | confirmed-by-rsvp | declined-by-rsvp | left-behind | not-yet-purchased | registered | unconfirmed | recycled | not-yet-registered | waitlisted
	OwnerUserId            string
	MessageContents        string
	Subject                string
	FromName               string
	ReplyEmail             *string
	CcEmails               *[]string
	BccEmails              *[]string
	UseEventSpecificLayout *bool
	BackgroundColor        *string
	DomainMaskEmail        *string
}

func (t *EmailMessage) SendAMessageToTicketTypes(p *SendAMessageToTicketTypesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	for i := range p.TicketTypeIds {
		queryParameters.Add(`ticketTypeIds[]`, p.TicketTypeIds[i])
	}
	for i := range p.InvitationStatusTypes {
		queryParameters.Add(`invitationStatusTypes[]`, p.InvitationStatusTypes[i])
	}
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`messageContents`, p.MessageContents)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`fromName`, p.FromName)
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*p.UseEventSpecificLayout))
	}
	if p.BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *p.BackgroundColor)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToTicketTypes`,
		&queryParameters,
		nil,
		nil,
	)
}

type SimulateEmailMessageStatsForEventParameters struct {
	EventId string
}

func (t *EmailMessage) SimulateEmailMessageStatsForEvent(p *SimulateEmailMessageStatsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SimulateEmailMessageStatsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type UpdateScheduledBatchEmailMessageParameters struct {
	MongoQueueId    string
	EmailDesignId   string
	Targets         []string // confirmed | purchased | assigned | unconfirmed | attended
	MessageSendTime string
	Timezone        string
}

func (t *EmailMessage) UpdateScheduledBatchEmailMessage(p *UpdateScheduledBatchEmailMessageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, p.MongoQueueId)
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	for i := range p.Targets {
		queryParameters.Add(`targets[]`, p.Targets[i])
	}
	queryParameters.Add(`messageSendTime`, p.MessageSendTime)
	queryParameters.Add(`timezone`, p.Timezone)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/UpdateScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}
