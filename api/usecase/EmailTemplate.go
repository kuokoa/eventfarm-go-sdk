package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = fmt.Sprint("Foo")
var _ = http.NoBody

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
