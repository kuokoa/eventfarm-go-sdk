/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/kuokoa/eventfarm-go-sdk/rest"
)

type Event struct {
	restClient rest.RestClientInterface
}

func NewEvent(restClient rest.RestClientInterface) *Event {
	return &Event{restClient}
}

// GET: Queries

type CheckAltKeywordAvailabilityParameters struct {
	AltKeyword string
}

func (t *Event) CheckAltKeywordAvailability(p *CheckAltKeywordAvailabilityParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`altKeyword`, p.AltKeyword)

	return t.restClient.Get(
		`/v2/Event/UseCase/CheckAltKeywordAvailability`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetAllQuestionsForEventParameters struct {
	EventId  string
	WithData *[]string // Answers | TicketType | QuestionContexts
}

func (t *Event) GetAllQuestionsForEvent(p *GetAllQuestionsForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/GetAllQuestionsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetEventParameters struct {
	EventId  string
	WithData *[]string // Pool | Stacks | StacksWithAvailabilityCounts | Tags | EventTexts | TicketTypes | TicketBlocks | TicketBlocksWithAllotmentCounts | QuestionsAndAnswers | QuestionContext | AllQuestions | ParentEvent | PoolFeatures | EventTheme
}

func (t *Event) GetEvent(p *GetEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/GetEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetEventByAltKeywordParameters struct {
	AltKeyword string
	WithData   *[]string // Pool | Stacks | StacksWithAvailabilityCounts | Tags | EventTexts | TicketTypes | TicketBlocks | TicketBlocksWithAllotmentCounts | QuestionsAndAnswers | QuestionContext | AllQuestions | ParentEvent | PoolFeatures | EventTheme
}

func (t *Event) GetEventByAltKeyword(p *GetEventByAltKeywordParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`altKeyword`, p.AltKeyword)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/GetEventByAltKeyword`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetEventCountsForPoolParameters struct {
	PoolId            string
	EarliestStartDate *int64
}

func (t *Event) GetEventCountsForPool(p *GetEventCountsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.EarliestStartDate != nil {
		queryParameters.Add(`earliestStartDate`, strconv.FormatInt(*p.EarliestStartDate, 10))
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/GetEventCountsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type GetQuestionParameters struct {
	QuestionId string
	WithData   *[]string // Answers | TicketType | QuestionResponseCounts | AnswerQuestionResponseCounts | QuestionContexts
}

func (t *Event) GetQuestion(p *GetQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/GetQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListChildrenForEventParameters struct {
	ParentEventId           string
	Query                   *string
	WithData                *[]string // Pool | Stacks | Tags | TicketTypes | TicketBlocks | EventTexts | QuestionsAndAnswers | ThumbnailUrl
	Page                    *int64    // >= 1
	ItemsPerPage            *int64    // 1-100
	SortBy                  *string
	SortDirection           *string
	EventDateFilterType     *string
	PoolId                  *string
	Tags                    *[]string
	AttributesFilter        *[]string
	AttributesExcludeFilter *[]string
}

func (t *Event) ListChildrenForEvent(p *ListChildrenForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`parentEventId`, p.ParentEventId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Tags != nil {
		for i := range *p.Tags {
			queryParameters.Add(`tags[]`, (*p.Tags)[i])
		}
	}
	if p.AttributesFilter != nil {
		for i := range *p.AttributesFilter {
			queryParameters.Add(`attributesFilter[]`, (*p.AttributesFilter)[i])
		}
	}
	if p.AttributesExcludeFilter != nil {
		for i := range *p.AttributesExcludeFilter {
			queryParameters.Add(`attributesExcludeFilter[]`, (*p.AttributesExcludeFilter)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/ListChildrenForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListChildrenForEventForUserParameters struct {
	ParentEventId       string
	UserId              string
	Query               *string
	WithData            *[]string // Pool | Stacks | Tags | TicketTypes | TicketBlocks | QuestionsAndAnswers | ThumbnailUrl
	Page                *int64    // >= 1
	ItemsPerPage        *int64    // 1-100
	SortBy              *string
	SortDirection       *string
	EventDateFilterType *string
	PoolId              *string
	Tags                *[]string
}

func (t *Event) ListChildrenForEventForUser(p *ListChildrenForEventForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`parentEventId`, p.ParentEventId)
	queryParameters.Add(`userId`, p.UserId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Tags != nil {
		for i := range *p.Tags {
			queryParameters.Add(`tags[]`, (*p.Tags)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/ListChildrenForEventForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListEventsForPoolParameters struct {
	PoolId                  string
	Query                   *string
	AttributesFilter        *[]string // distribute | donate | fee | editname | reveal | allow-notes | duplicate-emails | navigation | social-media | social-media-bar | map-location | show-description | ipad-purchase | simple-layout | label-print | skip-event-allocate-display | geo-restrict | visa-checkout | archived | guest-can-change-response | efx-enabled | show-calendar | show-qr-confirmation | event-app-enabled | child-events-enabled
	AttributesExcludeFilter *[]string // distribute | donate | fee | editname | reveal | allow-notes | duplicate-emails | navigation | social-media | social-media-bar | map-location | show-description | ipad-purchase | simple-layout | label-print | skip-event-allocate-display | geo-restrict | visa-checkout | archived | guest-can-change-response | efx-enabled | show-calendar | show-qr-confirmation | event-app-enabled | child-events-enabled
	WithData                *[]string // Pool | Stacks | Tags | TicketTypes | TicketBlocks | QuestionsAndAnswers | ThumbnailUrl
	LastModifiedTimestamp   *int64
	Page                    *int64 // >= 1
	ItemsPerPage            *int64 // 1-500
	SortBy                  *string
	SortDirection           *string
	EventDateFilterType     *string
	Tags                    *[]string
	EarliestStartTimestamp  *int64 // >= 0
}

func (t *Event) ListEventsForPool(p *ListEventsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.AttributesFilter != nil {
		for i := range *p.AttributesFilter {
			queryParameters.Add(`attributesFilter[]`, (*p.AttributesFilter)[i])
		}
	}
	if p.AttributesExcludeFilter != nil {
		for i := range *p.AttributesExcludeFilter {
			queryParameters.Add(`attributesExcludeFilter[]`, (*p.AttributesExcludeFilter)[i])
		}
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.Tags != nil {
		for i := range *p.Tags {
			queryParameters.Add(`tags[]`, (*p.Tags)[i])
		}
	}
	if p.EarliestStartTimestamp != nil {
		queryParameters.Add(`earliestStartTimestamp`, strconv.FormatInt(*p.EarliestStartTimestamp, 10))
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/ListEventsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListEventsForUserParameters struct {
	UserId                  string
	Query                   *string
	AttributesFilter        *[]string // distribute | donate | fee | editname | reveal | allow-notes | duplicate-emails | navigation | social-media | social-media-bar | map-location | show-description | ipad-purchase | simple-layout | label-print | skip-event-allocate-display | geo-restrict | visa-checkout | archived | guest-can-change-response | efx-enabled | show-calendar | show-qr-confirmation | event-app-enabled | child-events-enabled
	AttributesExcludeFilter *[]string // distribute | donate | fee | editname | reveal | allow-notes | duplicate-emails | navigation | social-media | social-media-bar | map-location | show-description | ipad-purchase | simple-layout | label-print | skip-event-allocate-display | geo-restrict | visa-checkout | archived | guest-can-change-response | efx-enabled | show-calendar | show-qr-confirmation | event-app-enabled | child-events-enabled
	WithData                *[]string // Pool | Stacks | Tags | TicketTypes | TicketBlocks | QuestionsAndAnswers | ThumbnailUrl
	LastModifiedTimestamp   *int64
	Page                    *int64 // >= 1
	ItemsPerPage            *int64 // 1-500
	SortBy                  *string
	SortDirection           *string
	EventDateFilterType     *string
	PoolId                  *string
	Tags                    *[]string
	EarliestStartTimestamp  *int64 // >= 0
	OrganizerOnly           *bool
	ParentEventsOnly        *bool
}

func (t *Event) ListEventsForUser(p *ListEventsForUserParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, p.UserId)
	if p.Query != nil {
		queryParameters.Add(`query`, *p.Query)
	}
	if p.AttributesFilter != nil {
		for i := range *p.AttributesFilter {
			queryParameters.Add(`attributesFilter[]`, (*p.AttributesFilter)[i])
		}
	}
	if p.AttributesExcludeFilter != nil {
		for i := range *p.AttributesExcludeFilter {
			queryParameters.Add(`attributesExcludeFilter[]`, (*p.AttributesExcludeFilter)[i])
		}
	}
	if p.WithData != nil {
		for i := range *p.WithData {
			queryParameters.Add(`withData[]`, (*p.WithData)[i])
		}
	}
	if p.LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.FormatInt(*p.LastModifiedTimestamp, 10))
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.FormatInt(*p.Page, 10))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.FormatInt(*p.ItemsPerPage, 10))
	}
	if p.SortBy != nil {
		queryParameters.Add(`sortBy`, *p.SortBy)
	}
	if p.SortDirection != nil {
		queryParameters.Add(`sortDirection`, *p.SortDirection)
	}
	if p.EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *p.EventDateFilterType)
	}
	if p.PoolId != nil {
		queryParameters.Add(`poolId`, *p.PoolId)
	}
	if p.Tags != nil {
		for i := range *p.Tags {
			queryParameters.Add(`tags[]`, (*p.Tags)[i])
		}
	}
	if p.EarliestStartTimestamp != nil {
		queryParameters.Add(`earliestStartTimestamp`, strconv.FormatInt(*p.EarliestStartTimestamp, 10))
	}
	if p.OrganizerOnly != nil {
		queryParameters.Add(`organizerOnly`, strconv.FormatBool(*p.OrganizerOnly))
	}
	if p.ParentEventsOnly != nil {
		queryParameters.Add(`parentEventsOnly`, strconv.FormatBool(*p.ParentEventsOnly))
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/ListEventsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

type ListQuestionsByEventAndContextParameters struct {
	EventId              string
	QuestionContextTypes *[]string // registration | lead
}

func (t *Event) ListQuestionsByEventAndContext(p *ListQuestionsByEventAndContextParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	if p.QuestionContextTypes != nil {
		for i := range *p.QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*p.QuestionContextTypes)[i])
		}
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/ListQuestionsByEventAndContext`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands

type AddChildEventParameters struct {
	EventId      string
	ChildEventId string
}

func (t *Event) AddChildEvent(p *AddChildEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`childEventId`, p.ChildEventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/AddChildEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) AddChildEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/AddChildEvent`,
		data,
		nil,
		nil,
	)
}

type AddParentToEventParameters struct {
	EventId       string
	ParentEventId string
}

func (t *Event) AddParentToEvent(p *AddParentToEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`parentEventId`, p.ParentEventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/AddParentToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) AddParentToEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/AddParentToEvent`,
		data,
		nil,
		nil,
	)
}

type AddTagToEventParameters struct {
	EventId string
	TagName string
}

func (t *Event) AddTagToEvent(p *AddTagToEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`tagName`, p.TagName)

	return t.restClient.Post(
		`/v2/Event/UseCase/AddTagToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) AddTagToEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/AddTagToEvent`,
		data,
		nil,
		nil,
	)
}

