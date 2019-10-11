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

type SalesforceEventSetting struct {
	restClient rest.RestClientInterface
}

func NewSalesforceEventSetting(restClient rest.RestClientInterface) *SalesforceEventSetting {
	return &SalesforceEventSetting{restClient}
}

// GET: Queries

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

type SetInvitationCountForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	InvitationCount          int64 // 1-100
}

func (t *SalesforceEventSetting) SetInvitationCountForSalesforceEventSetting(p *SetInvitationCountForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	queryParameters.Add(`invitationCount`, strconv.FormatInt(p.InvitationCount, 10))

	return t.restClient.Post(
		`/v2/SalesforceEventSetting/UseCase/SetInvitationCountForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

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
