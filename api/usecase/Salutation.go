package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"bitbucket.ef.network/go/sdk"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type Salutation struct {
	restClient sdk.RestClientInterface
}

func NewSalutation(restClient sdk.RestClientInterface) *Salutation {
	return &Salutation{restClient}
}

// GET: Queries
// @param array|null Locales english|german|french|hebrew|polish|portuguese|spanish|spanish-south-america|thai|italian|chinese-traditional|chinese-mandarin|japanese|korean

type GetAllSalutationsParameters struct {
	Locales *[]string
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