type AddUserRoleToEventParameters struct {
	EventId             string
	Email               string
	FirstName           string
	LastName            string
	EventRole           string
	AuthenticatedUserId *string
}

func (t *Event) AddUserRoleToEvent(p *AddUserRoleToEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`email`, p.Email)
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	queryParameters.Add(`eventRole`, p.EventRole)
	if p.AuthenticatedUserId != nil {
		queryParameters.Add(`authenticatedUserId`, *p.AuthenticatedUserId)
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/AddUserRoleToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) AddUserRoleToEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/AddUserRoleToEvent`,
		data,
		nil,
		nil,
	)
}

type ArchiveEventParameters struct {
	EventId string
}

func (t *Event) ArchiveEvent(p *ArchiveEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/ArchiveEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) ArchiveEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/ArchiveEvent`,
		data,
		nil,
		nil,
	)
}

type ClearDefaultSitePageForEventParameters struct {
	EventId string
}

func (t *Event) ClearDefaultSitePageForEvent(p *ClearDefaultSitePageForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/ClearDefaultSitePageForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) ClearDefaultSitePageForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/ClearDefaultSitePageForEvent`,
		data,
		nil,
		nil,
	)
}

type CopyExistingEventConfigurationParameters struct {
	FromEventId                   string
	ToEventId                     string
	ShouldCopyUsers               *bool
	ShouldCopyEmailDesigns        *bool
	ShouldCopySettings            *bool
	ShouldCopyRegQuestions        *bool
	ShouldCopySitePages           *bool
	ShouldCopyTicketTypes         *bool
	ShouldCopyTicketBlocks        *bool
	ShouldCopyStacks              *bool
	ShouldCopyIntegrationSettings *bool
}

func (t *Event) CopyExistingEventConfiguration(p *CopyExistingEventConfigurationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`fromEventId`, p.FromEventId)
	queryParameters.Add(`toEventId`, p.ToEventId)
	if p.ShouldCopyUsers != nil {
		queryParameters.Add(`shouldCopyUsers`, strconv.FormatBool(*p.ShouldCopyUsers))
	}
	if p.ShouldCopyEmailDesigns != nil {
		queryParameters.Add(`shouldCopyEmailDesigns`, strconv.FormatBool(*p.ShouldCopyEmailDesigns))
	}
	if p.ShouldCopySettings != nil {
		queryParameters.Add(`shouldCopySettings`, strconv.FormatBool(*p.ShouldCopySettings))
	}
	if p.ShouldCopyRegQuestions != nil {
		queryParameters.Add(`shouldCopyRegQuestions`, strconv.FormatBool(*p.ShouldCopyRegQuestions))
	}
	if p.ShouldCopySitePages != nil {
		queryParameters.Add(`shouldCopySitePages`, strconv.FormatBool(*p.ShouldCopySitePages))
	}
	if p.ShouldCopyTicketTypes != nil {
		queryParameters.Add(`shouldCopyTicketTypes`, strconv.FormatBool(*p.ShouldCopyTicketTypes))
	}
	if p.ShouldCopyTicketBlocks != nil {
		queryParameters.Add(`shouldCopyTicketBlocks`, strconv.FormatBool(*p.ShouldCopyTicketBlocks))
	}
	if p.ShouldCopyStacks != nil {
		queryParameters.Add(`shouldCopyStacks`, strconv.FormatBool(*p.ShouldCopyStacks))
	}
	if p.ShouldCopyIntegrationSettings != nil {
		queryParameters.Add(`shouldCopyIntegrationSettings`, strconv.FormatBool(*p.ShouldCopyIntegrationSettings))
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/CopyExistingEventConfiguration`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) CopyExistingEventConfigurationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/CopyExistingEventConfiguration`,
		data,
		nil,
		nil,
	)
}

type CreateAnswerParameters struct {
	QuestionId string
	Text       string
	SortOrder  *int64
	IsDefault  *bool
	AnswerId   *string
}

func (t *Event) CreateAnswer(p *CreateAnswerParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`text`, p.Text)
	if p.SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.FormatInt(*p.SortOrder, 10))
	}
	if p.IsDefault != nil {
		queryParameters.Add(`isDefault`, strconv.FormatBool(*p.IsDefault))
	}
	if p.AnswerId != nil {
		queryParameters.Add(`answerId`, *p.AnswerId)
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/CreateAnswer`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) CreateAnswerWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/CreateAnswer`,
		data,
		nil,
		nil,
	)
}

type CreateCIOEventParameters struct {
	PoolId    string
	UserId    string
	EventName string
	StartTime *string
	EndTime   *string
	Timezone  *string
	EventId   *string
}

func (t *Event) CreateCIOEvent(p *CreateCIOEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`eventName`, p.EventName)
	if p.StartTime != nil {
		queryParameters.Add(`startTime`, *p.StartTime)
	}
	if p.EndTime != nil {
		queryParameters.Add(`endTime`, *p.EndTime)
	}
	if p.Timezone != nil {
		queryParameters.Add(`timezone`, *p.Timezone)
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/CreateCIOEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) CreateCIOEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/CreateCIOEvent`,
		data,
		nil,
		nil,
	)
}

