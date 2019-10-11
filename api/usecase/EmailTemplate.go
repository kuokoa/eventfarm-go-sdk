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

type EmailTemplate struct {
	restClient rest.RestClientInterface
}

func NewEmailTemplate(restClient rest.RestClientInterface) *EmailTemplate {
	return &EmailTemplate{restClient}
}

// GET: Queries

type GetEmailTemplateByTypeParameters struct {
	EmailTemplateType string
}

func (t *EmailTemplate) GetEmailTemplateByType(p *GetEmailTemplateByTypeParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailTemplateType`, p.EmailTemplateType)

	return t.restClient.Get(
		`/v2/EmailTemplate/UseCase/GetEmailTemplateByType`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
