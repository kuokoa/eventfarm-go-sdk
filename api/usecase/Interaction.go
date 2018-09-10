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

type Interaction struct {
	restClient gosdk.RestClientInterface
}

func NewInteraction(restClient gosdk.RestClientInterface) *Interaction {
	return &Interaction{restClient}
}

// GET: Queries
// @param string InvitationId
// @param array|null WithData Event|Stack
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
func (t *Interaction) ListInteractionsForInvitation(InvitationId string, WithData *[]string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Interaction/UseCase/ListInteractionsForInvitation`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string TagId
// @param array|null WithData Event|Stack
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
func (t *Interaction) ListInteractionsForTag(TagId string, WithData *[]string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`tagId`, TagId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/Interaction/UseCase/ListInteractionsForTag`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string InvitationId
// @param string BodyContent
// @param string TagId
// @param string Type
// @param string|null InteractionId
func (t *Interaction) CreateInteraction(InvitationId string, BodyContent string, TagId string, Type string, InteractionId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, InvitationId)
	queryParameters.Add(`bodyContent`, BodyContent)
	queryParameters.Add(`tagId`, TagId)
	queryParameters.Add(`type`, Type)
	if InteractionId != nil {
		queryParameters.Add(`interactionId`, *InteractionId)
	}

	return t.restClient.Post(
		`/v2/Interaction/UseCase/CreateInteraction`,
		&queryParameters,
		nil,
		nil,
	)
}
