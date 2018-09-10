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

type EmailSample struct {
	restClient gosdk.RestClientInterface
}

func NewEmailSample(restClient gosdk.RestClientInterface) *EmailSample {
	return &EmailSample{restClient}
}

// GET: Queries
// @param string EmailDesignId
func (t *EmailSample) GetEmailThumbnailUrl(EmailDesignId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, EmailDesignId)

	return t.restClient.Get(
		`/v2/EmailSample/UseCase/GetEmailThumbnailUrl`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EmailDesignId
// @param array|null WithData EmailPreview|EmailSpamResult
func (t *EmailSample) GetLatestEmailSampleForDesign(EmailDesignId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailDesignId`, EmailDesignId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *EmailSample) CreateEmailPreview(PreviewUrl string, PreviewClient string, OperatingSystem string, DisplayName string, ThumbnailUrl string, EmailSampleId string, EmailPreviewId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`previewUrl`, PreviewUrl)
	queryParameters.Add(`previewClient`, PreviewClient)
	queryParameters.Add(`operatingSystem`, OperatingSystem)
	queryParameters.Add(`displayName`, DisplayName)
	queryParameters.Add(`thumbnailUrl`, ThumbnailUrl)
	queryParameters.Add(`emailSampleId`, EmailSampleId)
	if EmailPreviewId != nil {
		queryParameters.Add(`emailPreviewId`, *EmailPreviewId)
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
func (t *EmailSample) CreateEmailSample(HtmlContent string, EmailDesignId string, OverrideMinimumInterval bool, EmailSampleId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`htmlContent`, HtmlContent)
	queryParameters.Add(`emailDesignId`, EmailDesignId)
	queryParameters.Add(`overrideMinimumInterval`, strconv.FormatBool(OverrideMinimumInterval))
	if EmailSampleId != nil {
		queryParameters.Add(`emailSampleId`, *EmailSampleId)
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
func (t *EmailSample) CreateEmailSpamResult(SpamClient string, TestType string, TestDetails string, IsSpam int, EmailSampleId string, EmailSpamResultId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`spamClient`, SpamClient)
	queryParameters.Add(`testType`, TestType)
	queryParameters.Add(`testDetails`, TestDetails)
	queryParameters.Add(`isSpam`, strconv.Itoa(IsSpam))
	queryParameters.Add(`emailSampleId`, EmailSampleId)
	if EmailSpamResultId != nil {
		queryParameters.Add(`emailSpamResultId`, *EmailSpamResultId)
	}

	return t.restClient.Post(
		`/v2/EmailSample/UseCase/CreateEmailSpamResult`,
		&queryParameters,
		nil,
		nil,
	)
}
