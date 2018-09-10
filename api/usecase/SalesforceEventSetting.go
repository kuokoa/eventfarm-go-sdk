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

type SalesforceEventSetting struct {
	restClient gosdk.RestClientInterface
}

func NewSalesforceEventSetting(restClient gosdk.RestClientInterface) *SalesforceEventSetting {
	return &SalesforceEventSetting{restClient}
}

// GET: Queries
// @param string SalesforceEventSettingId
func (t *SalesforceEventSetting) GetSalesforceEventSetting(SalesforceEventSettingId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)

	return t.restClient.Get(
		`/v2/SalesforceEventSetting/UseCase/GetSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *SalesforceEventSetting) ListSalesforceEventSettingsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/SalesforceEventSetting/UseCase/ListSalesforceEventSettingsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string SalesforceEventSettingId
func (t *SalesforceEventSetting) DisableExportForSalesforceEventSetting(SalesforceEventSettingId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/DisableExportForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
func (t *SalesforceEventSetting) DisableImportForSalesforceEventSetting(SalesforceEventSettingId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/DisableImportForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
func (t *SalesforceEventSetting) EnableExportForSalesforceEventSetting(SalesforceEventSettingId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/EnableExportForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
func (t *SalesforceEventSetting) EnableImportForSalesforceEventSetting(SalesforceEventSettingId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/EnableImportForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string CampaignId
// @param string CampaignName
func (t *SalesforceEventSetting) SetCampaignForSalesforceEventSetting(SalesforceEventSettingId string, CampaignId string, CampaignName string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	queryParameters.Add(`campaignId`, CampaignId)
	queryParameters.Add(`campaignName`, CampaignName)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetCampaignForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param int InvitationCount 1-100
func (t *SalesforceEventSetting) SetInvitationCountForSalesforceEventSetting(SalesforceEventSettingId string, InvitationCount int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	queryParameters.Add(`invitationCount`, strconv.Itoa(InvitationCount))

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetInvitationCountForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string InvitationCreationType unconfirmed-no-email|confirmed-no-email
func (t *SalesforceEventSetting) SetInvitationCreationTypeForSalesforceEventSetting(SalesforceEventSettingId string, InvitationCreationType string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	queryParameters.Add(`invitationCreationType`, InvitationCreationType)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetInvitationCreationTypeForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string NewContactRule do-nothing|create-contact|create-lead
func (t *SalesforceEventSetting) SetNewContactRuleForSalesforceEventSetting(SalesforceEventSettingId string, NewContactRule string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	queryParameters.Add(`newContactRule`, NewContactRule)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetNewContactRuleForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string StackId
func (t *SalesforceEventSetting) SetStackForSalesforceEventSetting(SalesforceEventSettingId string, StackId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	queryParameters.Add(`stackId`, StackId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetStackForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}
