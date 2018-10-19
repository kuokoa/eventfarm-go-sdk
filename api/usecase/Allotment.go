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

type Allotment struct {
	restClient sdk.RestClientInterface
}

func NewAllotment(restClient sdk.RestClientInterface) *Allotment {
	return &Allotment{restClient}
}

// GET: Queries
// @param string StackId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100

type ListAllotmentsForStackParameters struct {
	StackId      string
	Page         *int
	ItemsPerPage *int
}

func (t *Allotment) ListAllotmentsForStack(p *ListAllotmentsForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Allotment/UseCase/ListAllotmentsForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string TicketBlockId
// @param string StackId
// @param int Quantity
// @param string|null AllotmentId

type CreateAllotmentParameters struct {
	TicketBlockId string
	StackId       string
	Quantity      int
	AllotmentId   *string
}

func (t *Allotment) CreateAllotment(p *CreateAllotmentParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, p.TicketBlockId)
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`quantity`, strconv.Itoa(p.Quantity))
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

// @param string AllotmentId

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

// @param string AllotmentId
// @param int Quantity >= 1

type SetAllotmentQuantityParameters struct {
	AllotmentId string
	Quantity    int
}

func (t *Allotment) SetAllotmentQuantity(p *SetAllotmentQuantityParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`allotmentId`, p.AllotmentId)
	queryParameters.Add(`quantity`, strconv.Itoa(p.Quantity))

	return t.restClient.Post(
		`/v2/Allotment/UseCase/SetAllotmentQuantity`,
		&queryParameters,
		nil,
		nil,
	)
}
