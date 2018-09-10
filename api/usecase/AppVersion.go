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

type AppVersion struct {
	restClient gosdk.RestClientInterface
}

func NewAppVersion(restClient gosdk.RestClientInterface) *AppVersion {
	return &AppVersion{restClient}
}

// GET: Queries
// @param string AppVersionId
func (t *AppVersion) GetAppVersion(AppVersionId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`appVersionId`, AppVersionId)

	return t.restClient.Get(
		`/v2/AppVersion/UseCase/GetAppVersion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string AppVersionType check-in-ios|check-in-android|ticket-block-mgmt-ios|ticket-block-mgmt-android
func (t *AppVersion) GetAppVersionByType(AppVersionType string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`appVersionType`, AppVersionType)

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
func (t *AppVersion) SetAppVersionNumberByType(AppVersionType string, SoftVersion string, HardVersion string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`appVersionType`, AppVersionType)
	queryParameters.Add(`softVersion`, SoftVersion)
	queryParameters.Add(`hardVersion`, HardVersion)

	return t.restClient.Post(
		`/v2/AppVersion/UseCase/SetAppVersionNumberByType`,
		&queryParameters,
		nil,
		nil,
	)
}
