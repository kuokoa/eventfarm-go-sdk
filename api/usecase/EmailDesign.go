package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = http.NoBody

type EmailDesign struct {
	restClient sdk.RestClientInterface
}

func NewEmailDesign(restClient sdk.RestClientInterface) *EmailDesign {
	return &EmailDesign{restClient}
}

// GET: Queries
// @param string EmailDesignId

type GetEmailDesignParameters struct {
	EmailDesignId string
}

func (t *EmailDesign) GetEmailDesign(p *GetEmailDesignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)

	return t.restClient.Get(
		`/v2/EmailDesign/UseCase/GetEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EmailDesignImageId

type GetEmailDesignImageParameters struct {
	EmailDesignImageId string
}

func (t *EmailDesign) GetEmailDesignImage(p *GetEmailDesignImageParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignImageId`, p.EmailDesignImageId)

	return t.restClient.Get(
		`/v2/EmailDesign/UseCase/GetEmailDesignImage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null Page
// @param int|null ItemsPerPage

type ListEmailDesignsByEventParameters struct {
	EventId      string
	Page         *int
	ItemsPerPage *int
}

func (t *EmailDesign) ListEmailDesignsByEvent(p *ListEmailDesignsByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/EmailDesign/UseCase/ListEmailDesignsByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string Name
// @param string Layout
// @param string FromName
// @param string Subject
// @param string Content
// @param string BackgroundColor
// @param string EmailDesignTypeId
// @param string EventId
// @param string|null FromEmail
// @param string|null ReplyEmail
// @param array|null CcEmail
// @param array|null BccEmail
// @param string|null DomainMaskId
// @param string|null DomainMaskEmail
// @param string|null EmailDesignId

type CreateEmailDesignParameters struct {
	Name              string
	Layout            string
	FromName          string
	Subject           string
	Content           string
	BackgroundColor   string
	EmailDesignTypeId string
	EventId           string
	FromEmail         *string
	ReplyEmail        *string
	CcEmail           *[]string
	BccEmail          *[]string
	DomainMaskId      *string
	DomainMaskEmail   *string
	EmailDesignId     *string
}

func (t *EmailDesign) CreateEmailDesign(p *CreateEmailDesignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`layout`, p.Layout)
	queryParameters.Add(`fromName`, p.FromName)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`content`, p.Content)
	queryParameters.Add(`backgroundColor`, p.BackgroundColor)
	queryParameters.Add(`emailDesignTypeId`, p.EmailDesignTypeId)
	queryParameters.Add(`eventId`, p.EventId)
	if p.FromEmail != nil {
		queryParameters.Add(`fromEmail`, *p.FromEmail)
	}
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmail != nil {
		for i := range *p.CcEmail {
			queryParameters.Add(`ccEmail[]`, (*p.CcEmail)[i])
		}
	}
	if p.BccEmail != nil {
		for i := range *p.BccEmail {
			queryParameters.Add(`bccEmail[]`, (*p.BccEmail)[i])
		}
	}
	if p.DomainMaskId != nil {
		queryParameters.Add(`domainMaskId`, *p.DomainMaskId)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}
	if p.EmailDesignId != nil {
		queryParameters.Add(`emailDesignId`, *p.EmailDesignId)
	}

	return t.restClient.Post(
		`/v2/EmailDesign/UseCase/CreateEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string Name
// @param string Layout blank|alt-email-layout
// @param string FromName
// @param string Subject
// @param string BackgroundColor
// @param string EmailDesignTypeId
// @param string EventId
// @param string EmailTemplateType simple-template|simple-header|simple-template-border|default-invite|full-width-header
// @param string|null FromEmail
// @param string|null ReplyEmail
// @param array|null CcEmails
// @param array|null BccEmails
// @param string|null DomainMaskId
// @param string|null DomainMaskEmail
// @param string|null EmailDesignId

type CreateEmailDesignFromTemplateParameters struct {
	Name              string
	Layout            string
	FromName          string
	Subject           string
	BackgroundColor   string
	EmailDesignTypeId string
	EventId           string
	EmailTemplateType string
	FromEmail         *string
	ReplyEmail        *string
	CcEmails          *[]string
	BccEmails         *[]string
	DomainMaskId      *string
	DomainMaskEmail   *string
	EmailDesignId     *string
}

func (t *EmailDesign) CreateEmailDesignFromTemplate(p *CreateEmailDesignFromTemplateParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`layout`, p.Layout)
	queryParameters.Add(`fromName`, p.FromName)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`backgroundColor`, p.BackgroundColor)
	queryParameters.Add(`emailDesignTypeId`, p.EmailDesignTypeId)
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`emailTemplateType`, p.EmailTemplateType)
	if p.FromEmail != nil {
		queryParameters.Add(`fromEmail`, *p.FromEmail)
	}
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmails != nil {
		for i := range *p.CcEmails {
			queryParameters.Add(`ccEmails[]`, (*p.CcEmails)[i])
		}
	}
	if p.BccEmails != nil {
		for i := range *p.BccEmails {
			queryParameters.Add(`bccEmails[]`, (*p.BccEmails)[i])
		}
	}
	if p.DomainMaskId != nil {
		queryParameters.Add(`domainMaskId`, *p.DomainMaskId)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}
	if p.EmailDesignId != nil {
		queryParameters.Add(`emailDesignId`, *p.EmailDesignId)
	}

	return t.restClient.Post(
		`/v2/EmailDesign/UseCase/CreateEmailDesignFromTemplate`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string Image image/jpeg|image/png
// @param string|null EmailDesignImageId

type CreateEmailDesignImageParameters struct {
	EventId            string
	Image              string
	EmailDesignImageId *string
}

func (t *EmailDesign) CreateEmailDesignImage(p *CreateEmailDesignImageParameters) (r *http.Response, err error) {
	// TODO
	return
}

// @param string OriginalEmailDesignId
// @param string|null DuplicateEventId
// @param string|null DuplicateEmailDesignId

type DuplicateEmailDesignParameters struct {
	OriginalEmailDesignId  string
	DuplicateEventId       *string
	DuplicateEmailDesignId *string
}

func (t *EmailDesign) DuplicateEmailDesign(p *DuplicateEmailDesignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`originalEmailDesignId`, p.OriginalEmailDesignId)
	if p.DuplicateEventId != nil {
		queryParameters.Add(`duplicateEventId`, *p.DuplicateEventId)
	}
	if p.DuplicateEmailDesignId != nil {
		queryParameters.Add(`duplicateEmailDesignId`, *p.DuplicateEmailDesignId)
	}

	return t.restClient.Post(
		`/v2/EmailDesign/UseCase/DuplicateEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EmailDesignId

type RemoveEmailDesignParameters struct {
	EmailDesignId string
}

func (t *EmailDesign) RemoveEmailDesign(p *RemoveEmailDesignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)

	return t.restClient.Post(
		`/v2/EmailDesign/UseCase/RemoveEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EmailDesignId
// @param string Name
// @param string Layout
// @param string FromName
// @param string Subject
// @param string Content
// @param string BackgroundColor
// @param string EmailDesignTypeId
// @param string EventId
// @param string|null FromEmail
// @param string|null ReplyEmail
// @param array|null CcEmail
// @param array|null BccEmail
// @param string|null DomainMaskId
// @param string|null DomainMaskEmail

type UpdateEmailDesignParameters struct {
	EmailDesignId     string
	Name              string
	Layout            string
	FromName          string
	Subject           string
	Content           string
	BackgroundColor   string
	EmailDesignTypeId string
	EventId           string
	FromEmail         *string
	ReplyEmail        *string
	CcEmail           *[]string
	BccEmail          *[]string
	DomainMaskId      *string
	DomainMaskEmail   *string
}

func (t *EmailDesign) UpdateEmailDesign(p *UpdateEmailDesignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	queryParameters.Add(`name`, p.Name)
	queryParameters.Add(`layout`, p.Layout)
	queryParameters.Add(`fromName`, p.FromName)
	queryParameters.Add(`subject`, p.Subject)
	queryParameters.Add(`content`, p.Content)
	queryParameters.Add(`backgroundColor`, p.BackgroundColor)
	queryParameters.Add(`emailDesignTypeId`, p.EmailDesignTypeId)
	queryParameters.Add(`eventId`, p.EventId)
	if p.FromEmail != nil {
		queryParameters.Add(`fromEmail`, *p.FromEmail)
	}
	if p.ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *p.ReplyEmail)
	}
	if p.CcEmail != nil {
		for i := range *p.CcEmail {
			queryParameters.Add(`ccEmail[]`, (*p.CcEmail)[i])
		}
	}
	if p.BccEmail != nil {
		for i := range *p.BccEmail {
			queryParameters.Add(`bccEmail[]`, (*p.BccEmail)[i])
		}
	}
	if p.DomainMaskId != nil {
		queryParameters.Add(`domainMaskId`, *p.DomainMaskId)
	}
	if p.DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *p.DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailDesign/UseCase/UpdateEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}
