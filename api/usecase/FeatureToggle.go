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

type FeatureToggle struct {
	restClient rest.RestClientInterface
}

func NewFeatureToggle(restClient rest.RestClientInterface) *FeatureToggle {
	return &FeatureToggle{restClient}
}

// GET: Queries

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
