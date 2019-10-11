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

type Allotment struct {
	restClient rest.RestClientInterface
}

func NewAllotment(restClient rest.RestClientInterface) *Allotment {
	return &Allotment{restClient}
}

// GET: Queries

type ListAllotmentsForStackParameters struct {
	StackId      string
	Page         *int64 // >= 1
	ItemsPerPage *int64 // 1-100
}

func (t *Allotment) ListAllotmentsForStack(p *ListAllotmentsForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/Allotment/UseCase/ListAllotmentsForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateAllotmentParameters struct {
	TicketBlockId string
	StackId       string
	Quantity      int64
	AllotmentId   *string
}

func (t *Allotment) CreateAllotment(p *CreateAllotmentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`quantity`, strconv.FormatInt(p.Quantity, 10))
	if p.AllotmentId != nil {
		queryParameters.Add(`allotmentId`, *p.AllotmentId)
	}

	return t.restClient.Post(
		`/v2/Allotment/UseCase/CreateAllotment`,
		&queryParameters,
		nil,
		nil,
	)
}

type DeleteAllotmentParameters struct {
	AllotmentId string
}

func (t *Allotment) DeleteAllotment(p *DeleteAllotmentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`allotmentId`, p.AllotmentId)

	return t.restClient.Post(
		`/v2/Allotment/UseCase/DeleteAllotment`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetAllotmentQuantityParameters struct {
	AllotmentId string
	Quantity    int64 // >= 1
}

func (t *Allotment) SetAllotmentQuantity(p *SetAllotmentQuantityParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`allotmentId`, p.AllotmentId)
	queryParameters.Add(`quantity`, strconv.FormatInt(p.Quantity, 10))

	return t.restClient.Post(
		`/v2/Allotment/UseCase/SetAllotmentQuantity`,
		&queryParameters,
		nil,
		nil,
	)
}
