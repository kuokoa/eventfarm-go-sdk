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

type EFx struct {
	restClient rest.RestClientInterface
}

func NewEFx(restClient rest.RestClientInterface) *EFx {
	return &EFx{restClient}
}

// GET: Queries

type GetAdminConfigForEventParameters struct {
	EventId string
}

func (t *EFx) GetAdminConfigForEvent(p *GetAdminConfigForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/EFx/UseCase/GetAdminConfigForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetAllModulesForEventParameters struct {
	EventId             string
	WithDisabledModules *bool
}

func (t *EFx) GetAllModulesForEvent(p *GetAllModulesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithDisabledModules != nil {
		queryParameters.Add(`withDisabledModules`, strconv.FormatBool(*p.WithDisabledModules))
	}

	return t.restClient.Get(
		`/v2/EFx/UseCase/GetAllModulesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type DisableForEventParameters struct {
	EventId string
}

func (t *EFx) DisableForEvent(p *DisableForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type DisableModuleForEventParameters struct {
	EventId    string
	ModuleType string
}

func (t *EFx) DisableModuleForEvent(p *DisableModuleForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`moduleType`, p.ModuleType)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableModuleForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type DisableNFCForEventParameters struct {
	EventId string
}

func (t *EFx) DisableNFCForEvent(p *DisableNFCForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableNFCForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type DisableSMSForEventParameters struct {
	EventId string
}

func (t *EFx) DisableSMSForEvent(p *DisableSMSForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/DisableSMSForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type EnableForEventParameters struct {
	EventId string
}

func (t *EFx) EnableForEvent(p *EnableForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type EnableModuleForEventParameters struct {
	EventId    string
	ModuleType string
}

func (t *EFx) EnableModuleForEvent(p *EnableModuleForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`moduleType`, p.ModuleType)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableModuleForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type EnableNFCForEventParameters struct {
	EventId string
}

func (t *EFx) EnableNFCForEvent(p *EnableNFCForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableNFCForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type EnableSMSForEventParameters struct {
	EventId string
}

func (t *EFx) EnableSMSForEvent(p *EnableSMSForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/EFx/UseCase/EnableSMSForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type RequestForEventParameters struct {
	EventId             string
	UserId              string
	RequestedEFxModules *[]string // access-control | concierge | digital-memory-bank | messaging | smsquiz | product-pickup | raffle | reservation | roaming-photographer | smart-bar | teams
}

func (t *EFx) RequestForEvent(p *RequestForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`userId`, p.UserId)
	if p.RequestedEFxModules != nil {
		for i := range *p.RequestedEFxModules {
			queryParameters.Add(`requestedEFxModules[]`, (*p.RequestedEFxModules)[i])
		}
	}

	return t.restClient.Post(
		`/v2/EFx/UseCase/RequestForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type SendSMSWithAppLinkParameters struct {
	PhoneNumber string
}

func (t *EFx) SendSMSWithAppLink(p *SendSMSWithAppLinkParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`phoneNumber`, p.PhoneNumber)

	return t.restClient.Post(
		`/v2/EFx/UseCase/SendSMSWithAppLink`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetAdminPinForEventParameters struct {
	EventId string
	Pin     string
}

func (t *EFx) SetAdminPinForEvent(p *SetAdminPinForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`pin`, p.Pin)

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetAdminPinForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetSMSForEventParameters struct {
	EventId string
	Message string
}

func (t *EFx) SetSMSForEvent(p *SetSMSForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`message`, p.Message)

	return t.restClient.Post(
		`/v2/EFx/UseCase/SetSMSForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}
