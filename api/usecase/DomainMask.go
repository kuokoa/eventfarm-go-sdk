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

type DomainMask struct {
	restClient gosdk.RestClientInterface
}

func NewDomainMask(restClient gosdk.RestClientInterface) *DomainMask {
	return &DomainMask{restClient}
}

// GET: Queries
// @param string PoolId
func (t *DomainMask) GetAllDomainMasksByPool(PoolId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)

	return t.restClient.Get(
		`/v2/DomainMask/UseCase/GetAllDomainMasksByPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string DomainMaskId
func (t *DomainMask) GetDomainMask(DomainMaskId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`domainMaskId`, DomainMaskId)

	return t.restClient.Get(
		`/v2/DomainMask/UseCase/GetDomainMask`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PoolId
// @param string Domain
func (t *DomainMask) CreateDomainMask(PoolId string, Domain string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`domain`, Domain)

	return t.restClient.Post(
		`/v2/DomainMask/UseCase/CreateDomainMask`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string DomainMaskId
func (t *DomainMask) RemoveDomainMask(DomainMaskId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`domainMaskId`, DomainMaskId)

	return t.restClient.Post(
		`/v2/DomainMask/UseCase/RemoveDomainMask`,
		&queryParameters,
		nil,
		nil,
	)
}
