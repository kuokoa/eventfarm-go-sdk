/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk"
)

type Salesforce struct {
	restClient sdk.RestClientInterface
}

func NewSalesforce(restClient sdk.RestClientInterface) *Salesforce {
	return &Salesforce{restClient}
}

// GET: Queries
// @param string PoolId

type GetSalesforceLimitsForPoolParameters struct {
	PoolId string
}

func (t *Salesforce) GetSalesforceLimitsForPool(p *GetSalesforceLimitsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/GetSalesforceLimitsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId

type GetSalesforceStatusForEventParameters struct {
	EventId string
}

func (t *Salesforce) GetSalesforceStatusForEvent(p *GetSalesforceStatusForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/GetSalesforceStatusForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type GetSalesforceSyncUserInfoForPoolParameters struct {
	PoolId string
}

func (t *Salesforce) GetSalesforceSyncUserInfoForPool(p *GetSalesforceSyncUserInfoForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/GetSalesforceSyncUserInfoForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type ListSalesforceCampaignsForPoolParameters struct {
	PoolId string
}

func (t *Salesforce) ListSalesforceCampaignsForPool(p *ListSalesforceCampaignsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceCampaignsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type ListSalesforceContactFieldsForPoolParameters struct {
	PoolId string
}

func (t *Salesforce) ListSalesforceContactFieldsForPool(p *ListSalesforceContactFieldsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceContactFieldsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type ListSalesforceLeadFieldsForPoolParameters struct {
	PoolId string
}

func (t *Salesforce) ListSalesforceLeadFieldsForPool(p *ListSalesforceLeadFieldsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceLeadFieldsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string CampaignId

type ListSalesforceStatusNamesForEventParameters struct {
	EventId    string
	CampaignId string
}

func (t *Salesforce) ListSalesforceStatusNamesForEvent(p *ListSalesforceStatusNamesForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`campaignId`, p.CampaignId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceStatusNamesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EventId

type ExportEventToSalesforceParameters struct {
	EventId string
}

func (t *Salesforce) ExportEventToSalesforce(p *ExportEventToSalesforceParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Salesforce/UseCase/ExportEventToSalesforce`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId

type ExportInvitationToSalesforceParameters struct {
	InvitationId string
}

func (t *Salesforce) ExportInvitationToSalesforce(p *ExportInvitationToSalesforceParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)

	return t.restClient.Post(
		`/v2/Salesforce/UseCase/ExportInvitationToSalesforce`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string|null UserId

type ImportCampaignMembersForEventParameters struct {
	EventId string
	UserId  *string
}

func (t *Salesforce) ImportCampaignMembersForEvent(p *ImportCampaignMembersForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.UserId != nil {
		queryParameters.Add(`userId`, *p.UserId)
	}

	return t.restClient.Post(
		`/v2/Salesforce/UseCase/ImportCampaignMembersForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string|null UserId

type ImportCampaignMembersForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	UserId                   *string
}

func (t *Salesforce) ImportCampaignMembersForSalesforceEventSetting(p *ImportCampaignMembersForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	if p.UserId != nil {
		queryParameters.Add(`userId`, *p.UserId)
	}

	return t.restClient.Post(
		`/v2/Salesforce/UseCase/ImportCampaignMembersForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}
