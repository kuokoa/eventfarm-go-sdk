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

type SalesforcePoolSetting struct {
	restClient gosdk.RestClientInterface
}

func NewSalesforcePoolSetting(restClient gosdk.RestClientInterface) *SalesforcePoolSetting {
	return &SalesforcePoolSetting{restClient}
}

// GET: Queries
// @param string PoolId
func (t *SalesforcePoolSetting) GetSalesforcePoolSetting(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/SalesforcePoolSetting/UseCase/GetSalesforcePoolSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PoolId
// @param string|null SyncUserId
// @param string|null DefaultLeadCompany
// @param string|null DefaultLeadSource
func (t *SalesforcePoolSetting) CreateSalesforcePoolSetting(PoolId string, SyncUserId *string, DefaultLeadCompany *string, DefaultLeadSource *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	if SyncUserId != nil {
		queryParameters.Add(`syncUserId`, *SyncUserId)
	}
	if DefaultLeadCompany != nil {
		queryParameters.Add(`defaultLeadCompany`, *DefaultLeadCompany)
	}
	if DefaultLeadSource != nil {
		queryParameters.Add(`defaultLeadSource`, *DefaultLeadSource)
	}

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/CreateSalesforcePoolSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
func (t *SalesforcePoolSetting) RemoveSyncUserForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/RemoveSyncUserForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string LeadCompany
func (t *SalesforcePoolSetting) SetLeadCompanyForPool(PoolId string, LeadCompany string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`leadCompany`, LeadCompany)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/SetLeadCompanyForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string LeadSource
func (t *SalesforcePoolSetting) SetLeadSourceForPool(PoolId string, LeadSource string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`leadSource`, LeadSource)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/SetLeadSourceForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string SyncUserId
func (t *SalesforcePoolSetting) SetSyncUserForPool(PoolId string, SyncUserId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`syncUserId`, SyncUserId)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/SetSyncUserForPool`,
		&queryParameters,
		nil,
		nil,
	)
}
