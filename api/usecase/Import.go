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

type Import struct {
	restClient gosdk.RestClientInterface
}

func NewImport(restClient gosdk.RestClientInterface) *Import {
	return &Import{restClient}
}

// GET: Queries
// @param string UserImportId
// @param array|null WithData GoodRecords|DuplicateRecords|ErrorRecords|ImportFailureRecords
func (t *Import) GetUserImport(UserImportId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, UserImportId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *Import) GetUserImportFile(UserImportId string, FileId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, UserImportId)
	queryParameters.Add(`fileId`, FileId)

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
func (t *Import) PostProcessAndImportInvitations(UserImportId string, EventId string, StackId *string, GuestsPerInvitation *int, InvitationCreationType *string, GroupName *string, GroupId *string, RedirectUrl *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, UserImportId)
	queryParameters.Add(`eventId`, EventId)
	if StackId != nil {
		queryParameters.Add(`stackId`, *StackId)
	}
	if GuestsPerInvitation != nil {
		queryParameters.Add(`guestsPerInvitation`, strconv.Itoa(*GuestsPerInvitation))
	}
	if InvitationCreationType != nil {
		queryParameters.Add(`invitationCreationType`, *InvitationCreationType)
	}
	if GroupName != nil {
		queryParameters.Add(`groupName`, *GroupName)
	}
	if GroupId != nil {
		queryParameters.Add(`groupId`, *GroupId)
	}
	if RedirectUrl != nil {
		queryParameters.Add(`redirectUrl`, *RedirectUrl)
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
func (t *Import) PostProcessAndImportUsers(UserImportId string, GroupName *string, GroupId *string, RedirectUrl *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userImportId`, UserImportId)
	if GroupName != nil {
		queryParameters.Add(`groupName`, *GroupName)
	}
	if GroupId != nil {
		queryParameters.Add(`groupId`, *GroupId)
	}
	if RedirectUrl != nil {
		queryParameters.Add(`redirectUrl`, *RedirectUrl)
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
func (t *Import) PreProcessSpreadsheetForUserImport(UserId string, PoolId string, Spreadsheet string) (r *http.Response, err error) {
	// TODO
	return
}
