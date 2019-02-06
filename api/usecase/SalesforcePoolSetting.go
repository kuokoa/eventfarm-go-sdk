/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type SalesforcePoolSetting struct {
	restClient rest.RestClientInterface
}

func NewSalesforcePoolSetting(restClient rest.RestClientInterface) *SalesforcePoolSetting {
	return &SalesforcePoolSetting{restClient}
}

// GET: Queries

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
