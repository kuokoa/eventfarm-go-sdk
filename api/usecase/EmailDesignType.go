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

type EmailDesignType struct {
	restClient rest.RestClientInterface
}

func NewEmailDesignType(restClient rest.RestClientInterface) *EmailDesignType {
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
