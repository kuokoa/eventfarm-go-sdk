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

type Question struct {
	restClient rest.RestClientInterface
}

func NewQuestion(restClient rest.RestClientInterface) *Question {
	return &Question{restClient}
}

// GET: Queries
// @param string EventId
// @param array|null WithData TicketType|Answers
// @param bool|null ShouldHideDeleted true|false
// @param string|null Query
// @param string|null SortBy
// @param string|null SortDirection
// @param int|null Page
// @param int|null ItemsPerPage

type ListQuestionsForEventParameters struct {
	EventId           string
	WithData          *[]string
	ShouldHideDeleted *bool
	Query             *string
	SortBy            *string
	SortDirection     *string
	Page              *int
	ItemsPerPage      *int
}

func (t *Question) ListQuestionsForEvent(p *ListQuestionsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*p.ShouldHideDeleted))
	}
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Question/UseCase/ListQuestionsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