type CreateEventParameters struct {
	PoolId       string
	UserId       string
	EventName    string
	AltKeyword   *string
	ContactEmail *string
	StartTime    *string
	EndTime      *string
	Timezone     *string
	EventId      *string
}

func (t *Event) CreateEvent(p *CreateEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`eventName`, p.EventName)
	if p.AltKeyword != nil {
		queryParameters.Add(`altKeyword`, *p.AltKeyword)
	}
	if p.ContactEmail != nil {
		queryParameters.Add(`contactEmail`, *p.ContactEmail)
	}
	if p.StartTime != nil {
		queryParameters.Add(`startTime`, *p.StartTime)
	}
	if p.EndTime != nil {
		queryParameters.Add(`endTime`, *p.EndTime)
	}
	if p.Timezone != nil {
		queryParameters.Add(`timezone`, *p.Timezone)
	}
	if p.EventId != nil {
		queryParameters.Add(`eventId`, *p.EventId)
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/CreateEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) CreateEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/CreateEvent`,
		data,
		nil,
		nil,
	)
}

type CreateQuestionParameters struct {
	EventId              string
	Text                 string
	QuestionType         string
	SortOrder            *int64
	IsRequired           *bool
	IsIndividual         *bool
	TicketTypeId         *string
	QuestionId           *string
	QuestionContextTypes *[]string // registration | lead
}

func (t *Event) CreateQuestion(p *CreateQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`text`, p.Text)
	queryParameters.Add(`questionType`, p.QuestionType)
	if p.SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.FormatInt(*p.SortOrder, 10))
	}
	if p.IsRequired != nil {
		queryParameters.Add(`isRequired`, strconv.FormatBool(*p.IsRequired))
	}
	if p.IsIndividual != nil {
		queryParameters.Add(`isIndividual`, strconv.FormatBool(*p.IsIndividual))
	}
	if p.TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *p.TicketTypeId)
	}
	if p.QuestionId != nil {
		queryParameters.Add(`questionId`, *p.QuestionId)
	}
	if p.QuestionContextTypes != nil {
		for i := range *p.QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*p.QuestionContextTypes)[i])
		}
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/CreateQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) CreateQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/CreateQuestion`,
		data,
		nil,
		nil,
	)
}

