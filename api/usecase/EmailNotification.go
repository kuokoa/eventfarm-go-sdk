/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type EmailNotification struct {
	restClient rest.RestClientInterface
}

func NewEmailNotification(restClient rest.RestClientInterface) *EmailNotification {
	return &EmailNotification{restClient}
}

// GET: Queries

type GetOpenActionsForEventOverLastMonthParameters struct {
	EventId string
}

func (t *EmailNotification) GetOpenActionsForEventOverLastMonth(p *GetOpenActionsForEventOverLastMonthParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EmailNotification/UseCase/GetOpenActionsForEventOverLastMonth`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateSparkpostNotificationParameters struct {
	EmailMessageId      string
	Type                string
	CreatedAt           int64
	EventId             *string
	EmailNotificationId *string
}

func (t *EmailNotification) CreateSparkpostNotification(p *CreateSparkpostNotificationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailMessageId`, p.EmailMessageId)
	queryParameters.Add(`type`, p.Type)
	queryParameters.Add(`createdAt`, strconv.FormatInt(p.CreatedAt, 10))
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.EmailNotificationId != nil {
		queryParameters.Add(`emailNotificationId`, *p.EmailNotificationId)
	}

	return t.restClient.Post(
		`/v2/EmailNotification/UseCase/CreateSparkpostNotification`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailNotification) CreateSparkpostNotificationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailNotification/UseCase/CreateSparkpostNotification`,
		data,
		nil,
		nil,
	)
}

type SimulateEmailNotificationsForEventParameters struct {
	EventId      string
	TotalRecords *int64
}

func (t *EmailNotification) SimulateEmailNotificationsForEvent(p *SimulateEmailNotificationsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.TotalRecords != nil {
		queryParameters.Add(`totalRecords`, strconv.FormatInt(*p.TotalRecords, 10))
	}

	return t.restClient.Post(
		`/v2/EmailNotification/UseCase/SimulateEmailNotificationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *EmailNotification) SimulateEmailNotificationsForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/EmailNotification/UseCase/SimulateEmailNotificationsForEvent`,
		data,
		nil,
		nil,
	)
}
