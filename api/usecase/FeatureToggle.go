package usecase

import (
	"net/http"
	"net/url"

	"bitbucket.ef.network/go/sdk"
)

type FeatureToggle struct {
	restClient sdk.RestClientInterface
}

func NewFeatureToggle(restClient sdk.RestClientInterface) *FeatureToggle {
	return &FeatureToggle{restClient}
}

// GET: Queries
// @param string FeatureName
// @param string|null UserId

type GetFeatureGrantParameters struct {
	FeatureName string
	UserId      *string
}

func (t *FeatureToggle) GetFeatureGrant(p *GetFeatureGrantParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`featureName`, p.FeatureName)
	if p.UserId != nil {
		queryParameters.Add(`userId`, *p.UserId)
	}

	return t.restClient.Get(
		`/v2/FeatureToggle/UseCase/GetFeatureGrant`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