type DeleteAnswerParameters struct {
	AnswerId string
}

func (t *Event) DeleteAnswer(p *DeleteAnswerParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DeleteAnswer`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DeleteAnswerWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DeleteAnswer`,
		data,
		nil,
		nil,
	)
}

type DeleteQuestionParameters struct {
	QuestionId string
}

func (t *Event) DeleteQuestion(p *DeleteQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DeleteQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DeleteQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DeleteQuestion`,
		data,
		nil,
		nil,
	)
}

type DisableAmexCardParameters struct {
	EventId string
}

func (t *Event) DisableAmexCard(p *DisableAmexCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableAmexCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableAmexCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableAmexCard`,
		data,
		nil,
		nil,
	)
}

type DisableCanvasParameters struct {
	EventId string
}

func (t *Event) DisableCanvas(p *DisableCanvasParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableCanvas`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableCanvasWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableCanvas`,
		data,
		nil,
		nil,
	)
}

type DisableDiscoverCardParameters struct {
	EventId string
}

func (t *Event) DisableDiscoverCard(p *DisableDiscoverCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDiscoverCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableDiscoverCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableDiscoverCard`,
		data,
		nil,
		nil,
	)
}

type DisableDistributionParameters struct {
	EventId string
}

func (t *Event) DisableDistribution(p *DisableDistributionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDistribution`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableDistributionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableDistribution`,
		data,
		nil,
		nil,
	)
}

type DisableDonationParameters struct {
	EventId string
}

func (t *Event) DisableDonation(p *DisableDonationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDonation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableDonationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableDonation`,
		data,
		nil,
		nil,
	)
}

