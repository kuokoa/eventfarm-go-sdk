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

type EmailDesign struct {
	restClient gosdk.RestClientInterface
}

func NewEmailDesign(restClient gosdk.RestClientInterface) *EmailDesign {
	return &EmailDesign{restClient}
}

// GET: Queries
// @param string EmailDesignId
func (t *EmailDesign) GetEmailDesign(EmailDesignId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, EmailDesignId)

	return t.restClient.Get(
		`/v2/EmailDesign/UseCase/GetEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EmailDesignImageId
func (t *EmailDesign) GetEmailDesignImage(EmailDesignImageId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignImageId`, EmailDesignImageId)

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
func (t *EmailDesign) ListEmailDesignsByEvent(EventId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
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
func (t *EmailDesign) CreateEmailDesign(Name string, Layout string, FromName string, Subject string, Content string, BackgroundColor string, EmailDesignTypeId string, EventId string, FromEmail *string, ReplyEmail *string, CcEmail *[]string, BccEmail *[]string, DomainMaskId *string, DomainMaskEmail *string, EmailDesignId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, Name)
	queryParameters.Add(`layout`, Layout)
	queryParameters.Add(`fromName`, FromName)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`content`, Content)
	queryParameters.Add(`backgroundColor`, BackgroundColor)
	queryParameters.Add(`emailDesignTypeId`, EmailDesignTypeId)
	queryParameters.Add(`eventId`, EventId)
	if FromEmail != nil {
		queryParameters.Add(`fromEmail`, *FromEmail)
	}
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmail != nil {
		for i := range *CcEmail {
			queryParameters.Add(`ccEmail[]`, (*CcEmail)[i])
		}
	}
	if BccEmail != nil {
		for i := range *BccEmail {
			queryParameters.Add(`bccEmail[]`, (*BccEmail)[i])
		}
	}
	if DomainMaskId != nil {
		queryParameters.Add(`domainMaskId`, *DomainMaskId)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}
	if EmailDesignId != nil {
		queryParameters.Add(`emailDesignId`, *EmailDesignId)
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
func (t *EmailDesign) CreateEmailDesignFromTemplate(Name string, Layout string, FromName string, Subject string, BackgroundColor string, EmailDesignTypeId string, EventId string, EmailTemplateType string, FromEmail *string, ReplyEmail *string, CcEmails *[]string, BccEmails *[]string, DomainMaskId *string, DomainMaskEmail *string, EmailDesignId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`name`, Name)
	queryParameters.Add(`layout`, Layout)
	queryParameters.Add(`fromName`, FromName)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`backgroundColor`, BackgroundColor)
	queryParameters.Add(`emailDesignTypeId`, EmailDesignTypeId)
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`emailTemplateType`, EmailTemplateType)
	if FromEmail != nil {
		queryParameters.Add(`fromEmail`, *FromEmail)
	}
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmails != nil {
		for i := range *CcEmails {
			queryParameters.Add(`ccEmails[]`, (*CcEmails)[i])
		}
	}
	if BccEmails != nil {
		for i := range *BccEmails {
			queryParameters.Add(`bccEmails[]`, (*BccEmails)[i])
		}
	}
	if DomainMaskId != nil {
		queryParameters.Add(`domainMaskId`, *DomainMaskId)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}
	if EmailDesignId != nil {
		queryParameters.Add(`emailDesignId`, *EmailDesignId)
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
func (t *EmailDesign) CreateEmailDesignImage(EventId string, Image string, EmailDesignImageId *string) (r *http.Response, err error) {
	// TODO
	return
}

// @param string OriginalEmailDesignId
// @param string|null DuplicateEventId
// @param string|null DuplicateEmailDesignId
func (t *EmailDesign) DuplicateEmailDesign(OriginalEmailDesignId string, DuplicateEventId *string, DuplicateEmailDesignId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`originalEmailDesignId`, OriginalEmailDesignId)
	if DuplicateEventId != nil {
		queryParameters.Add(`duplicateEventId`, *DuplicateEventId)
	}
	if DuplicateEmailDesignId != nil {
		queryParameters.Add(`duplicateEmailDesignId`, *DuplicateEmailDesignId)
	}

	return t.restClient.Post(
		`/v2/EmailDesign/UseCase/DuplicateEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EmailDesignId
func (t *EmailDesign) RemoveEmailDesign(EmailDesignId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, EmailDesignId)

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
func (t *EmailDesign) UpdateEmailDesign(EmailDesignId string, Name string, Layout string, FromName string, Subject string, Content string, BackgroundColor string, EmailDesignTypeId string, EventId string, FromEmail *string, ReplyEmail *string, CcEmail *[]string, BccEmail *[]string, DomainMaskId *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, EmailDesignId)
	queryParameters.Add(`name`, Name)
	queryParameters.Add(`layout`, Layout)
	queryParameters.Add(`fromName`, FromName)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`content`, Content)
	queryParameters.Add(`backgroundColor`, BackgroundColor)
	queryParameters.Add(`emailDesignTypeId`, EmailDesignTypeId)
	queryParameters.Add(`eventId`, EventId)
	if FromEmail != nil {
		queryParameters.Add(`fromEmail`, *FromEmail)
	}
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmail != nil {
		for i := range *CcEmail {
			queryParameters.Add(`ccEmail[]`, (*CcEmail)[i])
		}
	}
	if BccEmail != nil {
		for i := range *BccEmail {
			queryParameters.Add(`bccEmail[]`, (*BccEmail)[i])
		}
	}
	if DomainMaskId != nil {
		queryParameters.Add(`domainMaskId`, *DomainMaskId)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailDesign/UseCase/UpdateEmailDesign`,
		&queryParameters,
		nil,
		nil,
	)
}
