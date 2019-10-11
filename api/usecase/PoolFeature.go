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
