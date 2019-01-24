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

type Stack struct {
	restClient rest.RestClientInterface
}

func NewStack(restClient rest.RestClientInterface) *Stack {
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
// @param array|null WithData TicketType|Event|availibilityCounts
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
	Page                     *int64
	ItemsPerPage             *int64
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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
	Page                     *int64
	ItemsPerPage             *int64
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
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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
	Page         *int64
	ItemsPerPage *int64
}

func (t *Stack) ListStacksForTicketType(p *ListStacksForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
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
	Quantity        int64
	MaxQty          int64
	Price           *float64
	ServiceFee      *float64
	OpeningTime     *int64
	ClosingTime     *int64
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
	queryParameters.Add(`quantity`, strconv.FormatInt(p.Quantity, 10))
	queryParameters.Add(`maxQty`, strconv.FormatInt(p.MaxQty, 10))
	if p.Price != nil {
		queryParameters.Add(`price`, fmt.Sprintf("%f", *p.Price))
	}
	if p.ServiceFee != nil {
		queryParameters.Add(`serviceFee`, fmt.Sprintf("%f", *p.ServiceFee))
	}
	if p.OpeningTime != nil {
		queryParameters.Add(`openingTime`, strconv.FormatInt(*p.OpeningTime, 10))
	}
	if p.ClosingTime != nil {
		queryParameters.Add(`closingTime`, strconv.FormatInt(*p.ClosingTime, 10))
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

// @param string EventId
// @param string TicketTypeId
// @param bool PrivateInvite true|false
// @param bool Fcfs true|false
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

type CreateStackFromSettingsParameters struct {
	EventId         string
	TicketTypeId    string
	PrivateInvite   bool
	Fcfs            bool
	Quantity        int64
	MaxQty          int64
	Price           *float64
	ServiceFee      *float64
	OpeningTime     *int64
	ClosingTime     *int64
	Transferable    *bool
	InviteDesignId  *string
	ConfirmDesignId *string
	DeclineDesignId *string
	StackId         *string
}

func (t *Stack) CreateStackFromSettings(p *CreateStackFromSettingsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`privateInvite`, strconv.FormatBool(p.PrivateInvite))
	queryParameters.Add(`fcfs`, strconv.FormatBool(p.Fcfs))
	queryParameters.Add(`quantity`, strconv.FormatInt(p.Quantity, 10))
	queryParameters.Add(`maxQty`, strconv.FormatInt(p.MaxQty, 10))
	if p.Price != nil {
		queryParameters.Add(`price`, fmt.Sprintf("%f", *p.Price))
	}
	if p.ServiceFee != nil {
		queryParameters.Add(`serviceFee`, fmt.Sprintf("%f", *p.ServiceFee))
	}
	if p.OpeningTime != nil {
		queryParameters.Add(`openingTime`, strconv.FormatInt(*p.OpeningTime, 10))
	}
	if p.ClosingTime != nil {
		queryParameters.Add(`closingTime`, strconv.FormatInt(*p.ClosingTime, 10))
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
		`/v2/Stack/UseCase/CreateStackFromSettings`,
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
	ClosingTime int64
}

func (t *Stack) SetClosingTimeForStack(p *SetClosingTimeForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`closingTime`, strconv.FormatInt(p.ClosingTime, 10))

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
	MaxQuantity int64
}

func (t *Stack) SetMaxQuantityForStack(p *SetMaxQuantityForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`maxQuantity`, strconv.FormatInt(p.MaxQuantity, 10))

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
	OpeningTime int64
}

func (t *Stack) SetOpeningTimeForStack(p *SetOpeningTimeForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`openingTime`, strconv.FormatInt(p.OpeningTime, 10))

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
	Price   float64
}

func (t *Stack) SetPriceForStack(p *SetPriceForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`price`, fmt.Sprintf("%f", p.Price))

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
	Quantity int64
}

func (t *Stack) SetQuantityForStack(p *SetQuantityForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`quantity`, strconv.FormatInt(p.Quantity, 10))

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
	ServiceFee float64
}

