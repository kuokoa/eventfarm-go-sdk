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

type EmailDesignType struct {
	restClient sdk.RestClientInterface
}

func NewEmailDesignType(restClient sdk.RestClientInterface) *EmailDesignType {
	return &EmailDesignType{restClient}
}

// GET: Queries

type GetAllEmailDesignTypesParameters struct {
}

func (t *EmailDesignType) GetAllEmailDesignTypes(p *GetAllEmailDesignTypesParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Get(
		`/v2/EmailDesignType/UseCase/GetAllEmailDesignTypes`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
