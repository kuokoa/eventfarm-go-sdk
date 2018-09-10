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

type TicketType struct {
	restClient gosdk.RestClientInterface
}

func NewTicketType(restClient gosdk.RestClientInterface) *TicketType {
	return &TicketType{restClient}
}

// GET: Queries
// @param string EventId
// @param bool|null ShouldHideDeleted true|false
func (t *TicketType) ListTicketTypesForEvent(EventId string, ShouldHideDeleted *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*ShouldHideDeleted))
	}

	return t.restClient.Get(
		`/v2/TicketType/UseCase/ListTicketTypesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EventId
// @param string Name
// @param string Code
// @param int Quantity
// @param int SortOrder
// @param bool|null IsDeleted true|false
// @param string|null Description
// @param string|null TicketTypeId
func (t *TicketType) CreateTicketType(EventId string, Name string, Code string, Quantity int, SortOrder int, IsDeleted *bool, Description *string, TicketTypeId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`name`, Name)
	queryParameters.Add(`code`, Code)
	queryParameters.Add(`quantity`, strconv.Itoa(Quantity))
	queryParameters.Add(`sortOrder`, strconv.Itoa(SortOrder))
	if IsDeleted != nil {
		queryParameters.Add(`isDeleted`, strconv.FormatBool(*IsDeleted))
	}
	if Description != nil {
		queryParameters.Add(`description`, *Description)
	}
	if TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *TicketTypeId)
	}

	return t.restClient.Post(
		`/v2/TicketType/UseCase/CreateTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
func (t *TicketType) DeleteTicketType(TicketTypeId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, TicketTypeId)

	return t.restClient.Post(
		`/v2/TicketType/UseCase/DeleteTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param string Description
func (t *TicketType) SetDescriptionForTicketType(TicketTypeId string, Description string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, TicketTypeId)
	queryParameters.Add(`description`, Description)

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetDescriptionForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param int DisplayOrder
func (t *TicketType) SetDisplayOrderForTicketType(TicketTypeId string, DisplayOrder int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, TicketTypeId)
	queryParameters.Add(`displayOrder`, strconv.Itoa(DisplayOrder))

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetDisplayOrderForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param string Name
func (t *TicketType) SetNameForTicketType(TicketTypeId string, Name string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, TicketTypeId)
	queryParameters.Add(`name`, Name)

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetNameForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param int Quantity
func (t *TicketType) SetQuantityForTicketType(TicketTypeId string, Quantity int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, TicketTypeId)
	queryParameters.Add(`quantity`, strconv.Itoa(Quantity))

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetQuantityForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}