func (t *Stack) SetServiceFeeForStack(p *SetServiceFeeForStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	queryParameters.Add(`serviceFee`, fmt.Sprintf("%f", p.ServiceFee))

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

// @param string StackId
// @param string|null MethodSlug
// @param float|null Price
// @param float|null ServiceFee
// @param int|null Quantity
// @param int|null MaxQuantity
// @param bool|null Transferable true|false

type UpdateStackParameters struct {
	StackId      string
	MethodSlug   *string
	Price        *float64
	ServiceFee   *float64
	Quantity     *int64
	MaxQuantity  *int64
	Transferable *bool
}

func (t *Stack) UpdateStack(p *UpdateStackParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	if p.MethodSlug != nil {
		queryParameters.Add(`methodSlug`, *p.MethodSlug)
	}
	if p.Price != nil {
		queryParameters.Add(`price`, fmt.Sprintf("%f", *p.Price))
	}
	if p.ServiceFee != nil {
		queryParameters.Add(`serviceFee`, fmt.Sprintf("%f", *p.ServiceFee))
	}
	if p.Quantity != nil {
		queryParameters.Add(`quantity`, strconv.FormatInt(*p.Quantity, 10))
	}
	if p.MaxQuantity != nil {
		queryParameters.Add(`maxQuantity`, strconv.FormatInt(*p.MaxQuantity, 10))
	}
	if p.Transferable != nil {
		queryParameters.Add(`transferable`, strconv.FormatBool(*p.Transferable))
	}

	return t.restClient.Post(
		`/v2/Stack/UseCase/UpdateStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string|null EventId
// @param string|null TicketTypeId
// @param bool|null PrivateInvite true|false
// @param bool|null Fcfs true|false
// @param int|null Quantity
// @param int|null MaxQty
// @param float|null Price
// @param float|null ServiceFee
// @param int|null OpeningTime
// @param int|null ClosingTime
// @param bool|null Transferable true|false
// @param string|null InviteDesignId
// @param string|null ConfirmDesignId
// @param string|null DeclineDesignId

type UpdateStackFromSettingsParameters struct {
	StackId         string
	EventId         *string
	TicketTypeId    *string
	PrivateInvite   *bool
	Fcfs            *bool
	Quantity        *int64
	MaxQty          *int64
	Price           *float64
	ServiceFee      *float64
	OpeningTime     *int64
	ClosingTime     *int64
	Transferable    *bool
	InviteDesignId  *string
	ConfirmDesignId *string
	DeclineDesignId *string
}

func (t *Stack) UpdateStackFromSettings(p *UpdateStackFromSettingsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, p.StackId)
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}
	if p.TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *p.TicketTypeId)
	}
	if p.PrivateInvite != nil {
		queryParameters.Add(`privateInvite`, strconv.FormatBool(*p.PrivateInvite))
	}
	if p.Fcfs != nil {
		queryParameters.Add(`fcfs`, strconv.FormatBool(*p.Fcfs))
	}
	if p.Quantity != nil {
		queryParameters.Add(`quantity`, strconv.FormatInt(*p.Quantity, 10))
	}
	if p.MaxQty != nil {
		queryParameters.Add(`maxQty`, strconv.FormatInt(*p.MaxQty, 10))
	}
	if p.Price != nil {
		queryParameters.Add(`price`, fmt.Sprintf("%f", *p.Price))
	}
	if p.ServiceFee != nil {
		queryParameters.Add(`serviceFee`, fmt.Sprintf("%f", *p.ServiceFee))
	}
	if p.OpeningTime != nil {
		queryParameters.Add(`openingTime`, strconv.FormatInt(*p.OpeningTime, 10))
	}
	if p.ClosingTime != nil {
		queryParameters.Add(`closingTime`, strconv.FormatInt(*p.ClosingTime, 10))
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

	return t.restClient.Post(
		`/v2/Stack/UseCase/UpdateStackFromSettings`,
		&queryParameters,
		nil,
		nil,
	)
}
