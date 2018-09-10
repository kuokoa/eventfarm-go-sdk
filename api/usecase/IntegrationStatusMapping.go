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

type IntegrationStatusMapping struct {
	restClient gosdk.RestClientInterface
}

func NewIntegrationStatusMapping(restClient gosdk.RestClientInterface) *IntegrationStatusMapping {
	return &IntegrationStatusMapping{restClient}
}

// GET: Queries
// @param string SalesforceEventSettingId
// @param string|null StatusMappingType salesforce-campaign-member
func (t *IntegrationStatusMapping) ListStatusMappingsForSalesforceEventSetting(SalesforceEventSettingId string, StatusMappingType *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`salesforceEventSettingId`, SalesforceEventSettingId)
	if StatusMappingType != nil {
		queryParameters.Add(`statusMappingType`, *StatusMappingType)
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
func (t *IntegrationStatusMapping) CreateIntegrationStatusMapping(StatusMappingType string, IntegrationSettingType string, IntegrationSettingId string, StatusId string, IntegrationStatusName string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`statusMappingType`, StatusMappingType)
	queryParameters.Add(`integrationSettingType`, IntegrationSettingType)
	queryParameters.Add(`integrationSettingId`, IntegrationSettingId)
	queryParameters.Add(`statusId`, StatusId)
	queryParameters.Add(`integrationStatusName`, IntegrationStatusName)

	return t.restClient.Post(
		`/v2/IntegrationStatusMapping/UseCase/CreateIntegrationStatusMapping`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationStatusMappingId
// @param string StatusId assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted|checked-in
func (t *IntegrationStatusMapping) SetIntegrationStatusMappingStatusId(IntegrationStatusMappingId string, StatusId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationStatusMappingId`, IntegrationStatusMappingId)
	queryParameters.Add(`statusId`, StatusId)

	return t.restClient.Post(
		`/v2/IntegrationStatusMapping/UseCase/SetIntegrationStatusMappingStatusId`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string IntegrationStatusMappingId
// @param string IntegrationStatusValue
func (t *IntegrationStatusMapping) SetIntegrationStatusMappingValue(IntegrationStatusMappingId string, IntegrationStatusValue string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`integrationStatusMappingId`, IntegrationStatusMappingId)
	queryParameters.Add(`integrationStatusValue`, IntegrationStatusValue)

	return t.restClient.Post(
		`/v2/IntegrationStatusMapping/UseCase/SetIntegrationStatusMappingValue`,
		&queryParameters,
		nil,
		nil,
	)
}
