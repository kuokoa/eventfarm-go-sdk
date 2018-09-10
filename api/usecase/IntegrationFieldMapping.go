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

type IntegrationFieldMapping struct {
	restClient gosdk.RestClientInterface
}

func NewIntegrationFieldMapping(restClient gosdk.RestClientInterface) *IntegrationFieldMapping {
	return &IntegrationFieldMapping{restClient}
}

// GET: Queries
// @param string IntegrationFieldMappingId
func (t *IntegrationFieldMapping) GetIntegrationFieldMapping(IntegrationFieldMappingId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, IntegrationFieldMappingId)

	return t.restClient.Get(
		`/v2/IntegrationFieldMapping/UseCase/GetIntegrationFieldMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SalesforceEventSettingId
// @param string|null FieldMappingType
func (t *IntegrationFieldMapping) ListFieldMappingsForSalesforceEventSetting(SalesforceEventSettingId string, FieldMappingType *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	if FieldMappingType != nil {
		queryParameters.Add(`fieldMappingType`, *FieldMappingType)
	}

	return t.restClient.Get(
		`/v2/IntegrationFieldMapping/UseCase/ListFieldMappingsForSalesforceEventSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string FieldMappingType salesforce-import-lead|salesforce-import-contact|salesforce-export-lead|salesforce-export-contact|marketo-field-mapping
// @param string IntegrationSettingType
// @param string IntegrationSettingId
// @param string FieldType question|user-attribute|user-name|user-identifier
// @param string FieldId
// @param string IntegrationFieldValue
// @param bool CanUpdateEventFarmField true|false
// @param bool CanUpdateIntegrationField true|false
// @param bool CanDeleteMapping true|false
// @param string|null UpdateRule never|different|blank
// @param string|null FieldName
func (t *IntegrationFieldMapping) CreateIntegrationFieldMapping(FieldMappingType string, IntegrationSettingType string, IntegrationSettingId string, FieldType string, FieldId string, IntegrationFieldValue string, CanUpdateEventFarmField bool, CanUpdateIntegrationField bool, CanDeleteMapping bool, UpdateRule *string, FieldName *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`fieldMappingType`, FieldMappingType)
	queryParameters.Add(`integrationSettingType`, IntegrationSettingType)
	queryParameters.Add(`integrationSettingId`, IntegrationSettingId)
	queryParameters.Add(`fieldType`, FieldType)
	queryParameters.Add(`fieldId`, FieldId)
	queryParameters.Add(`integrationFieldValue`, IntegrationFieldValue)
	queryParameters.Add(`canUpdateEventFarmField`, strconv.FormatBool(CanUpdateEventFarmField))
	queryParameters.Add(`canUpdateIntegrationField`, strconv.FormatBool(CanUpdateIntegrationField))
	queryParameters.Add(`canDeleteMapping`, strconv.FormatBool(CanDeleteMapping))
	if UpdateRule != nil {
		queryParameters.Add(`updateRule`, *UpdateRule)
	}
	if FieldName != nil {
		queryParameters.Add(`fieldName`, *FieldName)
	}

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/CreateIntegrationFieldMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationFieldMappingId
func (t *IntegrationFieldMapping) DeleteIntegrationFieldMapping(IntegrationFieldMappingId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, IntegrationFieldMappingId)

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/DeleteIntegrationFieldMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationFieldMappingId
// @param string FieldType question|user-attribute|user-name|user-identifier
// @param string FieldId
// @param string|null FieldName
func (t *IntegrationFieldMapping) SetIntegrationFieldMappingFields(IntegrationFieldMappingId string, FieldType string, FieldId string, FieldName *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, IntegrationFieldMappingId)
	queryParameters.Add(`fieldType`, FieldType)
	queryParameters.Add(`fieldId`, FieldId)
	if FieldName != nil {
		queryParameters.Add(`fieldName`, *FieldName)
	}

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/SetIntegrationFieldMappingFields`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationFieldMappingId
// @param string IntegrationFieldName
func (t *IntegrationFieldMapping) SetIntegrationFieldMappingName(IntegrationFieldMappingId string, IntegrationFieldName string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, IntegrationFieldMappingId)
	queryParameters.Add(`integrationFieldName`, IntegrationFieldName)

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/SetIntegrationFieldMappingName`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationFieldMappingId
// @param string UpdateRule never|different|blank
func (t *IntegrationFieldMapping) SetIntegrationFieldMappingUpdateRule(IntegrationFieldMappingId string, UpdateRule string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationFieldMappingId`, IntegrationFieldMappingId)
	queryParameters.Add(`updateRule`, UpdateRule)

	return t.restClient.Post(
		`/v2/IntegrationFieldMapping/UseCase/SetIntegrationFieldMappingUpdateRule`,
		&queryParameters,
		nil,
		nil,
	)
}
