package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type SalesforcePoolSetting struct {
	restClient sdk.RestClientInterface
}

func NewSalesforcePoolSetting(restClient sdk.RestClientInterface) *SalesforcePoolSetting {
	return &SalesforcePoolSetting{restClient}
}

// GET: Queries
// @param string PoolId

type GetSalesforcePoolSettingParameters struct {
	PoolId string
}

func (t *SalesforcePoolSetting) GetSalesforcePoolSetting(p *GetSalesforcePoolSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

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

type CreateSalesforcePoolSettingParameters struct {
	PoolId             string
	SyncUserId         *string
	DefaultLeadCompany *string
	DefaultLeadSource  *string
}

func (t *SalesforcePoolSetting) CreateSalesforcePoolSetting(p *CreateSalesforcePoolSettingParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.SyncUserId != nil {
		queryParameters.Add(`syncUserId`, *p.SyncUserId)
	}
	if p.DefaultLeadCompany != nil {
		queryParameters.Add(`defaultLeadCompany`, *p.DefaultLeadCompany)
	}
	if p.DefaultLeadSource != nil {
		queryParameters.Add(`defaultLeadSource`, *p.DefaultLeadSource)
	}

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/CreateSalesforcePoolSetting`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId

type RemoveSyncUserForPoolParameters struct {
	PoolId string
}

func (t *SalesforcePoolSetting) RemoveSyncUserForPool(p *RemoveSyncUserForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/RemoveSyncUserForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string LeadCompany

type SetLeadCompanyForPoolParameters struct {
	PoolId      string
	LeadCompany string
}

func (t *SalesforcePoolSetting) SetLeadCompanyForPool(p *SetLeadCompanyForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`leadCompany`, p.LeadCompany)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/SetLeadCompanyForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string LeadSource

type SetLeadSourceForPoolParameters struct {
	PoolId     string
	LeadSource string
}

func (t *SalesforcePoolSetting) SetLeadSourceForPool(p *SetLeadSourceForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`leadSource`, p.LeadSource)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/SetLeadSourceForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string SyncUserId

type SetSyncUserForPoolParameters struct {
	PoolId     string
	SyncUserId string
}

func (t *SalesforcePoolSetting) SetSyncUserForPool(p *SetSyncUserForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`syncUserId`, p.SyncUserId)

	return t.restClient.Post(
		`/v2/SalesforcePoolSetting/UseCase/SetSyncUserForPool`,
		&queryParameters,
		nil,
		nil,
	)
}
