package usecase

import (
	"bitbucket.ef.network/go/sdk"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Disable unused import error
var _ = strconv.IntSize
var _ url.Error
var _ = fmt.Sprint("Foo")
var _ = http.NoBody

type Event struct {
	restClient sdk.RestClientInterface
}

func NewEvent(restClient sdk.RestClientInterface) *Event {
	return &Event{restClient}
}

// GET: Queries
// @param string AltKeyword

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

// @param string EventId
// @param array|null WithData Answers|TicketType|QuestionContexts

type GetAllQuestionsForEventParameters struct {
	EventId  string
	WithData *[]string
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

// @param string EventId
// @param array|null WithData Pool|Stacks|StacksWithAvailabilityCounts|Tags|EventTexts|TicketTypes|TicketBlocks|TicketBlocksWithAllotmentCounts|QuestionsAndAnswers|QuestionContext|AllQuestions

type GetEventParameters struct {
	EventId  string
	WithData *[]string
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

// @param string PoolId
// @param int|null EarliestStartDate

type GetEventCountsForPoolParameters struct {
	PoolId            string
	EarliestStartDate *int
}

func (t *Event) GetEventCountsForPool(p *GetEventCountsForPoolParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, p.PoolId)
	if p.EarliestStartDate != nil {
		queryParameters.Add(`earliestStartDate`, strconv.Itoa(*p.EarliestStartDate))
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/GetEventCountsForPool`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string QuestionId
// @param array|null WithData Answers|TicketType|QuestionResponseCounts|AnswerQuestionResponseCounts|QuestionContexts

type GetQuestionParameters struct {
	QuestionId string
	WithData   *[]string
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

// @param string ParentEventId
// @param string|null Query
// @param array|null WithData Pool|Stacks|Tags|TicketTypes|TicketBlocks|QuestionsAndAnswers|ThumbnailUrl
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
// @param string|null SortBy event-start|event-end|name|event-created
// @param string|null SortDirection ascending|descending
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null PoolId
// @param array|null Tags

type ListChildrenForEventParameters struct {
	ParentEventId       string
	Query               *string
	WithData            *[]string
	Page                *int
	ItemsPerPage        *int
	SortBy              *string
	SortDirection       *string
	EventDateFilterType *string
	PoolId              *string
	Tags                *[]string
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
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
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
		`/v2/Event/UseCase/ListChildrenForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string ParentEventId
// @param string UserId
// @param string|null Query
// @param array|null WithData Pool|Stacks|Tags|TicketTypes|TicketBlocks|QuestionsAndAnswers|ThumbnailUrl
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
// @param string|null SortBy event-start|event-end|name|event-created
// @param string|null SortDirection ascending|descending
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null PoolId
// @param array|null Tags

type ListChildrenForEventForUserParameters struct {
	ParentEventId       string
	UserId              string
	Query               *string
	WithData            *[]string
	Page                *int
	ItemsPerPage        *int
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
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
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

// @param string UserId
// @param string|null Query
// @param array|null AttributesFilter distribute|donate|fee|editname|reveal|allow-notes|duplicate-emails|navigation|social-media|social-media-bar|map-location|show-description|ipad-purchase|simple-layout|label-print|skip-event-allocate-display|geo-restrict|visa-checkout|archived|guest-can-change-response|efx-enabled|show-calendar
// @param array|null AttributesExcludeFilter distribute|donate|fee|editname|reveal|allow-notes|duplicate-emails|navigation|social-media|social-media-bar|map-location|show-description|ipad-purchase|simple-layout|label-print|skip-event-allocate-display|geo-restrict|visa-checkout|archived|guest-can-change-response|efx-enabled|show-calendar
// @param array|null WithData Pool|Stacks|Tags|TicketTypes|TicketBlocks|QuestionsAndAnswers|ThumbnailUrl
// @param int|null LastModifiedTimestamp
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-500
// @param string|null SortBy event-start|event-end|name|event-created
// @param string|null SortDirection ascending|descending
// @param string|null EventDateFilterType current-future|past-all|past-3-months|past-3-months-and-future|past-6-months
// @param string|null PoolId
// @param array|null Tags
// @param int|null EarliestStartTimestamp >= 0

type ListEventsForUserParameters struct {
	UserId                  string
	Query                   *string
	AttributesFilter        *[]string
	AttributesExcludeFilter *[]string
	WithData                *[]string
	LastModifiedTimestamp   *int
	Page                    *int
	ItemsPerPage            *int
	SortBy                  *string
	SortDirection           *string
	EventDateFilterType     *string
	PoolId                  *string
	Tags                    *[]string
	EarliestStartTimestamp  *int
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
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*p.LastModifiedTimestamp))
	}
	if p.Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*p.Page))
	}
	if p.ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*p.ItemsPerPage))
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
		queryParameters.Add(`earliestStartTimestamp`, strconv.Itoa(*p.EarliestStartTimestamp))
	}

	return t.restClient.Get(
		`/v2/Event/UseCase/ListEventsForUser`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param array|null QuestionContextTypes registration|lead

type ListQuestionsByEventAndContextParameters struct {
	EventId              string
	QuestionContextTypes *[]string
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
// @param string EventId
// @param string ChildEventId

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

// @param string EventId
// @param string ParentEventId

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

// @param string EventId
// @param string Email
// @param string FirstName
// @param string LastName
// @param string EventRole organizer|assistant|support|check-in-staff|read-only
// @param string|null AuthenticatedUserId

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

// @param string EventId

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

// @param string EventId

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

// @param string FromEventId
// @param string ToEventId
// @param bool|null ShouldCopyUsers true|false
// @param bool|null ShouldCopyEmailDesigns true|false
// @param bool|null ShouldCopySettings true|false
// @param bool|null ShouldCopyRegQuestions true|false
// @param bool|null ShouldCopySitePages true|false
// @param bool|null ShouldCopyTicketTypes true|false
// @param bool|null ShouldCopyTicketBlocks true|false
// @param bool|null ShouldCopyStacks true|false
// @param bool|null ShouldCopyIntegrationSettings true|false

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

// @param string QuestionId
// @param string Text
// @param int|null SortOrder
// @param bool|null IsDefault true|false
// @param string|null AnswerId

type CreateAnswerParameters struct {
	QuestionId string
	Text       string
	SortOrder  *int
	IsDefault  *bool
	AnswerId   *string
}

func (t *Event) CreateAnswer(p *CreateAnswerParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`text`, p.Text)
	if p.SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.Itoa(*p.SortOrder))
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

// @param string PoolId
// @param string UserId
// @param string EventName
// @param string|null StartTime
// @param string|null EndTime
// @param string|null Timezone Africa/Abidjan|Africa/Accra|Africa/Addis_Ababa|Africa/Algiers|Africa/Asmara|Africa/Bamako|Africa/Bangui|Africa/Banjul|Africa/Bissau|Africa/Blantyre|Africa/Brazzaville|Africa/Bujumbura|Africa/Cairo|Africa/Casablanca|Africa/Ceuta|Africa/Conakry|Africa/Dakar|Africa/Dar_es_Salaam|Africa/Djibouti|Africa/Douala|Africa/El_Aaiun|Africa/Freetown|Africa/Gaborone|Africa/Harare|Africa/Johannesburg|Africa/Juba|Africa/Kampala|Africa/Khartoum|Africa/Kigali|Africa/Kinshasa|Africa/Lagos|Africa/Libreville|Africa/Lome|Africa/Luanda|Africa/Lubumbashi|Africa/Lusaka|Africa/Malabo|Africa/Maputo|Africa/Maseru|Africa/Mbabane|Africa/Mogadishu|Africa/Monrovia|Africa/Nairobi|Africa/Ndjamena|Africa/Niamey|Africa/Nouakchott|Africa/Ouagadougou|Africa/Porto-Novo|Africa/Sao_Tome|Africa/Tripoli|Africa/Tunis|Africa/Windhoek|America/Adak|America/Anchorage|America/Anguilla|America/Antigua|America/Araguaina|America/Argentina/Buenos_Aires|America/Argentina/Catamarca|America/Argentina/Cordoba|America/Argentina/Jujuy|America/Argentina/La_Rioja|America/Argentina/Mendoza|America/Argentina/Rio_Gallegos|America/Argentina/Salta|America/Argentina/San_Juan|America/Argentina/San_Luis|America/Argentina/Tucuman|America/Argentina/Ushuaia|America/Aruba|America/Asuncion|America/Atikokan|America/Bahia|America/Bahia_Banderas|America/Barbados|America/Belem|America/Belize|America/Blanc-Sablon|America/Boa_Vista|America/Bogota|America/Boise|America/Cambridge_Bay|America/Campo_Grande|America/Cancun|America/Caracas|America/Cayenne|America/Cayman|America/Chicago|America/Chihuahua|America/Costa_Rica|America/Creston|America/Cuiaba|America/Curacao|America/Danmarkshavn|America/Dawson|America/Dawson_Creek|America/Denver|America/Detroit|America/Dominica|America/Edmonton|America/Eirunepe|America/El_Salvador|America/Fort_Nelson|America/Fortaleza|America/Glace_Bay|America/Godthab|America/Goose_Bay|America/Grand_Turk|America/Grenada|America/Guadeloupe|America/Guatemala|America/Guayaquil|America/Guyana|America/Halifax|America/Havana|America/Hermosillo|America/Indiana/Indianapolis|America/Indiana/Knox|America/Indiana/Marengo|America/Indiana/Petersburg|America/Indiana/Tell_City|America/Indiana/Vevay|America/Indiana/Vincennes|America/Indiana/Winamac|America/Inuvik|America/Iqaluit|America/Jamaica|America/Juneau|America/Kentucky/Louisville|America/Kentucky/Monticello|America/Kralendijk|America/La_Paz|America/Lima|America/Los_Angeles|America/Lower_Princes|America/Maceio|America/Managua|America/Manaus|America/Marigot|America/Martinique|America/Matamoros|America/Mazatlan|America/Menominee|America/Merida|America/Metlakatla|America/Mexico_City|America/Miquelon|America/Moncton|America/Monterrey|America/Montevideo|America/Montserrat|America/Nassau|America/New_York|America/Nipigon|America/Nome|America/Noronha|America/North_Dakota/Beulah|America/North_Dakota/Center|America/North_Dakota/New_Salem|America/Ojinaga|America/Panama|America/Pangnirtung|America/Paramaribo|America/Phoenix|America/Port-au-Prince|America/Port_of_Spain|America/Porto_Velho|America/Puerto_Rico|America/Punta_Arenas|America/Rainy_River|America/Rankin_Inlet|America/Recife|America/Regina|America/Resolute|America/Rio_Branco|America/Santarem|America/Santiago|America/Santo_Domingo|America/Sao_Paulo|America/Scoresbysund|America/Sitka|America/St_Barthelemy|America/St_Johns|America/St_Kitts|America/St_Lucia|America/St_Thomas|America/St_Vincent|America/Swift_Current|America/Tegucigalpa|America/Thule|America/Thunder_Bay|America/Tijuana|America/Toronto|America/Tortola|America/Vancouver|America/Whitehorse|America/Winnipeg|America/Yakutat|America/Yellowknife|Antarctica/Casey|Antarctica/Davis|Antarctica/DumontDUrville|Antarctica/Macquarie|Antarctica/Mawson|Antarctica/McMurdo|Antarctica/Palmer|Antarctica/Rothera|Antarctica/Syowa|Antarctica/Troll|Antarctica/Vostok|Arctic/Longyearbyen|Asia/Aden|Asia/Almaty|Asia/Amman|Asia/Anadyr|Asia/Aqtau|Asia/Aqtobe|Asia/Ashgabat|Asia/Atyrau|Asia/Baghdad|Asia/Bahrain|Asia/Baku|Asia/Bangkok|Asia/Barnaul|Asia/Beirut|Asia/Bishkek|Asia/Brunei|Asia/Chita|Asia/Choibalsan|Asia/Colombo|Asia/Damascus|Asia/Dhaka|Asia/Dili|Asia/Dubai|Asia/Dushanbe|Asia/Famagusta|Asia/Gaza|Asia/Hebron|Asia/Ho_Chi_Minh|Asia/Hong_Kong|Asia/Hovd|Asia/Irkutsk|Asia/Jakarta|Asia/Jayapura|Asia/Jerusalem|Asia/Kabul|Asia/Kamchatka|Asia/Karachi|Asia/Kathmandu|Asia/Khandyga|Asia/Kolkata|Asia/Krasnoyarsk|Asia/Kuala_Lumpur|Asia/Kuching|Asia/Kuwait|Asia/Macau|Asia/Magadan|Asia/Makassar|Asia/Manila|Asia/Muscat|Asia/Nicosia|Asia/Novokuznetsk|Asia/Novosibirsk|Asia/Omsk|Asia/Oral|Asia/Phnom_Penh|Asia/Pontianak|Asia/Pyongyang|Asia/Qatar|Asia/Qyzylorda|Asia/Riyadh|Asia/Sakhalin|Asia/Samarkand|Asia/Seoul|Asia/Shanghai|Asia/Singapore|Asia/Srednekolymsk|Asia/Taipei|Asia/Tashkent|Asia/Tbilisi|Asia/Tehran|Asia/Thimphu|Asia/Tokyo|Asia/Tomsk|Asia/Ulaanbaatar|Asia/Urumqi|Asia/Ust-Nera|Asia/Vientiane|Asia/Vladivostok|Asia/Yakutsk|Asia/Yangon|Asia/Yekaterinburg|Asia/Yerevan|Atlantic/Azores|Atlantic/Bermuda|Atlantic/Canary|Atlantic/Cape_Verde|Atlantic/Faroe|Atlantic/Madeira|Atlantic/Reykjavik|Atlantic/South_Georgia|Atlantic/St_Helena|Atlantic/Stanley|Australia/Adelaide|Australia/Brisbane|Australia/Broken_Hill|Australia/Currie|Australia/Darwin|Australia/Eucla|Australia/Hobart|Australia/Lindeman|Australia/Lord_Howe|Australia/Melbourne|Australia/Perth|Australia/Sydney|Europe/Amsterdam|Europe/Andorra|Europe/Astrakhan|Europe/Athens|Europe/Belgrade|Europe/Berlin|Europe/Bratislava|Europe/Brussels|Europe/Bucharest|Europe/Budapest|Europe/Busingen|Europe/Chisinau|Europe/Copenhagen|Europe/Dublin|Europe/Gibraltar|Europe/Guernsey|Europe/Helsinki|Europe/Isle_of_Man|Europe/Istanbul|Europe/Jersey|Europe/Kaliningrad|Europe/Kiev|Europe/Kirov|Europe/Lisbon|Europe/Ljubljana|Europe/London|Europe/Luxembourg|Europe/Madrid|Europe/Malta|Europe/Mariehamn|Europe/Minsk|Europe/Monaco|Europe/Moscow|Europe/Oslo|Europe/Paris|Europe/Podgorica|Europe/Prague|Europe/Riga|Europe/Rome|Europe/Samara|Europe/San_Marino|Europe/Sarajevo|Europe/Saratov|Europe/Simferopol|Europe/Skopje|Europe/Sofia|Europe/Stockholm|Europe/Tallinn|Europe/Tirane|Europe/Ulyanovsk|Europe/Uzhgorod|Europe/Vaduz|Europe/Vatican|Europe/Vienna|Europe/Vilnius|Europe/Volgograd|Europe/Warsaw|Europe/Zagreb|Europe/Zaporozhye|Europe/Zurich|Indian/Antananarivo|Indian/Chagos|Indian/Christmas|Indian/Cocos|Indian/Comoro|Indian/Kerguelen|Indian/Mahe|Indian/Maldives|Indian/Mauritius|Indian/Mayotte|Indian/Reunion|Pacific/Apia|Pacific/Auckland|Pacific/Bougainville|Pacific/Chatham|Pacific/Chuuk|Pacific/Easter|Pacific/Efate|Pacific/Enderbury|Pacific/Fakaofo|Pacific/Fiji|Pacific/Funafuti|Pacific/Galapagos|Pacific/Gambier|Pacific/Guadalcanal|Pacific/Guam|Pacific/Honolulu|Pacific/Kiritimati|Pacific/Kosrae|Pacific/Kwajalein|Pacific/Majuro|Pacific/Marquesas|Pacific/Midway|Pacific/Nauru|Pacific/Niue|Pacific/Norfolk|Pacific/Noumea|Pacific/Pago_Pago|Pacific/Palau|Pacific/Pitcairn|Pacific/Pohnpei|Pacific/Port_Moresby|Pacific/Rarotonga|Pacific/Saipan|Pacific/Tahiti|Pacific/Tarawa|Pacific/Tongatapu|Pacific/Wake|Pacific/Wallis|UTC
// @param string|null EventId

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

// @param string PoolId
// @param string UserId
// @param string EventName
// @param string|null AltKeyword
// @param string|null ContactEmail
// @param string|null StartTime
// @param string|null EndTime
// @param string|null Timezone Africa/Abidjan|Africa/Accra|Africa/Addis_Ababa|Africa/Algiers|Africa/Asmara|Africa/Bamako|Africa/Bangui|Africa/Banjul|Africa/Bissau|Africa/Blantyre|Africa/Brazzaville|Africa/Bujumbura|Africa/Cairo|Africa/Casablanca|Africa/Ceuta|Africa/Conakry|Africa/Dakar|Africa/Dar_es_Salaam|Africa/Djibouti|Africa/Douala|Africa/El_Aaiun|Africa/Freetown|Africa/Gaborone|Africa/Harare|Africa/Johannesburg|Africa/Juba|Africa/Kampala|Africa/Khartoum|Africa/Kigali|Africa/Kinshasa|Africa/Lagos|Africa/Libreville|Africa/Lome|Africa/Luanda|Africa/Lubumbashi|Africa/Lusaka|Africa/Malabo|Africa/Maputo|Africa/Maseru|Africa/Mbabane|Africa/Mogadishu|Africa/Monrovia|Africa/Nairobi|Africa/Ndjamena|Africa/Niamey|Africa/Nouakchott|Africa/Ouagadougou|Africa/Porto-Novo|Africa/Sao_Tome|Africa/Tripoli|Africa/Tunis|Africa/Windhoek|America/Adak|America/Anchorage|America/Anguilla|America/Antigua|America/Araguaina|America/Argentina/Buenos_Aires|America/Argentina/Catamarca|America/Argentina/Cordoba|America/Argentina/Jujuy|America/Argentina/La_Rioja|America/Argentina/Mendoza|America/Argentina/Rio_Gallegos|America/Argentina/Salta|America/Argentina/San_Juan|America/Argentina/San_Luis|America/Argentina/Tucuman|America/Argentina/Ushuaia|America/Aruba|America/Asuncion|America/Atikokan|America/Bahia|America/Bahia_Banderas|America/Barbados|America/Belem|America/Belize|America/Blanc-Sablon|America/Boa_Vista|America/Bogota|America/Boise|America/Cambridge_Bay|America/Campo_Grande|America/Cancun|America/Caracas|America/Cayenne|America/Cayman|America/Chicago|America/Chihuahua|America/Costa_Rica|America/Creston|America/Cuiaba|America/Curacao|America/Danmarkshavn|America/Dawson|America/Dawson_Creek|America/Denver|America/Detroit|America/Dominica|America/Edmonton|America/Eirunepe|America/El_Salvador|America/Fort_Nelson|America/Fortaleza|America/Glace_Bay|America/Godthab|America/Goose_Bay|America/Grand_Turk|America/Grenada|America/Guadeloupe|America/Guatemala|America/Guayaquil|America/Guyana|America/Halifax|America/Havana|America/Hermosillo|America/Indiana/Indianapolis|America/Indiana/Knox|America/Indiana/Marengo|America/Indiana/Petersburg|America/Indiana/Tell_City|America/Indiana/Vevay|America/Indiana/Vincennes|America/Indiana/Winamac|America/Inuvik|America/Iqaluit|America/Jamaica|America/Juneau|America/Kentucky/Louisville|America/Kentucky/Monticello|America/Kralendijk|America/La_Paz|America/Lima|America/Los_Angeles|America/Lower_Princes|America/Maceio|America/Managua|America/Manaus|America/Marigot|America/Martinique|America/Matamoros|America/Mazatlan|America/Menominee|America/Merida|America/Metlakatla|America/Mexico_City|America/Miquelon|America/Moncton|America/Monterrey|America/Montevideo|America/Montserrat|America/Nassau|America/New_York|America/Nipigon|America/Nome|America/Noronha|America/North_Dakota/Beulah|America/North_Dakota/Center|America/North_Dakota/New_Salem|America/Ojinaga|America/Panama|America/Pangnirtung|America/Paramaribo|America/Phoenix|America/Port-au-Prince|America/Port_of_Spain|America/Porto_Velho|America/Puerto_Rico|America/Punta_Arenas|America/Rainy_River|America/Rankin_Inlet|America/Recife|America/Regina|America/Resolute|America/Rio_Branco|America/Santarem|America/Santiago|America/Santo_Domingo|America/Sao_Paulo|America/Scoresbysund|America/Sitka|America/St_Barthelemy|America/St_Johns|America/St_Kitts|America/St_Lucia|America/St_Thomas|America/St_Vincent|America/Swift_Current|America/Tegucigalpa|America/Thule|America/Thunder_Bay|America/Tijuana|America/Toronto|America/Tortola|America/Vancouver|America/Whitehorse|America/Winnipeg|America/Yakutat|America/Yellowknife|Antarctica/Casey|Antarctica/Davis|Antarctica/DumontDUrville|Antarctica/Macquarie|Antarctica/Mawson|Antarctica/McMurdo|Antarctica/Palmer|Antarctica/Rothera|Antarctica/Syowa|Antarctica/Troll|Antarctica/Vostok|Arctic/Longyearbyen|Asia/Aden|Asia/Almaty|Asia/Amman|Asia/Anadyr|Asia/Aqtau|Asia/Aqtobe|Asia/Ashgabat|Asia/Atyrau|Asia/Baghdad|Asia/Bahrain|Asia/Baku|Asia/Bangkok|Asia/Barnaul|Asia/Beirut|Asia/Bishkek|Asia/Brunei|Asia/Chita|Asia/Choibalsan|Asia/Colombo|Asia/Damascus|Asia/Dhaka|Asia/Dili|Asia/Dubai|Asia/Dushanbe|Asia/Famagusta|Asia/Gaza|Asia/Hebron|Asia/Ho_Chi_Minh|Asia/Hong_Kong|Asia/Hovd|Asia/Irkutsk|Asia/Jakarta|Asia/Jayapura|Asia/Jerusalem|Asia/Kabul|Asia/Kamchatka|Asia/Karachi|Asia/Kathmandu|Asia/Khandyga|Asia/Kolkata|Asia/Krasnoyarsk|Asia/Kuala_Lumpur|Asia/Kuching|Asia/Kuwait|Asia/Macau|Asia/Magadan|Asia/Makassar|Asia/Manila|Asia/Muscat|Asia/Nicosia|Asia/Novokuznetsk|Asia/Novosibirsk|Asia/Omsk|Asia/Oral|Asia/Phnom_Penh|Asia/Pontianak|Asia/Pyongyang|Asia/Qatar|Asia/Qyzylorda|Asia/Riyadh|Asia/Sakhalin|Asia/Samarkand|Asia/Seoul|Asia/Shanghai|Asia/Singapore|Asia/Srednekolymsk|Asia/Taipei|Asia/Tashkent|Asia/Tbilisi|Asia/Tehran|Asia/Thimphu|Asia/Tokyo|Asia/Tomsk|Asia/Ulaanbaatar|Asia/Urumqi|Asia/Ust-Nera|Asia/Vientiane|Asia/Vladivostok|Asia/Yakutsk|Asia/Yangon|Asia/Yekaterinburg|Asia/Yerevan|Atlantic/Azores|Atlantic/Bermuda|Atlantic/Canary|Atlantic/Cape_Verde|Atlantic/Faroe|Atlantic/Madeira|Atlantic/Reykjavik|Atlantic/South_Georgia|Atlantic/St_Helena|Atlantic/Stanley|Australia/Adelaide|Australia/Brisbane|Australia/Broken_Hill|Australia/Currie|Australia/Darwin|Australia/Eucla|Australia/Hobart|Australia/Lindeman|Australia/Lord_Howe|Australia/Melbourne|Australia/Perth|Australia/Sydney|Europe/Amsterdam|Europe/Andorra|Europe/Astrakhan|Europe/Athens|Europe/Belgrade|Europe/Berlin|Europe/Bratislava|Europe/Brussels|Europe/Bucharest|Europe/Budapest|Europe/Busingen|Europe/Chisinau|Europe/Copenhagen|Europe/Dublin|Europe/Gibraltar|Europe/Guernsey|Europe/Helsinki|Europe/Isle_of_Man|Europe/Istanbul|Europe/Jersey|Europe/Kaliningrad|Europe/Kiev|Europe/Kirov|Europe/Lisbon|Europe/Ljubljana|Europe/London|Europe/Luxembourg|Europe/Madrid|Europe/Malta|Europe/Mariehamn|Europe/Minsk|Europe/Monaco|Europe/Moscow|Europe/Oslo|Europe/Paris|Europe/Podgorica|Europe/Prague|Europe/Riga|Europe/Rome|Europe/Samara|Europe/San_Marino|Europe/Sarajevo|Europe/Saratov|Europe/Simferopol|Europe/Skopje|Europe/Sofia|Europe/Stockholm|Europe/Tallinn|Europe/Tirane|Europe/Ulyanovsk|Europe/Uzhgorod|Europe/Vaduz|Europe/Vatican|Europe/Vienna|Europe/Vilnius|Europe/Volgograd|Europe/Warsaw|Europe/Zagreb|Europe/Zaporozhye|Europe/Zurich|Indian/Antananarivo|Indian/Chagos|Indian/Christmas|Indian/Cocos|Indian/Comoro|Indian/Kerguelen|Indian/Mahe|Indian/Maldives|Indian/Mauritius|Indian/Mayotte|Indian/Reunion|Pacific/Apia|Pacific/Auckland|Pacific/Bougainville|Pacific/Chatham|Pacific/Chuuk|Pacific/Easter|Pacific/Efate|Pacific/Enderbury|Pacific/Fakaofo|Pacific/Fiji|Pacific/Funafuti|Pacific/Galapagos|Pacific/Gambier|Pacific/Guadalcanal|Pacific/Guam|Pacific/Honolulu|Pacific/Kiritimati|Pacific/Kosrae|Pacific/Kwajalein|Pacific/Majuro|Pacific/Marquesas|Pacific/Midway|Pacific/Nauru|Pacific/Niue|Pacific/Norfolk|Pacific/Noumea|Pacific/Pago_Pago|Pacific/Palau|Pacific/Pitcairn|Pacific/Pohnpei|Pacific/Port_Moresby|Pacific/Rarotonga|Pacific/Saipan|Pacific/Tahiti|Pacific/Tarawa|Pacific/Tongatapu|Pacific/Wake|Pacific/Wallis|UTC
// @param string|null EventId

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

// @param string EventId
// @param string Text
// @param string QuestionType checkbox|radio|multi|text|select|date|waiver
// @param int|null SortOrder
// @param bool|null IsRequired true|false
// @param bool|null IsIndividual true|false
// @param string|null TicketTypeId
// @param string|null QuestionId
// @param array|null QuestionContextTypes registration|lead

type CreateQuestionParameters struct {
	EventId              string
	Text                 string
	QuestionType         string
	SortOrder            *int
	IsRequired           *bool
	IsIndividual         *bool
	TicketTypeId         *string
	QuestionId           *string
	QuestionContextTypes *[]string
}

func (t *Event) CreateQuestion(p *CreateQuestionParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, p.EventId)
	queryParameters.Add(`text`, p.Text)
	queryParameters.Add(`questionType`, p.QuestionType)
	if p.SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.Itoa(*p.SortOrder))
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

