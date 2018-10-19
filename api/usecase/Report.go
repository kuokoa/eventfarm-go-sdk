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

type Report struct {
	restClient sdk.RestClientInterface
}

func NewReport(restClient sdk.RestClientInterface) *Report {
	return &Report{restClient}
}

// GET: Queries
// @param int StartTime
// @param int EndTime

type ReportTotalEventsRunningBetweenDatesParameters struct {
	StartTime int
	EndTime   int
}

func (t *Report) ReportTotalEventsRunningBetweenDates(p *ReportTotalEventsRunningBetweenDatesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`startTime`, strconv.Itoa(p.StartTime))
	queryParameters.Add(`endTime`, strconv.Itoa(p.EndTime))

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
