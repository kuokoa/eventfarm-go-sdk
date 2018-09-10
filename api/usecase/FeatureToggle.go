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

type FeatureToggle struct {
	restClient gosdk.RestClientInterface
}

func NewFeatureToggle(restClient gosdk.RestClientInterface) *FeatureToggle {
	return &FeatureToggle{restClient}
}

// GET: Queries
// @param string FeatureName
// @param string|null UserId
func (t *FeatureToggle) GetFeatureGrant(FeatureName string, UserId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`featureName`, FeatureName)
	if UserId != nil {
		queryParameters.Add(`userId`, *UserId)
	}

	return t.restClient.Get(
		`/v2/FeatureToggle/UseCase/GetFeatureGrant`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
