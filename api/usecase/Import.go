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

type Import struct {
	restClient rest.RestClientInterface
}

func NewImport(restClient rest.RestClientInterface) *Import {
	return &Import{restClient}
}

// GET: Queries

type GetUserImportParameters struct {
	UserImportId string
	WithData     *[]interface{} // GoodRecords | DuplicateRecords | ErrorRecords | ImportFailureRecords
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

type PostProcessAndImportInvitationsParameters struct {
	UserImportId           string
	EventId                string
	StackId                *string
	GuestsPerInvitation    *int64 // >= 1
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
		queryParameters.Add(`guestsPerInvitation`, strconv.FormatInt(*p.GuestsPerInvitation, 10))
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

type PreProcessSpreadsheetForUserImportParameters struct {
	UserId      string
	PoolId      string
	Spreadsheet string
}

func (t *Import) PreProcessSpreadsheetForUserImport(p *PreProcessSpreadsheetForUserImportParameters) (r *http.Response, err error) {
	// TODO
	return
}