type DisableDuplicateEmailsParameters struct {
	EventId string
}

func (t *Event) DisableDuplicateEmails(p *DisableDuplicateEmailsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDuplicateEmails`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableDuplicateEmailsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableDuplicateEmails`,
		data,
		nil,
		nil,
	)
}

type DisableEditNameParameters struct {
	EventId string
}

func (t *Event) DisableEditName(p *DisableEditNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableEditName`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableEditNameWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableEditName`,
		data,
		nil,
		nil,
	)
}

type DisableEfxParameters struct {
	EventId string
}

func (t *Event) DisableEfx(p *DisableEfxParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableEfx`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableEfxWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableEfx`,
		data,
		nil,
		nil,
	)
}

type DisableEventAppParameters struct {
	EventId string
}

func (t *Event) DisableEventApp(p *DisableEventAppParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableEventApp`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableEventAppWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableEventApp`,
		data,
		nil,
		nil,
	)
}

type DisableGuestCanChangeResponseParameters struct {
	EventId string
}

func (t *Event) DisableGuestCanChangeResponse(p *DisableGuestCanChangeResponseParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableGuestCanChangeResponse`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableGuestCanChangeResponseWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableGuestCanChangeResponse`,
		data,
		nil,
		nil,
	)
}

type DisableInvitationRevealParameters struct {
	EventId string
}

func (t *Event) DisableInvitationReveal(p *DisableInvitationRevealParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableInvitationReveal`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableInvitationRevealWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableInvitationReveal`,
		data,
		nil,
		nil,
	)
}

type DisableMastercardCardParameters struct {
	EventId string
}

func (t *Event) DisableMastercardCard(p *DisableMastercardCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableMastercardCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableMastercardCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableMastercardCard`,
		data,
		nil,
		nil,
	)
}

type DisableQRCodeConfirmationParameters struct {
	EventId string
}

func (t *Event) DisableQRCodeConfirmation(p *DisableQRCodeConfirmationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableQRCodeConfirmation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableQRCodeConfirmationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableQRCodeConfirmation`,
		data,
		nil,
		nil,
	)
}

type DisableQuestionParameters struct {
	QuestionId string
}

func (t *Event) DisableQuestion(p *DisableQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableQuestion`,
		data,
		nil,
		nil,
	)
}

type DisableUserFieldParameters struct {
	EventId string
	Field   string
}

func (t *Event) DisableUserField(p *DisableUserFieldParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`field`, p.Field)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableUserField`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableUserFieldWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableUserField`,
		data,
		nil,
		nil,
	)
}

type DisableVisaCardParameters struct {
	EventId string
}

func (t *Event) DisableVisaCard(p *DisableVisaCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableVisaCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) DisableVisaCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/DisableVisaCard`,
		data,
		nil,
		nil,
	)
}

type EnableAmexCardParameters struct {
	EventId string
}

func (t *Event) EnableAmexCard(p *EnableAmexCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableAmexCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableAmexCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableAmexCard`,
		data,
		nil,
		nil,
	)
}

type EnableCanvasParameters struct {
	EventId string
}

func (t *Event) EnableCanvas(p *EnableCanvasParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableCanvas`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableCanvasWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableCanvas`,
		data,
		nil,
		nil,
	)
}

type EnableDiscoverCardParameters struct {
	EventId string
}

func (t *Event) EnableDiscoverCard(p *EnableDiscoverCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDiscoverCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableDiscoverCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableDiscoverCard`,
		data,
		nil,
		nil,
	)
}

type EnableDistributionParameters struct {
	EventId string
}

func (t *Event) EnableDistribution(p *EnableDistributionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDistribution`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableDistributionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableDistribution`,
		data,
		nil,
		nil,
	)
}

type EnableDonationParameters struct {
	EventId string
}

