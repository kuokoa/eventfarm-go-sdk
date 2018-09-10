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

type EmailTemplate struct {
	restClient gosdk.RestClientInterface
}

func NewEmailTemplate(restClient gosdk.RestClientInterface) *EmailTemplate {
	return &EmailTemplate{restClient}
}

// GET: Queries
// @param string EmailTemplateType simple-template|simple-header|simple-template-border|default-invite|full-width-header
func (t *EmailTemplate) GetEmailTemplateByType(EmailTemplateType string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailTemplateType`, EmailTemplateType)

	return t.restClient.Get(
		`/v2/EmailTemplate/UseCase/GetEmailTemplateByType`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
