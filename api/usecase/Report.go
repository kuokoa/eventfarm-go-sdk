/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"github.com/eventfarm/go-sdk/rest"
	"net/http"
	"net/url"
	"strconv"
)

type Report struct {
	restClient rest.RestClientInterface
}

func NewReport(restClient rest.RestClientInterface) *Report {
	return &Report{restClient}
}

// GET: Queries

type ListReportsForEventParameters struct {
	EventId       string
	Page          *int64 // >= 1
	ItemsPerPage  *int64 // 1-100
	SortBy        *string
	SortDirection *string
}

func (t *Report) ListReportsForEvent(p *ListReportsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
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
		`/v2/Report/UseCase/ListReportsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

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

type CreateActivityLogReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateActivityLogReport(p *CreateActivityLogReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateActivityLogReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateAdminEventReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateAdminEventReport(p *CreateAdminEventReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateAdminEventReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateConfirmedGuestReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateConfirmedGuestReport(p *CreateConfirmedGuestReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateConfirmedGuestReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateEmailDeliverabilityReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateEmailDeliverabilityReport(p *CreateEmailDeliverabilityReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateEmailDeliverabilityReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateGraphicalCheckinReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateGraphicalCheckinReport(p *CreateGraphicalCheckinReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateGraphicalCheckinReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateGuestSummaryReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateGuestSummaryReport(p *CreateGuestSummaryReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateGuestSummaryReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreatePurchasedReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreatePurchasedReport(p *CreatePurchasedReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreatePurchasedReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateSentEmailReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateSentEmailReport(p *CreateSentEmailReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateSentEmailReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateTicketBlockSummaryReportParameters struct {
	EventId       string
	OwnerUserId   string
	TicketBlockId string
	Name          *string
	ReportId      *string
}

func (t *Report) CreateTicketBlockSummaryReport(p *CreateTicketBlockSummaryReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTicketBlockSummaryReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateTicketBlockUserReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateTicketBlockUserReport(p *CreateTicketBlockUserReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTicketBlockUserReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateTransferReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateTransferReport(p *CreateTransferReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTransferReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateUnconfirmedGuestReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateUnconfirmedGuestReport(p *CreateUnconfirmedGuestReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateUnconfirmedGuestReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateWaitListReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateWaitListReport(p *CreateWaitListReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateWaitListReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type CreateWaiverReportParameters struct {
	EventId     string
	OwnerUserId string
	Name        *string
	ReportId    *string
}

func (t *Report) CreateWaiverReport(p *CreateWaiverReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	if p.Name != nil {
		queryParameters.Add(`name`, *p.Name)
	}
	if p.ReportId != nil {
		queryParameters.Add(`reportId`, *p.ReportId)
	}

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateWaiverReport`,
		&queryParameters,
		nil,
		nil,
	)
}
