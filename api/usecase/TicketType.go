/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type TicketType struct {
	restClient rest.RestClientInterface
}

func NewTicketType(restClient rest.RestClientInterface) *TicketType {
	return &TicketType{restClient}
}

// GET: Queries

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

type CreateTicketTypeParameters struct {
	EventId      string
	Name         string
	Code         string
	Quantity     int64
	SortOrder    int64
	IsDeleted    *bool
	Description  *string
	TicketTypeId *string
}

func (t *TicketType) CreateTicketType(p *CreateTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`code`, p.Code)
	queryParameters.Add(`quantity`, strconv.FormatInt(p.Quantity, 10))
	queryParameters.Add(`sortOrder`, strconv.FormatInt(p.SortOrder, 10))
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

func (t *TicketType) CreateTicketTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/TicketType/UseCase/CreateTicketType`,
		data,
		nil,
		nil,
	)
}

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

func (t *TicketType) DeleteTicketTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/TicketType/UseCase/DeleteTicketType`,
		data,
		nil,
		nil,
	)
}

type SetDescriptionForTicketTypeParameters struct {
	TicketTypeId string
	Description  *string
}

func (t *TicketType) SetDescriptionForTicketType(p *SetDescriptionForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	if p.Description != nil {
		queryParameters.Add(`description`, *p.Description)
	}

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetDescriptionForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *TicketType) SetDescriptionForTicketTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/TicketType/UseCase/SetDescriptionForTicketType`,
		data,
		nil,
		nil,
	)
}

type SetDisplayOrderForTicketTypeParameters struct {
	TicketTypeId string
	DisplayOrder int64
}

func (t *TicketType) SetDisplayOrderForTicketType(p *SetDisplayOrderForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`displayOrder`, strconv.FormatInt(p.DisplayOrder, 10))

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetDisplayOrderForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *TicketType) SetDisplayOrderForTicketTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/TicketType/UseCase/SetDisplayOrderForTicketType`,
		data,
		nil,
		nil,
	)
}

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

func (t *TicketType) SetNameForTicketTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/TicketType/UseCase/SetNameForTicketType`,
		data,
		nil,
		nil,
	)
}

type SetQuantityForTicketTypeParameters struct {
	TicketTypeId string
	Quantity     int64
}

func (t *TicketType) SetQuantityForTicketType(p *SetQuantityForTicketTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`ticketTypeId`, p.TicketTypeId)
	queryParameters.Add(`quantity`, strconv.FormatInt(p.Quantity, 10))

	return t.restClient.Post(
		`/v2/TicketType/UseCase/SetQuantityForTicketType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *TicketType) SetQuantityForTicketTypeWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/TicketType/UseCase/SetQuantityForTicketType`,
		data,
		nil,
		nil,
	)
}
