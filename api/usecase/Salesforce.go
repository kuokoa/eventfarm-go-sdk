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

type Salesforce struct {
	restClient gosdk.RestClientInterface
}

func NewSalesforce(restClient gosdk.RestClientInterface) *Salesforce {
	return &Salesforce{restClient}
}

// GET: Queries
// @param string PoolId
func (t *Salesforce) GetSalesforceLimitsForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/GetSalesforceLimitsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Salesforce) GetSalesforceStatusForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/GetSalesforceStatusForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *Salesforce) GetSalesforceSyncUserInfoForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/GetSalesforceSyncUserInfoForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *Salesforce) ListSalesforceCampaignsForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceCampaignsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *Salesforce) ListSalesforceContactFieldsForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceContactFieldsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *Salesforce) ListSalesforceLeadFieldsForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceLeadFieldsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string CampaignId
func (t *Salesforce) ListSalesforceStatusNamesForEvent(EventId string, CampaignId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`campaignId`, CampaignId)

	return t.restClient.Get(
		`/v2/Salesforce/UseCase/ListSalesforceStatusNamesForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EventId
func (t *Salesforce) ExportEventToSalesforce(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Salesforce/UseCase/ExportEventToSalesforce`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string InvitationId
func (t *Salesforce) ExportInvitationToSalesforce(InvitationId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)

	return t.restClient.Post(
		`/v2/Salesforce/UseCase/ExportInvitationToSalesforce`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string|null UserId
func (t *Salesforce) ImportCampaignMembersForEvent(EventId string, UserId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if UserId != nil {
		queryParameters.Add(`userId`, *UserId)
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
func (t *Salesforce) ImportCampaignMembersForSalesforceEventSetting(SalesforceEventSettingId string, UserId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	if UserId != nil {
		queryParameters.Add(`userId`, *UserId)
	}

	return t.restClient.Post(
		`/v2/Salesforce/UseCase/ImportCampaignMembersForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}
