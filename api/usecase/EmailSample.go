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

type EmailSample struct {
	restClient sdk.RestClientInterface
}

func NewEmailSample(restClient sdk.RestClientInterface) *EmailSample {
	return &EmailSample{restClient}
}

// GET: Queries
// @param string EmailDesignId

type GetEmailThumbnailUrlParameters struct {
	EmailDesignId string
}

func (t *EmailSample) GetEmailThumbnailUrl(p *GetEmailThumbnailUrlParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)

	return t.restClient.Get(
		`/v2/EmailSample/UseCase/GetEmailThumbnailUrl`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EmailDesignId
// @param array|null WithData EmailPreview|EmailSpamResult

type GetLatestEmailSampleForDesignParameters struct {
	EmailDesignId string
	WithData      *[]string
}

func (t *EmailSample) GetLatestEmailSampleForDesign(p *GetLatestEmailSampleForDesignParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/EmailSample/UseCase/GetLatestEmailSampleForDesign`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string PreviewUrl
// @param string PreviewClient
// @param string OperatingSystem
// @param string DisplayName
// @param string ThumbnailUrl
// @param string EmailSampleId
// @param string|null EmailPreviewId

type CreateEmailPreviewParameters struct {
	PreviewUrl      string
	PreviewClient   string
	OperatingSystem string
	DisplayName     string
	ThumbnailUrl    string
	EmailSampleId   string
	EmailPreviewId  *string
}

func (t *EmailSample) CreateEmailPreview(p *CreateEmailPreviewParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`previewUrl`, p.PreviewUrl)
	queryParameters.Add(`previewClient`, p.PreviewClient)
	queryParameters.Add(`operatingSystem`, p.OperatingSystem)
	queryParameters.Add(`displayName`, p.DisplayName)
	queryParameters.Add(`thumbnailUrl`, p.ThumbnailUrl)
	queryParameters.Add(`emailSampleId`, p.EmailSampleId)
	if p.EmailPreviewId != nil {
		queryParameters.Add(`emailPreviewId`, *p.EmailPreviewId)
	}

	return t.restClient.Post(
		`/v2/EmailSample/UseCase/CreateEmailPreview`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string HtmlContent
// @param string EmailDesignId
// @param bool OverrideMinimumInterval true|false
// @param string|null EmailSampleId

type CreateEmailSampleParameters struct {
	HtmlContent             string
	EmailDesignId           string
	OverrideMinimumInterval bool
	EmailSampleId           *string
}

func (t *EmailSample) CreateEmailSample(p *CreateEmailSampleParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`htmlContent`, p.HtmlContent)
	queryParameters.Add(`emailDesignId`, p.EmailDesignId)
	queryParameters.Add(`overrideMinimumInterval`, strconv.FormatBool(p.OverrideMinimumInterval))
	if p.EmailSampleId != nil {
		queryParameters.Add(`emailSampleId`, *p.EmailSampleId)
	}

	return t.restClient.Post(
		`/v2/EmailSample/UseCase/CreateEmailSample`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string SpamClient
// @param string TestType
// @param string TestDetails
// @param int IsSpam
// @param string EmailSampleId
// @param string|null EmailSpamResultId

type CreateEmailSpamResultParameters struct {
	SpamClient        string
	TestType          string
	TestDetails       string
	IsSpam            int
	EmailSampleId     string
	EmailSpamResultId *string
}

func (t *EmailSample) CreateEmailSpamResult(p *CreateEmailSpamResultParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`spamClient`, p.SpamClient)
	queryParameters.Add(`testType`, p.TestType)
	queryParameters.Add(`testDetails`, p.TestDetails)
	queryParameters.Add(`isSpam`, strconv.Itoa(p.IsSpam))
	queryParameters.Add(`emailSampleId`, p.EmailSampleId)
	if p.EmailSpamResultId != nil {
		queryParameters.Add(`emailSpamResultId`, *p.EmailSpamResultId)
	}

	return t.restClient.Post(
		`/v2/EmailSample/UseCase/CreateEmailSpamResult`,
		&queryParameters,
		nil,
		nil,
	)
}
