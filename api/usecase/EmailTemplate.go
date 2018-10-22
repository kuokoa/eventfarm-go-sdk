/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"

	"github.com/eventfarm/go-sdk"
)

type EmailTemplate struct {
	restClient sdk.RestClientInterface
}

func NewEmailTemplate(restClient sdk.RestClientInterface) *EmailTemplate {
	return &EmailTemplate{restClient}
}

// GET: Queries
// @param string EmailTemplateType simple-template|simple-header|simple-template-border|default-invite|full-width-header

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
