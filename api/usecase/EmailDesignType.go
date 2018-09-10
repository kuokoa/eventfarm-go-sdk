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

type EmailDesignType struct {
	restClient gosdk.RestClientInterface
}

func NewEmailDesignType(restClient gosdk.RestClientInterface) *EmailDesignType {
	return &EmailDesignType{restClient}
}

// GET: Queries
func (t *EmailDesignType) GetAllEmailDesignTypes() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/EmailDesignType/UseCase/GetAllEmailDesignTypes`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
