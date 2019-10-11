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

type BugReport struct {
	restClient rest.RestClientInterface
}

func NewBugReport(restClient rest.RestClientInterface) *BugReport {
	return &BugReport{restClient}
}

// GET: Queries

type GetBugReportParameters struct {
	BugReportId string
	WithData    *[]string // User
}

func (t *BugReport) GetBugReport(p *GetBugReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`bugReportId`, p.BugReportId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/BugReport/UseCase/GetBugReport`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListBugReportsParameters struct {
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-500
}

func (t *BugReport) ListBugReports(p *ListBugReportsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/BugReport/UseCase/ListBugReports`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateBugReportParameters struct {
	UserId      string
	Action      string
	Message     string
	Request     string
	Response    string
	BugReportId *string
}

func (t *BugReport) CreateBugReport(p *CreateBugReportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`action`, p.Action)
	queryParameters.Add(`message`, p.Message)
	queryParameters.Add(`request`, p.Request)
	queryParameters.Add(`response`, p.Response)
	if p.BugReportId != nil {
		queryParameters.Add(`bugReportId`, *p.BugReportId)
	}

	return t.restClient.Post(
		`/v2/BugReport/UseCase/CreateBugReport`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *BugReport) CreateBugReportWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/BugReport/UseCase/CreateBugReport`,
		data,
		nil,
		nil,
	)
}
