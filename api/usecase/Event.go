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

type Event struct {
	restClient gosdk.RestClientInterface
}

func NewEvent(restClient gosdk.RestClientInterface) *Event {
	return &Event{restClient}
}

// GET: Queries
// @param string AltKeyword
func (t *Event) CheckAltKeywordAvailability(AltKeyword string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`altKeyword`, AltKeyword)

	return t.restClient.Get(
		`/v2/Event/UseCase/CheckAltKeywordAvailability`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param array|null WithData Answers|TicketType|QuestionContexts
func (t *Event) GetAllQuestionsForEvent(EventId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *Event) GetEvent(EventId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *Event) GetEventCountsForPool(PoolId string, EarliestStartDate *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	if EarliestStartDate != nil {
		queryParameters.Add(`earliestStartDate`, strconv.Itoa(*EarliestStartDate))
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
func (t *Event) GetQuestion(QuestionId string, WithData *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, QuestionId)
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
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
func (t *Event) ListChildrenForEvent(ParentEventId string, Query *string, WithData *[]string, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string, EventDateFilterType *string, PoolId *string, Tags *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`parentEventId`, ParentEventId)
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
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
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *EventDateFilterType)
	}
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if Tags != nil {
		for i := range *Tags {
			queryParameters.Add(`tags[]`, (*Tags)[i])
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
func (t *Event) ListChildrenForEventForUser(ParentEventId string, UserId string, Query *string, WithData *[]string, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string, EventDateFilterType *string, PoolId *string, Tags *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`parentEventId`, ParentEventId)
	queryParameters.Add(`userId`, UserId)
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
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
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *EventDateFilterType)
	}
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if Tags != nil {
		for i := range *Tags {
			queryParameters.Add(`tags[]`, (*Tags)[i])
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
// @param array|null AttributesFilter distribute|donate|fee|editname|reveal|allow-notes|duplicate-emails|navigation|social-media|social-media-bar|map-location|show-description|ipad-purchase|simple-layout|label-print|skip-event-allocate-display|geo-restrict|visa-checkout|archived|guest-can-change-response|efx-enabled
// @param array|null AttributesExcludeFilter distribute|donate|fee|editname|reveal|allow-notes|duplicate-emails|navigation|social-media|social-media-bar|map-location|show-description|ipad-purchase|simple-layout|label-print|skip-event-allocate-display|geo-restrict|visa-checkout|archived|guest-can-change-response|efx-enabled
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
func (t *Event) ListEventsForUser(UserId string, Query *string, AttributesFilter *[]string, AttributesExcludeFilter *[]string, WithData *[]string, LastModifiedTimestamp *int, Page *int, ItemsPerPage *int, SortBy *string, SortDirection *string, EventDateFilterType *string, PoolId *string, Tags *[]string, EarliestStartTimestamp *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`userId`, UserId)
	if Query != nil {
		queryParameters.Add(`query`, *Query)
	}
	if AttributesFilter != nil {
		for i := range *AttributesFilter {
			queryParameters.Add(`attributesFilter[]`, (*AttributesFilter)[i])
		}
	}
	if AttributesExcludeFilter != nil {
		for i := range *AttributesExcludeFilter {
			queryParameters.Add(`attributesExcludeFilter[]`, (*AttributesExcludeFilter)[i])
		}
	}
	if WithData != nil {
		for i := range *WithData {
			queryParameters.Add(`withData[]`, (*WithData)[i])
		}
	}
	if LastModifiedTimestamp != nil {
		queryParameters.Add(`lastModifiedTimestamp`, strconv.Itoa(*LastModifiedTimestamp))
	}
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}
	if SortBy != nil {
		queryParameters.Add(`sortBy`, *SortBy)
	}
	if SortDirection != nil {
		queryParameters.Add(`sortDirection`, *SortDirection)
	}
	if EventDateFilterType != nil {
		queryParameters.Add(`eventDateFilterType`, *EventDateFilterType)
	}
	if PoolId != nil {
		queryParameters.Add(`poolId`, *PoolId)
	}
	if Tags != nil {
		for i := range *Tags {
			queryParameters.Add(`tags[]`, (*Tags)[i])
		}
	}
	if EarliestStartTimestamp != nil {
		queryParameters.Add(`earliestStartTimestamp`, strconv.Itoa(*EarliestStartTimestamp))
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
func (t *Event) ListQuestionsByEventAndContext(EventId string, QuestionContextTypes *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if QuestionContextTypes != nil {
		for i := range *QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*QuestionContextTypes)[i])
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
func (t *Event) AddChildEvent(EventId string, ChildEventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`childEventId`, ChildEventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/AddChildEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string ParentEventId
func (t *Event) AddParentToEvent(EventId string, ParentEventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`parentEventId`, ParentEventId)

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
func (t *Event) AddUserRoleToEvent(EventId string, Email string, FirstName string, LastName string, EventRole string, AuthenticatedUserId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`email`, Email)
	queryParameters.Add(`firstName`, FirstName)
	queryParameters.Add(`lastName`, LastName)
	queryParameters.Add(`eventRole`, EventRole)
	if AuthenticatedUserId != nil {
		queryParameters.Add(`authenticatedUserId`, *AuthenticatedUserId)
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/AddUserRoleToEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) ArchiveEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/ArchiveEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) ClearDefaultSitePageForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

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
func (t *Event) CopyExistingEventConfiguration(FromEventId string, ToEventId string, ShouldCopyUsers *bool, ShouldCopyEmailDesigns *bool, ShouldCopySettings *bool, ShouldCopyRegQuestions *bool, ShouldCopySitePages *bool, ShouldCopyTicketTypes *bool, ShouldCopyTicketBlocks *bool, ShouldCopyStacks *bool, ShouldCopyIntegrationSettings *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`fromEventId`, FromEventId)
	queryParameters.Add(`toEventId`, ToEventId)
	if ShouldCopyUsers != nil {
		queryParameters.Add(`shouldCopyUsers`, strconv.FormatBool(*ShouldCopyUsers))
	}
	if ShouldCopyEmailDesigns != nil {
		queryParameters.Add(`shouldCopyEmailDesigns`, strconv.FormatBool(*ShouldCopyEmailDesigns))
	}
	if ShouldCopySettings != nil {
		queryParameters.Add(`shouldCopySettings`, strconv.FormatBool(*ShouldCopySettings))
	}
	if ShouldCopyRegQuestions != nil {
		queryParameters.Add(`shouldCopyRegQuestions`, strconv.FormatBool(*ShouldCopyRegQuestions))
	}
	if ShouldCopySitePages != nil {
		queryParameters.Add(`shouldCopySitePages`, strconv.FormatBool(*ShouldCopySitePages))
	}
	if ShouldCopyTicketTypes != nil {
		queryParameters.Add(`shouldCopyTicketTypes`, strconv.FormatBool(*ShouldCopyTicketTypes))
	}
	if ShouldCopyTicketBlocks != nil {
		queryParameters.Add(`shouldCopyTicketBlocks`, strconv.FormatBool(*ShouldCopyTicketBlocks))
	}
	if ShouldCopyStacks != nil {
		queryParameters.Add(`shouldCopyStacks`, strconv.FormatBool(*ShouldCopyStacks))
	}
	if ShouldCopyIntegrationSettings != nil {
		queryParameters.Add(`shouldCopyIntegrationSettings`, strconv.FormatBool(*ShouldCopyIntegrationSettings))
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
func (t *Event) CreateAnswer(QuestionId string, Text string, SortOrder *int, IsDefault *bool, AnswerId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, QuestionId)
	queryParameters.Add(`text`, Text)
	if SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.Itoa(*SortOrder))
	}
	if IsDefault != nil {
		queryParameters.Add(`isDefault`, strconv.FormatBool(*IsDefault))
	}
	if AnswerId != nil {
		queryParameters.Add(`answerId`, *AnswerId)
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
func (t *Event) CreateCIOEvent(PoolId string, UserId string, EventName string, StartTime *string, EndTime *string, Timezone *string, EventId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`eventName`, EventName)
	if StartTime != nil {
		queryParameters.Add(`startTime`, *StartTime)
	}
	if EndTime != nil {
		queryParameters.Add(`endTime`, *EndTime)
	}
	if Timezone != nil {
		queryParameters.Add(`timezone`, *Timezone)
	}
	if EventId != nil {
		queryParameters.Add(`eventId`, *EventId)
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
func (t *Event) CreateEvent(PoolId string, UserId string, EventName string, AltKeyword *string, ContactEmail *string, StartTime *string, EndTime *string, Timezone *string, EventId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`eventName`, EventName)
	if AltKeyword != nil {
		queryParameters.Add(`altKeyword`, *AltKeyword)
	}
	if ContactEmail != nil {
		queryParameters.Add(`contactEmail`, *ContactEmail)
	}
	if StartTime != nil {
		queryParameters.Add(`startTime`, *StartTime)
	}
	if EndTime != nil {
		queryParameters.Add(`endTime`, *EndTime)
	}
	if Timezone != nil {
		queryParameters.Add(`timezone`, *Timezone)
	}
	if EventId != nil {
		queryParameters.Add(`eventId`, *EventId)
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
func (t *Event) CreateQuestion(EventId string, Text string, QuestionType string, SortOrder *int, IsRequired *bool, IsIndividual *bool, TicketTypeId *string, QuestionId *string, QuestionContextTypes *[]string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`text`, Text)
	queryParameters.Add(`questionType`, QuestionType)
	if SortOrder != nil {
		queryParameters.Add(`sortOrder`, strconv.Itoa(*SortOrder))
	}
	if IsRequired != nil {
		queryParameters.Add(`isRequired`, strconv.FormatBool(*IsRequired))
	}
	if IsIndividual != nil {
		queryParameters.Add(`isIndividual`, strconv.FormatBool(*IsIndividual))
	}
	if TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *TicketTypeId)
	}
	if QuestionId != nil {
		queryParameters.Add(`questionId`, *QuestionId)
	}
	if QuestionContextTypes != nil {
		for i := range *QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*QuestionContextTypes)[i])
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
func (t *Event) DeleteAnswer(AnswerId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, AnswerId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DeleteAnswer`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string QuestionId
func (t *Event) DeleteQuestion(QuestionId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, QuestionId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DeleteQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableAmexCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableAmexCard`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableCanvas(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableCanvas`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableDiscoverCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDiscoverCard`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableDistribution(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDistribution`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableDonation(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDonation`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableDuplicateEmails(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableDuplicateEmails`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableEditName(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableEditName`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableEfx(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableEfx`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableGuestCanChangeResponse(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableGuestCanChangeResponse`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableInvitationReveal(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableInvitationReveal`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableMastercardCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableMastercardCard`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string QuestionId
func (t *Event) DisableQuestion(QuestionId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, QuestionId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string Field address|company|phone|title|country|position
func (t *Event) DisableUserField(EventId string, Field string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`field`, Field)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableUserField`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) DisableVisaCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/DisableVisaCard`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableAmexCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableAmexCard`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableCanvas(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableCanvas`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableDiscoverCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDiscoverCard`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableDistribution(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDistribution`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableDonation(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDonation`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableDuplicateEmails(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableDuplicateEmails`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableEditName(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableEditName`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableEfx(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableEfx`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableGuestCanChangeResponse(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableGuestCanChangeResponse`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableInvitationReveal(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableInvitationReveal`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableMastercardCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableMastercardCard`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string QuestionId
func (t *Event) EnableQuestion(QuestionId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, QuestionId)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string Field address|company|phone|title|country|position
func (t *Event) EnableUserField(EventId string, Field string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`field`, Field)

	return t.restClient.Post(
		`/v2/Event/UseCase/EnableUserField`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) EnableVisaCard(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

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
func (t *Event) IncrementCIOAndMessageCSM(PoolId string, UserId string, FirstName string, LastName string, Email string, Company string, PhoneNumber string, HowManyEvents string, Industry string, TypeOfEvents []string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`poolId`, PoolId)
	queryParameters.Add(`userId`, UserId)
	queryParameters.Add(`firstName`, FirstName)
	queryParameters.Add(`lastName`, LastName)
	queryParameters.Add(`email`, Email)
	queryParameters.Add(`company`, Company)
	queryParameters.Add(`phoneNumber`, PhoneNumber)
	queryParameters.Add(`howManyEvents`, HowManyEvents)
	queryParameters.Add(`industry`, Industry)
	for i := range TypeOfEvents {
		queryParameters.Add(`typeOfEvents[]`, TypeOfEvents[i])
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
func (t *Event) RemoveChildEvent(EventId string, ChildEventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`childEventId`, ChildEventId)

	return t.restClient.Post(
		`/v2/Event/UseCase/RemoveChildEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) RemoveParentFromEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

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
func (t *Event) SetAltKeywordForEvent(EventId string, AltKeyword string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`altKeyword`, AltKeyword)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetAltKeywordForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string AnswerId
// @param int SortOrder
func (t *Event) SetAnswerSortOrder(AnswerId string, SortOrder int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, AnswerId)
	queryParameters.Add(`sortOrder`, strconv.Itoa(SortOrder))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetAnswerSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string SitePageId
func (t *Event) SetDefaultSitePageForEvent(EventId string, SitePageId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`sitePageId`, SitePageId)

	return t.restClient.Post(
		`/v2/Event/UseCase/SetDefaultSitePageForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string QuestionId
// @param int SortOrder
func (t *Event) SetQuestionSortOrder(QuestionId string, SortOrder int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, QuestionId)
	queryParameters.Add(`sortOrder`, strconv.Itoa(SortOrder))

	return t.restClient.Post(
		`/v2/Event/UseCase/SetQuestionSortOrder`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *Event) UnarchiveEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

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
func (t *Event) UpdateAnswer(AnswerId string, Text string, IsDefault *bool) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`answerId`, AnswerId)
	queryParameters.Add(`text`, Text)
	if IsDefault != nil {
		queryParameters.Add(`isDefault`, strconv.FormatBool(*IsDefault))
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
func (t *Event) UpdateQuestion(QuestionId string, Text string, QuestionType string, QuestionContextTypes *[]string, IsRequired *bool, IsIndividual *bool, TicketTypeId *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`questionId`, QuestionId)
	queryParameters.Add(`text`, Text)
	queryParameters.Add(`questionType`, QuestionType)
	if QuestionContextTypes != nil {
		for i := range *QuestionContextTypes {
			queryParameters.Add(`questionContextTypes[]`, (*QuestionContextTypes)[i])
		}
	}
	if IsRequired != nil {
		queryParameters.Add(`isRequired`, strconv.FormatBool(*IsRequired))
	}
	if IsIndividual != nil {
		queryParameters.Add(`isIndividual`, strconv.FormatBool(*IsIndividual))
	}
	if TicketTypeId != nil {
		queryParameters.Add(`ticketTypeId`, *TicketTypeId)
	}

	return t.restClient.Post(
		`/v2/Event/UseCase/UpdateQuestion`,
		&queryParameters,
		nil,
		nil,
	)
}