func (t *Event) EnableDonation(p *EnableDonationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDonation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableDonationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableDonation`,
		data,
		nil,
		nil,
	)
}

type EnableDuplicateEmailsParameters struct {
	EventId string
}

func (t *Event) EnableDuplicateEmails(p *EnableDuplicateEmailsParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDuplicateEmails`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableDuplicateEmailsWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableDuplicateEmails`,
		data,
		nil,
		nil,
	)
}

type EnableEditNameParameters struct {
	EventId string
}

func (t *Event) EnableEditName(p *EnableEditNameParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableEditName`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableEditNameWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableEditName`,
		data,
		nil,
		nil,
	)
}

type EnableEfxParameters struct {
	EventId string
}

func (t *Event) EnableEfx(p *EnableEfxParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableEfx`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableEfxWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableEfx`,
		data,
		nil,
		nil,
	)
}

type EnableEventAppParameters struct {
	EventId string
}

func (t *Event) EnableEventApp(p *EnableEventAppParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableEventApp`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableEventAppWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableEventApp`,
		data,
		nil,
		nil,
	)
}

type EnableGuestCanChangeResponseParameters struct {
	EventId string
}

func (t *Event) EnableGuestCanChangeResponse(p *EnableGuestCanChangeResponseParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableGuestCanChangeResponse`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableGuestCanChangeResponseWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableGuestCanChangeResponse`,
		data,
		nil,
		nil,
	)
}

type EnableInvitationRevealParameters struct {
	EventId string
}

func (t *Event) EnableInvitationReveal(p *EnableInvitationRevealParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableInvitationReveal`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableInvitationRevealWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableInvitationReveal`,
		data,
		nil,
		nil,
	)
}

type EnableMastercardCardParameters struct {
	EventId string
}

func (t *Event) EnableMastercardCard(p *EnableMastercardCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableMastercardCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableMastercardCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableMastercardCard`,
		data,
		nil,
		nil,
	)
}

type EnableQRCodeConfirmationParameters struct {
	EventId string
}

func (t *Event) EnableQRCodeConfirmation(p *EnableQRCodeConfirmationParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableQRCodeConfirmation`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableQRCodeConfirmationWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableQRCodeConfirmation`,
		data,
		nil,
		nil,
	)
}

type EnableQuestionParameters struct {
	QuestionId string
}

func (t *Event) EnableQuestion(p *EnableQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableQuestion`,
		data,
		nil,
		nil,
	)
}

type EnableUserFieldParameters struct {
	EventId string
	Field   string
}

func (t *Event) EnableUserField(p *EnableUserFieldParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`field`, p.Field)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableUserField`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableUserFieldWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableUserField`,
		data,
		nil,
		nil,
	)
}

type EnableVisaCardParameters struct {
	EventId string
}

func (t *Event) EnableVisaCard(p *EnableVisaCardParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableVisaCard`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) EnableVisaCardWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/EnableVisaCard`,
		data,
		nil,
		nil,
	)
}

type IncrementCIOAndMessageCSMParameters struct {
	PoolId        string
	UserId        string
	FirstName     string
	LastName      string
	Email         string
	Company       string
	PhoneNumber   string
	HowManyEvents string
	Industry      string
	TypeOfEvents  []string // brand-awareness-events | trade-shows | conferences | roadshows | internal-meetings | recruiting-events | fundraising-events
}

func (t *Event) IncrementCIOAndMessageCSM(p *IncrementCIOAndMessageCSMParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	queryParameters.Add(`userId`, p.UserId)
	queryParameters.Add(`firstName`, p.FirstName)
	queryParameters.Add(`lastName`, p.LastName)
	queryParameters.Add(`email`, p.Email)
	queryParameters.Add(`company`, p.Company)
	queryParameters.Add(`phoneNumber`, p.PhoneNumber)
	queryParameters.Add(`howManyEvents`, p.HowManyEvents)
	queryParameters.Add(`industry`, p.Industry)
	for i := range p.TypeOfEvents {
		queryParameters.Add(`typeOfEvents[]`, p.TypeOfEvents[i])
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/IncrementCIOAndMessageCSM`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) IncrementCIOAndMessageCSMWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/IncrementCIOAndMessageCSM`,
		data,
		nil,
		nil,
	)
}

type RemoveChildEventParameters struct {
	EventId      string
	ChildEventId string
}

func (t *Event) RemoveChildEvent(p *RemoveChildEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`childEventId`, p.ChildEventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/RemoveChildEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) RemoveChildEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/RemoveChildEvent`,
		data,
		nil,
		nil,
	)
}

type RemoveMessageForEventParameters struct {
	EventId     string
	MessageType string
}

func (t *Event) RemoveMessageForEvent(p *RemoveMessageForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`messageType`, p.MessageType)

	return t.restClient.Post(
		`/v2/Event/UseCase/RemoveMessageForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) RemoveMessageForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/RemoveMessageForEvent`,
		data,
		nil,
		nil,
	)
}

type RemoveParentFromEventParameters struct {
	EventId string
}