// @param string AnswerId

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

// @param string QuestionId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string QuestionId

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

// @param string EventId
// @param string Field address|company|phone|title|country|position

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string EventId

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

// @param string QuestionId

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

// @param string EventId
// @param string Field address|company|phone|title|country|position

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

// @param string EventId

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

// @param string PoolId
// @param string UserId
// @param string FirstName
// @param string LastName
// @param string Email
// @param string Company
// @param string PhoneNumber
// @param string HowManyEvents fewer-than-5|between-5-and-30|more-than-30
// @param string Industry consumer-goods|agency|education|finance|sports-and-entertainment|technology|other
// @param array TypeOfEvents brand-awareness-events|trade-shows|conferences|roadshows|internal-meetings|recruiting-events|fundraising-events

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
	TypeOfEvents  []string
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

// @param string EventId
// @param string ChildEventId

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

// @param string EventId

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

func (t *Event) SendMessageToGuestList() (r *http.Response, err error) {
	queryParameters := url.Values{}

	return t.restClient.Post(
		`/v2/Event/UseCase/SendMessageToGuestList`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string AltKeyword

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

// @param string AnswerId
// @param int SortOrder

type SetAnswerSortOrderParameters struct {
	AnswerId  string
	SortOrder int
}

func (t *Event) SetAnswerSortOrder(p *SetAnswerSortOrderParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, p.AnswerId)
	queryParameters.Add(`sortOrder`, strconv.Itoa(p.SortOrder))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetAnswerSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string SitePageId

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

// @param string QuestionId
// @param int SortOrder

type SetQuestionSortOrderParameters struct {
	QuestionId string
	SortOrder  int
}

func (t *Event) SetQuestionSortOrder(p *SetQuestionSortOrderParameters) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, p.QuestionId)
	queryParameters.Add(`sortOrder`, strconv.Itoa(p.SortOrder))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetQuestionSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId

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

// @param string AnswerId
// @param string Text
// @param bool|null IsDefault true|false

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

// @param string QuestionId
// @param string Text
// @param string QuestionType checkbox|radio|multi|text|select|date|waiver
// @param array|null QuestionContextTypes registration|lead
// @param bool|null IsRequired true|false
// @param bool|null IsIndividual true|false
// @param string|null TicketTypeId

type UpdateQuestionParameters struct {
	QuestionId           string
	Text                 string
	QuestionType         string
	QuestionContextTypes *[]string
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
