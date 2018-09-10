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

type Salutation struct {
	restClient gosdk.RestClientInterface
}

func NewSalutation(restClient gosdk.RestClientInterface) *Salutation {
	return &Salutation{restClient}
}

// GET: Queries
// @param array|null Locales english|german|french|hebrew|polish|portuguese|spanish|spanish-south-america|thai|italian|chinese-traditional|chinese-mandarin|japanese|korean
func (t *Salutation) GetAllSalutations(Locales *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	if Locales != nil {
		for i := range *Locales {
			queryParameters.Add(`locales[]`, (*Locales)[i])
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
