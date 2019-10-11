/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/eventfarm/go-sdk/rest"
)

type EmailDesign struct {
	restClient rest.RestClientInterface
}

func NewEmailDesign(restClient rest.RestClientInterface) *EmailDesign {
	return &EmailDesign{restClient}
}

// GET: Queries

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

type ListEmailDesignsByEventParameters struct {
	EventId      string
	Page         *int64
	ItemsPerPage *int64
}

func (t *EmailDesign) ListEmailDesignsByEvent(p *ListEmailDesignsByEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}

	return t.restClient.Get(
		`/v2/EmailDesign/UseCase/ListEmailDesignsByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

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

type CreateEmailDesignImageParameters struct {
	EventId            string
	Image              string
	EmailDesignImageId *string
}

func (t *EmailDesign) CreateEmailDesignImage(p *CreateEmailDesignImageParameters) (r *http.Response, err error) {
	// TODO
	return
}

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