func (t *Event) RemoveParentFromEvent(p *RemoveParentFromEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/RemoveParentFromEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) RemoveParentFromEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/RemoveParentFromEvent`,
		data,
		nil,
		nil,
	)
}

type RemoveTagFromEventParameters struct {
	EventId string
	TagSlug string
}

func (t *Event) RemoveTagFromEvent(p *RemoveTagFromEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`tagSlug`, p.TagSlug)

	return t.restClient.Post(
		`/v2/Event/UseCase/RemoveTagFromEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) RemoveTagFromEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/RemoveTagFromEvent`,
		data,
		nil,
		nil,
	)
}

type RemoveTrackingScriptForEventParameters struct {
	EventId            string
	TrackingScriptType string
}

func (t *Event) RemoveTrackingScriptForEvent(p *RemoveTrackingScriptForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`trackingScriptType`, p.TrackingScriptType)

	return t.restClient.Post(
		`/v2/Event/UseCase/RemoveTrackingScriptForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) RemoveTrackingScriptForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/RemoveTrackingScriptForEvent`,
		data,
		nil,
		nil,
	)
}

type SendContactEmailParameters struct {
	EventId      string
	FromName     string
	FromEmail    string
	EmailSubject string
	EmailBody    string
}

func (t *Event) SendContactEmail(p *SendContactEmailParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`fromName`, p.FromName)
	queryParameters.Add(`fromEmail`, p.FromEmail)
	queryParameters.Add(`emailSubject`, p.EmailSubject)
	queryParameters.Add(`emailBody`, p.EmailBody)

	return t.restClient.Post(
		`/v2/Event/UseCase/SendContactEmail`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SendContactEmailWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SendContactEmail`,
		data,
		nil,
		nil,
	)
}

func (t *Event) SendMessageToGuestList() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Post(
		`/v2/Event/UseCase/SendMessageToGuestList`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SendMessageToGuestListWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SendMessageToGuestList`,
		data,
		nil,
		nil,
	)
}

type SetAltKeywordForEventParameters struct {
	EventId    string
	AltKeyword string
}

func (t *Event) SetAltKeywordForEvent(p *SetAltKeywordForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`altKeyword`, p.AltKeyword)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetAltKeywordForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetAltKeywordForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetAltKeywordForEvent`,
		data,
		nil,
		nil,
	)
}

type SetAnswerSortOrderParameters struct {
	AnswerId  string
	SortOrder int64
}

func (t *Event) SetAnswerSortOrder(p *SetAnswerSortOrderParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)
	queryParameters.Add(`sortOrder`, strconv.FormatInt(p.SortOrder, 10))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetAnswerSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetAnswerSortOrderWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetAnswerSortOrder`,
		data,
		nil,
		nil,
	)
}

type SetContactEmailForEventParameters struct {
	EventId      string
	ContactEmail string
}

func (t *Event) SetContactEmailForEvent(p *SetContactEmailForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`contactEmail`, p.ContactEmail)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetContactEmailForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetContactEmailForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetContactEmailForEvent`,
		data,
		nil,
		nil,
	)
}

type SetDefaultSitePageForEventParameters struct {
	EventId    string
	SitePageId string
}

func (t *Event) SetDefaultSitePageForEvent(p *SetDefaultSitePageForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`sitePageId`, p.SitePageId)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetDefaultSitePageForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetDefaultSitePageForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetDefaultSitePageForEvent`,
		data,
		nil,
		nil,
	)
}

type SetDescriptionForEventParameters struct {
	EventId     string
	Description string
}

func (t *Event) SetDescriptionForEvent(p *SetDescriptionForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`description`, p.Description)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetDescriptionForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetDescriptionForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetDescriptionForEvent`,
		data,
		nil,
		nil,
	)
}

type SetFacebookHandleForEventParameters struct {
	EventId        string
	FacebookHandle string
}

func (t *Event) SetFacebookHandleForEvent(p *SetFacebookHandleForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`facebookHandle`, p.FacebookHandle)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetFacebookHandleForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetFacebookHandleForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetFacebookHandleForEvent`,
		data,
		nil,
		nil,
	)
}

type SetInstagramHandleForEventParameters struct {
	EventId         string
	InstagramHandle string
}

func (t *Event) SetInstagramHandleForEvent(p *SetInstagramHandleForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`instagramHandle`, p.InstagramHandle)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetInstagramHandleForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetInstagramHandleForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetInstagramHandleForEvent`,
		data,
		nil,
		nil,
	)
}

type SetLanguageForEventParameters struct {
	EventId  string
	Language string
}

func (t *Event) SetLanguageForEvent(p *SetLanguageForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`language`, p.Language)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetLanguageForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetLanguageForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetLanguageForEvent`,
		data,
		nil,
		nil,
	)
}

type SetMapSourceForEventParameters struct {
	EventId   string
	MapSource string
}

func (t *Event) SetMapSourceForEvent(p *SetMapSourceForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`mapSource`, p.MapSource)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetMapSourceForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetMapSourceForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetMapSourceForEvent`,
		data,
		nil,
		nil,
	)
}

type SetMaxInvitationCountForEventParameters struct {
	EventId            string
	MaxInvitationCount int64
}

func (t *Event) SetMaxInvitationCountForEvent(p *SetMaxInvitationCountForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`maxInvitationCount`, strconv.FormatInt(p.MaxInvitationCount, 10))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetMaxInvitationCountForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetMaxInvitationCountForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetMaxInvitationCountForEvent`,
		data,
		nil,
		nil,
	)
}

type SetMessageForEventParameters struct {
	EventId     string
	MessageType string
	Message     string
}

func (t *Event) SetMessageForEvent(p *SetMessageForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`messageType`, p.MessageType)
	queryParameters.Add(`message`, p.Message)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetMessageForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetMessageForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetMessageForEvent`,
		data,
		nil,
		nil,
	)
}

