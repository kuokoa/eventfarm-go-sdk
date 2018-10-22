/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk"
)

type TicketType struct {
	restClient sdk.RestClientInterface
}

func NewTicketType(restClient sdk.RestClientInterface) *TicketType {
	return &TicketType{restClient}
}

// GET: Queries
// @param string EventId
// @param bool|null ShouldHideDeleted true|false

type ListTicketTypesForEventParameters struct {
	EventId           string
	ShouldHideDeleted *bool
}

func (t *TicketType) ListTicketTypesForEvent(p *ListTicketTypesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.ShouldHideDeleted != nil {
		queryParameters.Add(`shouldHideDeleted`, strconv.FormatBool(*p.ShouldHideDeleted))
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

type CreateTicketTypeParameters struct {
	EventId      string
	Name         string
	Code         string
	Quantity     int
	SortOrder    int
	IsDeleted    *bool
	Description  *string
	TicketTypeId *string
}

func (t *TicketType) CreateTicketType(p *CreateTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`code`, p.Code)
	queryParameters.Add(`quantity`, strconv.Itoa(p.Quantity))
	queryParameters.Add(`sortOrder`, strconv.Itoa(p.SortOrder))
	if p.IsDeleted != nil {
		queryParameters.Add(`isDeleted`, strconv.FormatBool(*p.IsDeleted))
	}
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}
	if p.TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *p.TicketTypeId)
	}

	return t.restClient.Post(
		`/v2/TicketType/UseCase/CreateTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId

type DeleteTicketTypeParameters struct {
	TicketTypeId string
}

func (t *TicketType) DeleteTicketType(p *DeleteTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)

	return t.restClient.Post(
		`/v2/TicketType/UseCase/DeleteTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param string Description

type SetDescriptionForTicketTypeParameters struct {
	TicketTypeId string
	Description  string
}

func (t *TicketType) SetDescriptionForTicketType(p *SetDescriptionForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`description`, p.Description)

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetDescriptionForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param int DisplayOrder

type SetDisplayOrderForTicketTypeParameters struct {
	TicketTypeId string
	DisplayOrder int
}

func (t *TicketType) SetDisplayOrderForTicketType(p *SetDisplayOrderForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`displayOrder`, strconv.Itoa(p.DisplayOrder))

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetDisplayOrderForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param string Name

type SetNameForTicketTypeParameters struct {
	TicketTypeId string
	Name         string
}

func (t *TicketType) SetNameForTicketType(p *SetNameForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetNameForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TicketTypeId
// @param int Quantity

type SetQuantityForTicketTypeParameters struct {
	TicketTypeId string
	Quantity     int
}

func (t *TicketType) SetQuantityForTicketType(p *SetQuantityForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`quantity`, strconv.Itoa(p.Quantity))

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetQuantityForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}
