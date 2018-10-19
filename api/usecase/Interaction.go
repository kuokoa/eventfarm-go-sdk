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

type Interaction struct {
	restClient sdk.RestClientInterface
}

func NewInteraction(restClient sdk.RestClientInterface) *Interaction {
	return &Interaction{restClient}
}

// GET: Queries
// @param string InvitationId
// @param array|null WithData Event|Stack
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500

type ListInteractionsForInvitationParameters struct {
	InvitationId string
	WithData     *[]string
	Page         *int
	ItemsPerPage *int
}

func (t *Interaction) ListInteractionsForInvitation(p *ListInteractionsForInvitationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
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

type ListInteractionsForTagParameters struct {
	TagId        string
	WithData     *[]string
	Page         *int
	ItemsPerPage *int
}

func (t *Interaction) ListInteractionsForTag(p *ListInteractionsForTagParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`tagId`, p.TagId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
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

type CreateInteractionParameters struct {
	InvitationId  string
	BodyContent   string
	TagId         string
	Type          string
	InteractionId *string
}

func (t *Interaction) CreateInteraction(p *CreateInteractionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`invitationId`, p.InvitationId)
	queryParameters.Add(`bodyContent`, p.BodyContent)
	queryParameters.Add(`tagId`, p.TagId)
	queryParameters.Add(`type`, p.Type)
	if p.InteractionId != nil {
		queryParameters.Add(`interactionId`, *p.InteractionId)
	}

	return t.restClient.Post(
		`/v2/Interaction/UseCase/CreateInteraction`,
		&queryParameters,
		nil,
		nil,
	)
}
