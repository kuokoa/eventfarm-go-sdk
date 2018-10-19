package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = fmt.Sprint("Foo")
var _ = http.NoBody

type SalesforceEventSetting struct {
	restClient sdk.RestClientInterface
}

func NewSalesforceEventSetting(restClient sdk.RestClientInterface) *SalesforceEventSetting {
	return &SalesforceEventSetting{restClient}
}

// GET: Queries
// @param string SalesforceEventSettingId

type GetSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
}

func (t *SalesforceEventSetting) GetSalesforceEventSetting(p *GetSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)

	return t.restClient.Get(
		`/v2/SalesforceEventSetting/UseCase/GetSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId

type ListSalesforceEventSettingsForEventParameters struct {
	EventId string
}

func (t *SalesforceEventSetting) ListSalesforceEventSettingsForEvent(p *ListSalesforceEventSettingsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Get(
		`/v2/SalesforceEventSetting/UseCase/ListSalesforceEventSettingsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string SalesforceEventSettingId

type DisableExportForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
}

func (t *SalesforceEventSetting) DisableExportForSalesforceEventSetting(p *DisableExportForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/DisableExportForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId

type DisableImportForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
}

func (t *SalesforceEventSetting) DisableImportForSalesforceEventSetting(p *DisableImportForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/DisableImportForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId

type EnableExportForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
}

func (t *SalesforceEventSetting) EnableExportForSalesforceEventSetting(p *EnableExportForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/EnableExportForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId

type EnableImportForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
}

func (t *SalesforceEventSetting) EnableImportForSalesforceEventSetting(p *EnableImportForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)

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

type SetCampaignForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	CampaignId               string
	CampaignName             string
}

func (t *SalesforceEventSetting) SetCampaignForSalesforceEventSetting(p *SetCampaignForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	queryParameters.Add(`campaignId`, p.CampaignId)
	queryParameters.Add(`campaignName`, p.CampaignName)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetCampaignForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param int InvitationCount 1-100

type SetInvitationCountForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	InvitationCount          int
}

func (t *SalesforceEventSetting) SetInvitationCountForSalesforceEventSetting(p *SetInvitationCountForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	queryParameters.Add(`invitationCount`, strconv.Itoa(p.InvitationCount))

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetInvitationCountForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string InvitationCreationType unconfirmed-no-email|confirmed-no-email

type SetInvitationCreationTypeForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	InvitationCreationType   string
}

func (t *SalesforceEventSetting) SetInvitationCreationTypeForSalesforceEventSetting(p *SetInvitationCreationTypeForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	queryParameters.Add(`invitationCreationType`, p.InvitationCreationType)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetInvitationCreationTypeForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string NewContactRule do-nothing|create-contact|create-lead

type SetNewContactRuleForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	NewContactRule           string
}

func (t *SalesforceEventSetting) SetNewContactRuleForSalesforceEventSetting(p *SetNewContactRuleForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	queryParameters.Add(`newContactRule`, p.NewContactRule)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetNewContactRuleForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string StackId

type SetStackForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	StackId                  string
}

func (t *SalesforceEventSetting) SetStackForSalesforceEventSetting(p *SetStackForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	queryParameters.Add(`stackId`, p.StackId)

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetStackForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}
