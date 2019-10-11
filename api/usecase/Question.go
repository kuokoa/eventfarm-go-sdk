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

type ListQuestionsForEventParameters struct {
	EventId           string
	WithData          *[]interface{} // TicketType | Answers
	ShouldHideDeleted *bool
	Query             *string
	SortBy            *string
	SortDirection     *string
	Page              *int64
	ItemsPerPage      *int64
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Question/UseCase/ListQuestionsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
