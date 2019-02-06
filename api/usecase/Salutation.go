/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk/rest"
)

type Salutation struct {
	restClient rest.RestClientInterface
}

func NewSalutation(restClient rest.RestClientInterface) *Salutation {
	return &Salutation{restClient}
}

// GET: Queries

type GetAllSalutationsParameters struct {
	Locales *[]string // english | german | french | hebrew | polish | portuguese | spanish | spanish-south-america | thai | italian | chinese-traditional | chinese-mandarin | japanese | korean
}

func (t *Salutation) GetAllSalutations(p *GetAllSalutationsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if p.Locales != nil {
		for i := range *p.Locales {
			queryParameters.Add(`locales[]`, (*p.Locales)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Salutation/UseCase/GetAllSalutations`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
