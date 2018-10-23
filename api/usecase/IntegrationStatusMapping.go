/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type IntegrationStatusMapping struct {
	restClient rest.RestClientInterface
}

func NewIntegrationStatusMapping(restClient rest.RestClientInterface) *IntegrationStatusMapping {
	return &IntegrationStatusMapping{restClient}
}

// GET: Queries
// @param string SalesforceEventSettingId
// @param string|null StatusMappingType salesforce-campaign-member

type ListStatusMappingsForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	StatusMappingType        *string
}

func (t *IntegrationStatusMapping) ListStatusMappingsForSalesforceEventSetting(p *ListStatusMappingsForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	if p.StatusMappingType != nil {
		queryParameters.Add(`statusMappingType`, *p.StatusMappingType)
	}

	return t.restClient.Get(
		`/v2/IntegrationStatusMapping/UseCase/ListStatusMappingsForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string StatusMappingType salesforce-campaign-member
// @param string IntegrationSettingType salesforce|marketo
// @param string IntegrationSettingId
// @param string StatusId assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted|checked-in
// @param string IntegrationStatusName

type CreateIntegrationStatusMappingParameters struct {
	StatusMappingType      string
	IntegrationSettingType string
	IntegrationSettingId   string
	StatusId               string
	IntegrationStatusName  string
}

func (t *IntegrationStatusMapping) CreateIntegrationStatusMapping(p *CreateIntegrationStatusMappingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`statusMappingType`, p.StatusMappingType)
	queryParameters.Add(`integrationSettingType`, p.IntegrationSettingType)
	queryParameters.Add(`integrationSettingId`, p.IntegrationSettingId)
	queryParameters.Add(`statusId`, p.StatusId)
	queryParameters.Add(`integrationStatusName`, p.IntegrationStatusName)

	return t.restClient.Post(
		`/v2/IntegrationStatusMapping/UseCase/CreateIntegrationStatusMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationStatusMappingId
// @param string StatusId assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted|checked-in

type SetIntegrationStatusMappingStatusIdParameters struct {
	IntegrationStatusMappingId string
	StatusId                   string
}

func (t *IntegrationStatusMapping) SetIntegrationStatusMappingStatusId(p *SetIntegrationStatusMappingStatusIdParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationStatusMappingId`, p.IntegrationStatusMappingId)
	queryParameters.Add(`statusId`, p.StatusId)

	return t.restClient.Post(
		`/v2/IntegrationStatusMapping/UseCase/SetIntegrationStatusMappingStatusId`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationStatusMappingId
// @param string IntegrationStatusValue

type SetIntegrationStatusMappingValueParameters struct {
	IntegrationStatusMappingId string
	IntegrationStatusValue     string
}

func (t *IntegrationStatusMapping) SetIntegrationStatusMappingValue(p *SetIntegrationStatusMappingValueParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationStatusMappingId`, p.IntegrationStatusMappingId)
	queryParameters.Add(`integrationStatusValue`, p.IntegrationStatusValue)

	return t.restClient.Post(
		`/v2/IntegrationStatusMapping/UseCase/SetIntegrationStatusMappingValue`,
		&queryParameters,
		nil,
		nil,
	)
}