type SetMinInvitationCountForEventParameters struct {
	EventId            string
	MinInvitationCount int64
}

func (t *Event) SetMinInvitationCountForEvent(p *SetMinInvitationCountForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`minInvitationCount`, strconv.FormatInt(p.MinInvitationCount, 10))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetMinInvitationCountForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetMinInvitationCountForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetMinInvitationCountForEvent`,
		data,
		nil,
		nil,
	)
}

type SetNameForEventParameters struct {
	EventId string
	Name    string
}

func (t *Event) SetNameForEvent(p *SetNameForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`name`, p.Name)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetNameForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetNameForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetNameForEvent`,
		data,
		nil,
		nil,
	)
}

type SetQuestionSortOrderParameters struct {
	QuestionId string
	SortOrder  int64
}

func (t *Event) SetQuestionSortOrder(p *SetQuestionSortOrderParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`sortOrder`, strconv.FormatInt(p.SortOrder, 10))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetQuestionSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetQuestionSortOrderWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetQuestionSortOrder`,
		data,
		nil,
		nil,
	)
}

type SetTimeForEventParameters struct {
	EventId   string
	StartTime string
	EndTime   string
	Timezone  string
}

func (t *Event) SetTimeForEvent(p *SetTimeForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`startTime`, p.StartTime)
	queryParameters.Add(`endTime`, p.EndTime)
	queryParameters.Add(`timezone`, p.Timezone)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetTimeForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetTimeForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetTimeForEvent`,
		data,
		nil,
		nil,
	)
}

type SetTrackingScriptForEventParameters struct {
	EventId            string
	TrackingScriptType string
	TrackingScript     string
}

func (t *Event) SetTrackingScriptForEvent(p *SetTrackingScriptForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`trackingScriptType`, p.TrackingScriptType)
	queryParameters.Add(`trackingScript`, p.TrackingScript)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetTrackingScriptForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetTrackingScriptForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetTrackingScriptForEvent`,
		data,
		nil,
		nil,
	)
}

type SetTwitterHandleForEventParameters struct {
	EventId       string
	TwitterHandle string
}

func (t *Event) SetTwitterHandleForEvent(p *SetTwitterHandleForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`twitterHandle`, p.TwitterHandle)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetTwitterHandleForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetTwitterHandleForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetTwitterHandleForEvent`,
		data,
		nil,
		nil,
	)
}

type SetVenueForEventParameters struct {
	EventId      string
	VenueName    string
	VenueAddress string
}

func (t *Event) SetVenueForEvent(p *SetVenueForEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`venueName`, p.VenueName)
	queryParameters.Add(`venueAddress`, p.VenueAddress)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetVenueForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) SetVenueForEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/SetVenueForEvent`,
		data,
		nil,
		nil,
	)
}

type UnarchiveEventParameters struct {
	EventId string
}

func (t *Event) UnarchiveEvent(p *UnarchiveEventParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/UnarchiveEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) UnarchiveEventWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/UnarchiveEvent`,
		data,
		nil,
		nil,
	)
}

type UpdateAnswerParameters struct {
	AnswerId  string
	Text      string
	IsDefault *bool
}

func (t *Event) UpdateAnswer(p *UpdateAnswerParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)
	queryParameters.Add(`text`, p.Text)
	if p.IsDefault != nil {
		queryParameters.Add(`isDefault`, strconv.FormatBool(*p.IsDefault))
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/UpdateAnswer`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) UpdateAnswerWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/UpdateAnswer`,
		data,
		nil,
		nil,
	)
}

type UpdateQuestionParameters struct {
	QuestionId           string
	Text                 string
	QuestionType         string
	QuestionContextTypes *[]string // registration | lead
	IsRequired           *bool
	IsIndividual         *bool
	TicketTypeId         *string
}

func (t *Event) UpdateQuestion(p *UpdateQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`text`, p.Text)
	queryParameters.Add(`questionType`, p.QuestionType)
	if p.QuestionContextTypes != nil {
		for i := range *p.QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*p.QuestionContextTypes)[i])
		}
	}
	if p.IsRequired != nil {
		queryParameters.Add(`isRequired`, strconv.FormatBool(*p.IsRequired))
	}
	if p.IsIndividual != nil {
		queryParameters.Add(`isIndividual`, strconv.FormatBool(*p.IsIndividual))
	}
	if p.TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *p.TicketTypeId)
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/UpdateQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

func (t *Event) UpdateQuestionWithJSON(data *map[string]interface{}) (r *http.Response, err error) {
	return t.restClient.PostJSON(
		`/v2/Event/UseCase/UpdateQuestion`,
		data,
		nil,
		nil,
	)
}
