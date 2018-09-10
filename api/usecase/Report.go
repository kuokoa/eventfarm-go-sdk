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

type Report struct {
	restClient gosdk.RestClientInterface
}

func NewReport(restClient gosdk.RestClientInterface) *Report {
	return &Report{restClient}
}

// GET: Queries
// @param int StartTime
// @param int EndTime
func (t *Report) ReportTotalEventsRunningBetweenDates(StartTime int, EndTime int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`startTime`, strconv.Itoa(StartTime))
	queryParameters.Add(`endTime`, strconv.Itoa(EndTime))

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
func (t *Report) CreateTicketBlockSummaryReport(EventId string, OwnerUserId string, TicketBlockId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`ticketBlockId`, TicketBlockId)

	return t.restClient.Post(
		`/v2/Report/UseCase/CreateTicketBlockSummaryReport`,
		&queryParameters,
		nil,
		nil,
	)
}
