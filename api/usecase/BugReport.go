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

type BugReport struct {
	restClient gosdk.RestClientInterface
}

func NewBugReport(restClient gosdk.RestClientInterface) *BugReport {
	return &BugReport{restClient}
}

// GET: Queries
// @param string BugReportId
// @param array|null WithData User
func (t *BugReport) GetBugReport(BugReportId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`bugReportId`, BugReportId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/BugReport/UseCase/GetBugReport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
func (t *BugReport) ListBugReports(Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/BugReport/UseCase/ListBugReports`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string UserId
// @param string Action
// @param string Message
// @param string Request
// @param string Response
// @param string|null BugReportId
func (t *BugReport) CreateBugReport(UserId string, Action string, Message string, Request string, Response string, BugReportId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`action`, Action)
	queryParameters.Add(`message`, Message)
	queryParameters.Add(`request`, Request)
	queryParameters.Add(`response`, Response)
	if BugReportId != nil {
		queryParameters.Add(`bugReportId`, *BugReportId)
	}

	return t.restClient.Post(
		`/v2/BugReport/UseCase/CreateBugReport`,
		&queryParameters,
		nil,
		nil,
	)
}
