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

type Stack struct {
	restClient gosdk.RestClientInterface
}

func NewStack(restClient gosdk.RestClientInterface) *Stack {
	return &Stack{restClient}
}

// GET: Queries
// @param string StackId
// @param array|null WithData Event|TicketType|AvailabilityCounts
func (t *Stack) GetStack(StackId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *Stack) ListStacksForEvent(EventId string, WithData *[]string, ExcludeStackMethodFilter *[]string, ShouldHideDeleted *bool, Query *string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if ExcludeStackMethodFilter != nil {
		for i := range *ExcludeStackMethodFilter {
			queryParameters.Add(`excludeStackMethodFilter[]`, (*ExcludeStackMethodFilter)[i])
		}
	}
	if ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*ShouldHideDeleted))
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Stack) ListStacksForPromotion(PromotionId string, WithData *[]string, ExcludeStackMethodFilter *[]string, ShouldHideDeleted *bool, Query *string, SortBy *string, SortDirection *string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if ExcludeStackMethodFilter != nil {
		for i := range *ExcludeStackMethodFilter {
			queryParameters.Add(`excludeStackMethodFilter[]`, (*ExcludeStackMethodFilter)[i])
		}
	}
	if ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*ShouldHideDeleted))
	}
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Stack) ListStacksForTicketType(TicketTypeId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, TicketTypeId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *Stack) AddPromotionToStacks(PromotionId string, StackIds []string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	for i := range StackIds {
		queryParameters.Add(`stackIds[]`, StackIds[i])
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
func (t *Stack) CreateStack(EventId string, TicketTypeId string, MethodId string, Quantity int, MaxQty int, Price *float, ServiceFee *float, OpeningTime *int, ClosingTime *int, Transferable *bool, InviteDesignId *string, ConfirmDesignId *string, DeclineDesignId *string, StackId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ticketTypeId`, TicketTypeId)
	queryParameters.Add(`methodId`, MethodId)
	queryParameters.Add(`quantity`, strconv.Itoa(Quantity))
	queryParameters.Add(`maxQty`, strconv.Itoa(MaxQty))
	if Price != nil {
		queryParameters.Add(`price`, *Price)
	}
	if ServiceFee != nil {
		queryParameters.Add(`serviceFee`, *ServiceFee)
	}
	if OpeningTime != nil {
		queryParameters.Add(`openingTime`, strconv.Itoa(*OpeningTime))
	}
	if ClosingTime != nil {
		queryParameters.Add(`closingTime`, strconv.Itoa(*ClosingTime))
	}
	if Transferable != nil {
		queryParameters.Add(`transferable`, strconv.FormatBool(*Transferable))
	}
	if InviteDesignId != nil {
		queryParameters.Add(`inviteDesignId`, *InviteDesignId)
	}
	if ConfirmDesignId != nil {
		queryParameters.Add(`confirmDesignId`, *ConfirmDesignId)
	}
	if DeclineDesignId != nil {
		queryParameters.Add(`declineDesignId`, *DeclineDesignId)
	}
	if StackId != nil {
		queryParameters.Add(`stackId`, *StackId)
	}

	return t.restClient.Post(
		`/v2/Stack/UseCase/CreateStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
func (t *Stack) DeleteStack(StackId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/DeleteStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PromotionId
// @param array StackIds
func (t *Stack) RemovePromotionFromStacks(PromotionId string, StackIds []string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`promotionId`, PromotionId)
	for i := range StackIds {
		queryParameters.Add(`stackIds[]`, StackIds[i])
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
func (t *Stack) SetClosingTimeForStack(StackId string, ClosingTime int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`closingTime`, strconv.Itoa(ClosingTime))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetClosingTimeForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string ConfirmDesignId
func (t *Stack) SetConfirmDesignForStack(StackId string, ConfirmDesignId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`confirmDesignId`, ConfirmDesignId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetConfirmDesignForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string DeclineDesignId
func (t *Stack) SetDeclineDesignForStack(StackId string, DeclineDesignId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`declineDesignId`, DeclineDesignId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetDeclineDesignForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string InviteDesignId
func (t *Stack) SetInviteDesignForStack(StackId string, InviteDesignId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`inviteDesignId`, InviteDesignId)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetInviteDesignForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param int MaxQuantity
func (t *Stack) SetMaxQuantityForStack(StackId string, MaxQuantity int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`maxQuantity`, strconv.Itoa(MaxQuantity))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetMaxQuantityForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param string MethodSlug
func (t *Stack) SetMethodForStack(StackId string, MethodSlug string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`methodSlug`, MethodSlug)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetMethodForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param int OpeningTime
func (t *Stack) SetOpeningTimeForStack(StackId string, OpeningTime int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`openingTime`, strconv.Itoa(OpeningTime))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetOpeningTimeForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param float Price
func (t *Stack) SetPriceForStack(StackId string, Price float) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`price`, Price)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetPriceForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param int Quantity
func (t *Stack) SetQuantityForStack(StackId string, Quantity int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`quantity`, strconv.Itoa(Quantity))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetQuantityForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param float ServiceFee
func (t *Stack) SetServiceFeeForStack(StackId string, ServiceFee float) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`serviceFee`, ServiceFee)

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetServiceFeeForStack`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string StackId
// @param bool Transferable true|false
func (t *Stack) SetTransferableForStack(StackId string, Transferable bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`stackId`, StackId)
	queryParameters.Add(`transferable`, strconv.FormatBool(Transferable))

	return t.restClient.Post(
		`/v2/Stack/UseCase/SetTransferableForStack`,
		&queryParameters,
		nil,
		nil,
	)
}
