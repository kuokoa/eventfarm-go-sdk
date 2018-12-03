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
// @param string TicketBlockId

type CreateTicketBlockSummaryReportParameters struct {
	EventId       string
	OwnerUserId   string
	TicketBlockId string
}

func (t *Report) CreateTicketBlockSummaryReport(p *CreateTicketBlockSummaryReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ownerUserId`, p.OwnerUserId)
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTicketBlockSummaryReport`,
		&queryParameters,
		nil,
		nil,
	)
}
