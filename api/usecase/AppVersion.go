/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type AppVersion struct {
	restClient rest.RestClientInterface
}

func NewAppVersion(restClient rest.RestClientInterface) *AppVersion {
	return &AppVersion{restClient}
}

// GET: Queries
// @param string AppVersionId

type GetAppVersionParameters struct {
	AppVersionId string
}

func (t *AppVersion) GetAppVersion(p *GetAppVersionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`appVersionId`, p.AppVersionId)

	return t.restClient.Get(
		`/v2/AppVersion/UseCase/GetAppVersion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string AppVersionType check-in-ios|check-in-android|ticket-block-mgmt-ios|ticket-block-mgmt-android

type GetAppVersionByTypeParameters struct {
	AppVersionType string
}

func (t *AppVersion) GetAppVersionByType(p *GetAppVersionByTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`appVersionType`, p.AppVersionType)

	return t.restClient.Get(
		`/v2/AppVersion/UseCase/GetAppVersionByType`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *AppVersion) GetSystemStatus() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/AppVersion/UseCase/GetSystemStatus`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *AppVersion) ListAppVersions() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/AppVersion/UseCase/ListAppVersions`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string AppVersionType check-in-ios|check-in-android|ticket-block-mgmt-ios|ticket-block-mgmt-android
// @param string SoftVersion
// @param string HardVersion

type SetAppVersionNumberByTypeParameters struct {
	AppVersionType string
	SoftVersion    string
	HardVersion    string
}

func (t *AppVersion) SetAppVersionNumberByType(p *SetAppVersionNumberByTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`appVersionType`, p.AppVersionType)
	queryParameters.Add(`softVersion`, p.SoftVersion)
	queryParameters.Add(`hardVersion`, p.HardVersion)

	return t.restClient.Post(
		`/v2/AppVersion/UseCase/SetAppVersionNumberByType`,
		&queryParameters,
		nil,
		nil,
	)
}
