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

type Allotment struct {
	restClient gosdk.RestClientInterface
}

func NewAllotment(restClient gosdk.RestClientInterface) *Allotment {
	return &Allotment{restClient}
}

// GET: Queries
// @param string StackId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *Allotment) ListAllotmentsForStack(StackId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Allotment) CreateAllotment(TicketBlockId string, StackId string, Quantity int, AllotmentId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketBlockId`, TicketBlockId)
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`quantity`, strconv.Itoa(Quantity))
	if AllotmentId != nil {
		queryParameters.Add(`allotmentId`, *AllotmentId)
	}

	return t.restClient.Post(
		`/v2/Allotment/UseCase/CreateAllotment`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string AllotmentId
func (t *Allotment) DeleteAllotment(AllotmentId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`allotmentId`, AllotmentId)

	return t.restClient.Post(
		`/v2/Allotment/UseCase/DeleteAllotment`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string AllotmentId
// @param int Quantity >= 1
func (t *Allotment) SetAllotmentQuantity(AllotmentId string, Quantity int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`allotmentId`, AllotmentId)
	queryParameters.Add(`quantity`, strconv.Itoa(Quantity))

	return t.restClient.Post(
		`/v2/Allotment/UseCase/SetAllotmentQuantity`,
		&queryParameters,
		nil,
		nil,
	)
}
