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

type PoolFeature struct {
	restClient gosdk.RestClientInterface
}

func NewPoolFeature(restClient gosdk.RestClientInterface) *PoolFeature {
	return &PoolFeature{restClient}
}

// GET: Queries
// @param string PoolId
func (t *PoolFeature) ListFeaturesForPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/PoolFeature/UseCase/ListFeaturesForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PoolId
// @param string FeatureId
// @param bool|null Enabled true|false
func (t *PoolFeature) AddFeatureForPool(PoolId string, FeatureId string, Enabled *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`featureId`, FeatureId)
	if Enabled != nil {
		queryParameters.Add(`enabled`, strconv.FormatBool(*Enabled))
	}

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/AddFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string FeatureId
func (t *PoolFeature) DisableFeatureForPool(PoolId string, FeatureId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`featureId`, FeatureId)

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/DisableFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string FeatureId
func (t *PoolFeature) EnableFeatureForPool(PoolId string, FeatureId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`featureId`, FeatureId)

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/EnableFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string PoolId
// @param string FeatureId
func (t *PoolFeature) RemoveFeatureForPool(PoolId string, FeatureId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`featureId`, FeatureId)

	return t.restClient.Post(
		`/v2/PoolFeature/UseCase/RemoveFeatureForPool`,
		&queryParameters,
		nil,
		nil,
	)
}
