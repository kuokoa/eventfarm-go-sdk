/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type PoolFeature struct {
	restClient rest.RestClientInterface
}

func NewPoolFeature(restClient rest.RestClientInterface) *PoolFeature {
	return &PoolFeature{restClient}
}

// GET: Queries

type ListFeaturesForPoolParameters struct {
	PoolId string
}

func (t *PoolFeature) ListFeaturesForPool(p *ListFeaturesForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/PoolFeature/UseCase/ListFeaturesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddFeatureForPoolParameters struct {
	PoolId    string
	FeatureId string
	Enabled   *bool
}

func (t *PoolFeature) AddFeatureForPool(p *AddFeatureForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`featureId`, p.FeatureId)
	if p.Enabled != nil {
		queryParameters.Add(`enabled`, strconv.FormatBool(*p.Enabled))
	}

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/AddFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolFeature) AddFeatureForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolFeature/UseCase/AddFeatureForPool`,
		data,
		nil,
		nil,
	)
}

type DisableFeatureForPoolParameters struct {
	PoolId    string
	FeatureId string
}

func (t *PoolFeature) DisableFeatureForPool(p *DisableFeatureForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`featureId`, p.FeatureId)

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/DisableFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolFeature) DisableFeatureForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolFeature/UseCase/DisableFeatureForPool`,
		data,
		nil,
		nil,
	)
}

type EnableFeatureForPoolParameters struct {
	PoolId    string
	FeatureId string
}

func (t *PoolFeature) EnableFeatureForPool(p *EnableFeatureForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`featureId`, p.FeatureId)

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/EnableFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolFeature) EnableFeatureForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolFeature/UseCase/EnableFeatureForPool`,
		data,
		nil,
		nil,
	)
}

type RemoveFeatureForPoolParameters struct {
	PoolId    string
	FeatureId string
}

func (t *PoolFeature) RemoveFeatureForPool(p *RemoveFeatureForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`featureId`, p.FeatureId)

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/RemoveFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *PoolFeature) RemoveFeatureForPoolWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/PoolFeature/UseCase/RemoveFeatureForPool`,
		data,
		nil,
		nil,
	)
}
