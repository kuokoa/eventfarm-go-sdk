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
