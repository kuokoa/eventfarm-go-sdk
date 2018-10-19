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

type Import struct {
	restClient sdk.RestClientInterface
}

func NewImport(restClient sdk.RestClientInterface) *Import {
	return &Import{restClient}
}

// GET: Queries
// @param string UserImportId
// @param array|null WithData GoodRecords|DuplicateRecords|ErrorRecords|ImportFailureRecords

type GetUserImportParameters struct {
	UserImportId string
	WithData     *[]string
}

func (t *Import) GetUserImport(p *GetUserImportParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, p.UserImportId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Import/UseCase/GetUserImport`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserImportId
// @param string FileId

type GetUserImportFileParameters struct {
	UserImportId string
	FileId       string
}

func (t *Import) GetUserImportFile(p *GetUserImportFileParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, p.UserImportId)
	queryParameters.Add(`fileId`, p.FileId)

	return t.restClient.Get(
		`/v2/Import/UseCase/GetUserImportFile`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string UserImportId
// @param string EventId
// @param string|null StackId
// @param int|null GuestsPerInvitation >= 1
// @param string|null InvitationCreationType unconfirmed-no-email|confirmed-no-email|send-email
// @param string|null GroupName
// @param string|null GroupId
// @param string|null RedirectUrl

type PostProcessAndImportInvitationsParameters struct {
	UserImportId           string
	EventId                string
	StackId                *string
	GuestsPerInvitation    *int
	InvitationCreationType *string
	GroupName              *string
	GroupId                *string
	RedirectUrl            *string
}

func (t *Import) PostProcessAndImportInvitations(p *PostProcessAndImportInvitationsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, p.UserImportId)
	queryParameters.Add(`eventId`, p.EventId)
	if p.StackId != nil {
		queryParameters.Add(`stackId`, *p.StackId)
	}
	if p.GuestsPerInvitation != nil {
		queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(*p.GuestsPerInvitation))
	}
	if p.InvitationCreationType != nil {
		queryParameters.Add(`invitationCreationType`, *p.InvitationCreationType)
	}
	if p.GroupName != nil {
		queryParameters.Add(`groupName`, *p.GroupName)
	}
	if p.GroupId != nil {
		queryParameters.Add(`groupId`, *p.GroupId)
	}
	if p.RedirectUrl != nil {
		queryParameters.Add(`redirectUrl`, *p.RedirectUrl)
	}

	return t.restClient.Post(
		`/v2/Import/UseCase/PostProcessAndImportInvitations`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserImportId
// @param string|null GroupName
// @param string|null GroupId
// @param string|null RedirectUrl

type PostProcessAndImportUsersParameters struct {
	UserImportId string
	GroupName    *string
	GroupId      *string
	RedirectUrl  *string
}

func (t *Import) PostProcessAndImportUsers(p *PostProcessAndImportUsersParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, p.UserImportId)
	if p.GroupName != nil {
		queryParameters.Add(`groupName`, *p.GroupName)
	}
	if p.GroupId != nil {
		queryParameters.Add(`groupId`, *p.GroupId)
	}
	if p.RedirectUrl != nil {
		queryParameters.Add(`redirectUrl`, *p.RedirectUrl)
	}

	return t.restClient.Post(
		`/v2/Import/UseCase/PostProcessAndImportUsers`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string UserId
// @param string PoolId
// @param string Spreadsheet

type PreProcessSpreadsheetForUserImportParameters struct {
	UserId      string
	PoolId      string
	Spreadsheet string
}

func (t *Import) PreProcessSpreadsheetForUserImport(p *PreProcessSpreadsheetForUserImportParameters) (r *http.Response, err error) {
	// TODO
	return
}
