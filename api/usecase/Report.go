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

type Report struct {
	restClient rest.RestClientInterface
}

func NewReport(restClient rest.RestClientInterface) *Report {
	return &Report{restClient}
}

// GET: Queries
// @param int StartTime
// @param int EndTime

type ReportTotalEventsRunningBetweenDatesParameters struct {
	StartTime int64
	EndTime   int64
}

func (t *Report) ReportTotalEventsRunningBetweenDates(p *ReportTotalEventsRunningBetweenDatesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`startTime`, strconv.FormatInt(p.StartTime, 10))
	queryParameters.Add(`endTime`, strconv.FormatInt(p.EndTime, 10))

	return t.restClient.Get(
		`/v2/Report/UseCase/ReportTotalEventsRunningBetweenDates`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EventId
// @param string OwnerUserId
// @param string|null Name

type CreateAdminEventReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
}

func (t *Report) CreateAdminEventReport(p *CreateAdminEventReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateAdminEventReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string|null Name

type CreateConfirmedGuestReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
}

func (t *Report) CreateConfirmedGuestReport(p *CreateConfirmedGuestReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateConfirmedGuestReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateEmailDeliverabilityReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateEmailDeliverabilityReport(p *CreateEmailDeliverabilityReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateEmailDeliverabilityReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateGraphicalCheckinReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateGraphicalCheckinReport(p *CreateGraphicalCheckinReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateGraphicalCheckinReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateGuestSummaryReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateGuestSummaryReport(p *CreateGuestSummaryReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateGuestSummaryReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreatePurchasedReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreatePurchasedReport(p *CreatePurchasedReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreatePurchasedReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateSentEmailReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateSentEmailReport(p *CreateSentEmailReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateSentEmailReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string TicketBlockId
// @param string Name

type CreateTicketBlockSummaryReportParameters struct {
	EventId       string
	OwnerUserId   string
	TicketBlockId string
	Name          string
}

func (t *Report) CreateTicketBlockSummaryReport(p *CreateTicketBlockSummaryReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTicketBlockSummaryReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateTicketBlockUserReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateTicketBlockUserReport(p *CreateTicketBlockUserReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTicketBlockUserReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateTransferReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateTransferReport(p *CreateTransferReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTransferReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateUnconfirmedGuestReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateUnconfirmedGuestReport(p *CreateUnconfirmedGuestReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateUnconfirmedGuestReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateWaitListReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateWaitListReport(p *CreateWaitListReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateWaitListReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string Name

type CreateWaiverReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        string
}

func (t *Report) CreateWaiverReport(p *CreateWaiverReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateWaiverReport`,
		&queryParameters,
		nil,
		nil,
	)
}
