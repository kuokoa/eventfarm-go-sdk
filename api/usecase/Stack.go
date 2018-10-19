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

type Stack struct {
	restClient sdk.RestClientInterface
}

func NewStack(restClient sdk.RestClientInterface) *Stack {
	return &Stack{restClient}
}

// GET: Queries
// @param string StackId
// @param array|null WithData Event|TicketType|AvailabilityCounts

type GetStackParameters struct {
	StackId  string
	WithData *[]string
}

func (t *Stack) GetStack(p *GetStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Stack/UseCase/GetStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param array|null WithData TicketType|Event
// @param array|null ExcludeStackMethodFilter
// @param bool|null ShouldHideDeleted true|false
// @param string|null Query
// @param string|null SortBy
// @param string|null SortDirection
// @param int|null Page
// @param int|null ItemsPerPage

type ListStacksForEventParameters struct {
	EventId                  string
	WithData                 *[]string
	ExcludeStackMethodFilter *[]string
	ShouldHideDeleted        *bool
	Query                    *string
	SortBy                   *string
	SortDirection            *string
	Page                     *int
	ItemsPerPage             *int
}

func (t *Stack) ListStacksForEvent(p *ListStacksForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.ExcludeStackMethodFilter != nil {
		for i := range *p.ExcludeStackMethodFilter {
			queryParameters.Add(`excludeStackMethodFilter[]`, (*p.ExcludeStackMethodFilter)[i])
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
		`/v2/Stack/UseCase/ListStacksForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param array|null WithData TicketType|Event
// @param array|null ExcludeStackMethodFilter
// @param bool|null ShouldHideDeleted true|false
// @param string|null Query
// @param string|null SortBy
// @param string|null SortDirection
// @param int|null Page
// @param int|null ItemsPerPage

type ListStacksForPromotionParameters struct {
	PromotionId              string
	WithData                 *[]string
	ExcludeStackMethodFilter *[]string
	ShouldHideDeleted        *bool
	Query                    *string
	SortBy                   *string
	SortDirection            *string
	Page                     *int
	ItemsPerPage             *int
}

func (t *Stack) ListStacksForPromotion(p *ListStacksForPromotionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.ExcludeStackMethodFilter != nil {
		for i := range *p.ExcludeStackMethodFilter {
			queryParameters.Add(`excludeStackMethodFilter[]`, (*p.ExcludeStackMethodFilter)[i])
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
		`/v2/Stack/UseCase/ListStacksForPromotion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param int|null Page
// @param int|null ItemsPerPage

type ListStacksForTicketTypeParameters struct {
	TicketTypeId string
	Page         *int
	ItemsPerPage *int
}

func (t *Stack) ListStacksForTicketType(p *ListStacksForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Stack/UseCase/ListStacksForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PromotionId
// @param array StackIds

type AddPromotionToStacksParameters struct {
	PromotionId string
	StackIds    []string
}

func (t *Stack) AddPromotionToStacks(p *AddPromotionToStacksParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	for i := range p.StackIds {
		queryParameters.Add(`stackIds[]`, p.StackIds[i])
	}

	return t.restClient.Post(
		`/v2/Stack/UseCase/AddPromotionToStacks`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string TicketTypeId
// @param string MethodId
// @param int Quantity
// @param int MaxQty
// @param float|null Price
// @param float|null ServiceFee
// @param int|null OpeningTime
// @param int|null ClosingTime
// @param bool|null Transferable true|false
// @param string|null InviteDesignId
// @param string|null ConfirmDesignId
// @param string|null DeclineDesignId
// @param string|null StackId

type CreateStackParameters struct {
	EventId         string
	TicketTypeId    string
	MethodId        string
	Quantity        int
	MaxQty          int
	Price           *float
	ServiceFee      *float
	OpeningTime     *int
	ClosingTime     *int
	Transferable    *bool
	InviteDesignId  *string
	ConfirmDesignId *string
	DeclineDesignId *string
	StackId         *string
}

func (t *Stack) CreateStack(p *CreateStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`methodId`, p.MethodId)
	queryParameters.Add(`quantity`, strconv.Itoa(p.Quantity))
	queryParameters.Add(`maxQty`, strconv.Itoa(p.MaxQty))
	if p.Price != nil {
		queryParameters.Add(`price`, *p.Price)
	}
	if p.ServiceFee != nil {
		queryParameters.Add(`serviceFee`, *p.ServiceFee)
	}
	if p.OpeningTime != nil {
		queryParameters.Add(`openingTime`, strconv.Itoa(*p.OpeningTime))
	}
	if p.ClosingTime != nil {
		queryParameters.Add(`closingTime`, strconv.Itoa(*p.ClosingTime))
	}
	if p.Transferable != nil {
		queryParameters.Add(`transferable`, strconv.FormatBool(*p.Transferable))
	}
	if p.InviteDesignId != nil {
		queryParameters.Add(`inviteDesignId`, *p.InviteDesignId)
	}
	if p.ConfirmDesignId != nil {
		queryParameters.Add(`confirmDesignId`, *p.ConfirmDesignId)
	}
	if p.DeclineDesignId != nil {
		queryParameters.Add(`declineDesignId`, *p.DeclineDesignId)
	}
	if p.StackId != nil {
		queryParameters.Add(`stackId`, *p.StackId)
	}

	return t.restClient.Post(
		`/v2/Stack/UseCase/CreateStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId

type DeleteStackParameters struct {
	StackId string
}

func (t *Stack) DeleteStack(p *DeleteStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/DeleteStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param array StackIds

type RemovePromotionFromStacksParameters struct {
	PromotionId string
	StackIds    []string
}

func (t *Stack) RemovePromotionFromStacks(p *RemovePromotionFromStacksParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, p.PromotionId)
	for i := range p.StackIds {
		queryParameters.Add(`stackIds[]`, p.StackIds[i])
	}

	return t.restClient.Post(
		`/v2/Stack/UseCase/RemovePromotionFromStacks`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param int ClosingTime

type SetClosingTimeForStackParameters struct {
	StackId     string
	ClosingTime int
}

func (t *Stack) SetClosingTimeForStack(p *SetClosingTimeForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`closingTime`, strconv.Itoa(p.ClosingTime))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetClosingTimeForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string ConfirmDesignId

type SetConfirmDesignForStackParameters struct {
	StackId         string
	ConfirmDesignId string
}

func (t *Stack) SetConfirmDesignForStack(p *SetConfirmDesignForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`confirmDesignId`, p.ConfirmDesignId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetConfirmDesignForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string DeclineDesignId

type SetDeclineDesignForStackParameters struct {
	StackId         string
	DeclineDesignId string
}

func (t *Stack) SetDeclineDesignForStack(p *SetDeclineDesignForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`declineDesignId`, p.DeclineDesignId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetDeclineDesignForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string InviteDesignId

type SetInviteDesignForStackParameters struct {
	StackId        string
	InviteDesignId string
}

func (t *Stack) SetInviteDesignForStack(p *SetInviteDesignForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`inviteDesignId`, p.InviteDesignId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetInviteDesignForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param int MaxQuantity

type SetMaxQuantityForStackParameters struct {
	StackId     string
	MaxQuantity int
}

func (t *Stack) SetMaxQuantityForStack(p *SetMaxQuantityForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`maxQuantity`, strconv.Itoa(p.MaxQuantity))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetMaxQuantityForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string MethodSlug

type SetMethodForStackParameters struct {
	StackId    string
	MethodSlug string
}

func (t *Stack) SetMethodForStack(p *SetMethodForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`methodSlug`, p.MethodSlug)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetMethodForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param int OpeningTime

type SetOpeningTimeForStackParameters struct {
	StackId     string
	OpeningTime int
}

func (t *Stack) SetOpeningTimeForStack(p *SetOpeningTimeForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`openingTime`, strconv.Itoa(p.OpeningTime))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetOpeningTimeForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param float Price

type SetPriceForStackParameters struct {
	StackId string
	Price   float
}

func (t *Stack) SetPriceForStack(p *SetPriceForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`price`, p.Price)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetPriceForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param int Quantity

type SetQuantityForStackParameters struct {
	StackId  string
	Quantity int
}

func (t *Stack) SetQuantityForStack(p *SetQuantityForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`quantity`, strconv.Itoa(p.Quantity))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetQuantityForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param float ServiceFee

type SetServiceFeeForStackParameters struct {
	StackId    string
	ServiceFee float
}

func (t *Stack) SetServiceFeeForStack(p *SetServiceFeeForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`serviceFee`, p.ServiceFee)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetServiceFeeForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param bool Transferable true|false

type SetTransferableForStackParameters struct {
	StackId      string
	Transferable bool
}

func (t *Stack) SetTransferableForStack(p *SetTransferableForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`transferable`, strconv.FormatBool(p.Transferable))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetTransferableForStack`,
		&queryParameters,
		nil,
		nil,
	)
}
