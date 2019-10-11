/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"fmt"
	"github.com/eventfarm/go-sdk/rest"
	"net/http"
	"net/url"
	"strconv"
)

type IntegrationFieldMapping struct {
	restClient rest.RestClientInterface
}

func NewIntegrationFieldMapping(restClient rest.RestClientInterface) *IntegrationFieldMapping {
	return &IntegrationFieldMapping{restClient}
}

// GET: Queries

type GetIntegrationFieldMappingParameters struct {
	IntegrationFieldMappingId string
}

func (t *IntegrationFieldMapping) GetIntegrationFieldMapping(p *GetIntegrationFieldMappingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, p.IntegrationFieldMappingId)

	return t.restClient.Get(
		`/v2/IntegrationFieldMapping/UseCase/GetIntegrationFieldMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListFieldMappingsForSalesforceEventSettingParameters struct {
	SalesforceEventSettingId string
	FieldMappingType         *string
}

func (t *IntegrationFieldMapping) ListFieldMappingsForSalesforceEventSetting(p *ListFieldMappingsForSalesforceEventSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, p.SalesforceEventSettingId)
	if p.FieldMappingType != nil {
		queryParameters.Add(`fieldMappingType`, *p.FieldMappingType)
	}

	return t.restClient.Get(
		`/v2/IntegrationFieldMapping/UseCase/ListFieldMappingsForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type CreateIntegrationFieldMappingParameters struct {
	FieldMappingType          string
	IntegrationSettingType    string
	IntegrationSettingId      string
	FieldType                 string
	FieldId                   string
	IntegrationFieldValue     string
	CanUpdateEventFarmField   bool
	CanUpdateIntegrationField bool
	CanDeleteMapping          bool
	UpdateRule                *string
	FieldName                 *string
}

func (t *IntegrationFieldMapping) CreateIntegrationFieldMapping(p *CreateIntegrationFieldMappingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`fieldMappingType`, p.FieldMappingType)
	queryParameters.Add(`integrationSettingType`, p.IntegrationSettingType)
	queryParameters.Add(`integrationSettingId`, p.IntegrationSettingId)
	queryParameters.Add(`fieldType`, p.FieldType)
	queryParameters.Add(`fieldId`, p.FieldId)
	queryParameters.Add(`integrationFieldValue`, p.IntegrationFieldValue)
	queryParameters.Add(`canUpdateEventFarmField`, strconv.FormatBool(p.CanUpdateEventFarmField))
	queryParameters.Add(`canUpdateIntegrationField`, strconv.FormatBool(p.CanUpdateIntegrationField))
	queryParameters.Add(`canDeleteMapping`, strconv.FormatBool(p.CanDeleteMapping))
	if p.UpdateRule != nil {
		queryParameters.Add(`updateRule`, *p.UpdateRule)
	}
	if p.FieldName != nil {
		queryParameters.Add(`fieldName`, *p.FieldName)
	}

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/CreateIntegrationFieldMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

type DeleteIntegrationFieldMappingParameters struct {
	IntegrationFieldMappingId string
}

func (t *IntegrationFieldMapping) DeleteIntegrationFieldMapping(p *DeleteIntegrationFieldMappingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, p.IntegrationFieldMappingId)

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/DeleteIntegrationFieldMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetIntegrationFieldMappingFieldsParameters struct {
	IntegrationFieldMappingId string
	FieldType                 string
	FieldId                   string
	FieldName                 *string
}

func (t *IntegrationFieldMapping) SetIntegrationFieldMappingFields(p *SetIntegrationFieldMappingFieldsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, p.IntegrationFieldMappingId)
	queryParameters.Add(`fieldType`, p.FieldType)
	queryParameters.Add(`fieldId`, p.FieldId)
	if p.FieldName != nil {
		queryParameters.Add(`fieldName`, *p.FieldName)
	}

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/SetIntegrationFieldMappingFields`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetIntegrationFieldMappingNameParameters struct {
	IntegrationFieldMappingId string
	IntegrationFieldName      string
}

func (t *IntegrationFieldMapping) SetIntegrationFieldMappingName(p *SetIntegrationFieldMappingNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, p.IntegrationFieldMappingId)
	queryParameters.Add(`integrationFieldName`, p.IntegrationFieldName)

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/SetIntegrationFieldMappingName`,
		&queryParameters,
		nil,
		nil,
	)
}

type SetIntegrationFieldMappingUpdateRuleParameters struct {
	IntegrationFieldMappingId string
	UpdateRule                string
}

func (t *IntegrationFieldMapping) SetIntegrationFieldMappingUpdateRule(p *SetIntegrationFieldMappingUpdateRuleParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, p.IntegrationFieldMappingId)
	queryParameters.Add(`updateRule`, p.UpdateRule)

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/SetIntegrationFieldMappingUpdateRule`,
		&queryParameters,
		nil,
		nil,
	)
}
