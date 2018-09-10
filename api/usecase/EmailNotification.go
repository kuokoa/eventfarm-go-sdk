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

type EmailNotification struct {
	restClient gosdk.RestClientInterface
}

func NewEmailNotification(restClient gosdk.RestClientInterface) *EmailNotification {
	return &EmailNotification{restClient}
}

// GET: Queries
// @param string EventId
func (t *EmailNotification) GetOpenActionsForEventOverLastMonth(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/EmailNotification/UseCase/GetOpenActionsForEventOverLastMonth`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EmailMessageId
// @param string Type
// @param int CreatedAt
// @param string|null EventId
// @param string|null EmailNotificationId
func (t *EmailNotification) CreateSparkpostNotification(EmailMessageId string, Type string, CreatedAt int, EventId *string, EmailNotificationId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailMessageId`, EmailMessageId)
	queryParameters.Add(`type`, Type)
	queryParameters.Add(`createdAt`, strconv.Itoa(CreatedAt))
	if EventId != nil {
		queryParameters.Add(`eventId`, *EventId)
	}
	if EmailNotificationId != nil {
		queryParameters.Add(`emailNotificationId`, *EmailNotificationId)
	}

	return t.restClient.Post(
		`/v2/EmailNotification/UseCase/CreateSparkpostNotification`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null TotalRecords
func (t *EmailNotification) SimulateEmailNotificationsForEvent(EventId string, TotalRecords *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if TotalRecords != nil {
		queryParameters.Add(`totalRecords`, strconv.Itoa(*TotalRecords))
	}

	return t.restClient.Post(
		`/v2/EmailNotification/UseCase/SimulateEmailNotificationsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}
