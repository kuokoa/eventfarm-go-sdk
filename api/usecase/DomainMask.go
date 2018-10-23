/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type DomainMask struct {
	restClient rest.RestClientInterface
}

func NewDomainMask(restClient rest.RestClientInterface) *DomainMask {
	return &DomainMask{restClient}
}

// GET: Queries
// @param string PoolId

type GetAllDomainMasksByPoolParameters struct {
	PoolId string
}

func (t *DomainMask) GetAllDomainMasksByPool(p *GetAllDomainMasksByPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)

	return t.restClient.Get(
		`/v2/DomainMask/UseCase/GetAllDomainMasksByPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string DomainMaskId

type GetDomainMaskParameters struct {
	DomainMaskId string
}

func (t *DomainMask) GetDomainMask(p *GetDomainMaskParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`domainMaskId`, p.DomainMaskId)

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

type CreateDomainMaskParameters struct {
	PoolId string
	Domain string
}

func (t *DomainMask) CreateDomainMask(p *CreateDomainMaskParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`domain`, p.Domain)

	return t.restClient.Post(
		`/v2/DomainMask/UseCase/CreateDomainMask`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string DomainMaskId

type RemoveDomainMaskParameters struct {
	DomainMaskId string
}

func (t *DomainMask) RemoveDomainMask(p *RemoveDomainMaskParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`domainMaskId`, p.DomainMaskId)

	return t.restClient.Post(
		`/v2/DomainMask/UseCase/RemoveDomainMask`,
		&queryParameters,
		nil,
		nil,
	)
}
