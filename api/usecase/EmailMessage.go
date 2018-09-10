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

type EmailMessage struct {
	restClient gosdk.RestClientInterface
}

func NewEmailMessage(restClient gosdk.RestClientInterface) *EmailMessage {
	return &EmailMessage{restClient}
}

// GET: Queries
// @param string EmailMessageId
func (t *EmailMessage) GetEmailMessage(EmailMessageId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`emailMessageId`, EmailMessageId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *EmailMessage) GetEmailMessageStatsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetEmailMessageStatsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string MongoQueueId
func (t *EmailMessage) GetScheduledEmailMessage(MongoQueueId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, MongoQueueId)

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/GetScheduledEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null Page
// @param int|null ItemsPerPage
func (t *EmailMessage) ListOutboxEmailMessagesByEvent(EventId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListOutboxEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null Page >= 1
// @param int|null ItemsPerPage 1-100
func (t *EmailMessage) ListScheduledBatchEmailMessagesByEvent(EventId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListScheduledBatchEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null Page
// @param int|null ItemsPerPage
func (t *EmailMessage) ListScheduledEmailMessagesByEvent(EventId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListScheduledEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param int|null Page
// @param int|null ItemsPerPage
func (t *EmailMessage) ListSentEmailMessagesByEvent(EventId string, Page *int, ItemsPerPage *int) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	if Page != nil {
		queryParameters.Add(`page`, strconv.Itoa(*Page))
	}
	if ItemsPerPage != nil {
		queryParameters.Add(`itemsPerPage`, strconv.Itoa(*ItemsPerPage))
	}

	return t.restClient.Get(
		`/v2/EmailMessage/UseCase/ListSentEmailMessagesByEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// POST: Commands
// @param string EventId
// @param string OwnerUserId
// @param string MessageContent
// @param array ToEmails
// @param string FromName
// @param string Subject
// @param string LayoutType blank|alt-email-layout
// @param string MessageType confirmation|donation|invitation|sendamessage|sendgroupmessage
// @param string BackgroundColor
// @param string|null FromEmail
// @param string|null ReplyEmail
// @param string|null DomainMaskEmail
func (t *EmailMessage) CreatePreviewEmailMessage(EventId string, OwnerUserId string, MessageContent string, ToEmails []string, FromName string, Subject string, LayoutType string, MessageType string, BackgroundColor string, FromEmail *string, ReplyEmail *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`messageContent`, MessageContent)
	for i := range ToEmails {
		queryParameters.Add(`toEmails[]`, ToEmails[i])
	}
	queryParameters.Add(`fromName`, FromName)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`layoutType`, LayoutType)
	queryParameters.Add(`messageType`, MessageType)
	queryParameters.Add(`backgroundColor`, BackgroundColor)
	if FromEmail != nil {
		queryParameters.Add(`fromEmail`, *FromEmail)
	}
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/CreatePreviewEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string EmailDesignId
// @param string Type status
// @param array Targets confirmed|purchased|assigned|unconfirmed|attended
// @param string MessageSendTime
// @param string Timezone Africa/Abidjan|Africa/Accra|Africa/Addis_Ababa|Africa/Algiers|Africa/Asmara|Africa/Bamako|Africa/Bangui|Africa/Banjul|Africa/Bissau|Africa/Blantyre|Africa/Brazzaville|Africa/Bujumbura|Africa/Cairo|Africa/Casablanca|Africa/Ceuta|Africa/Conakry|Africa/Dakar|Africa/Dar_es_Salaam|Africa/Djibouti|Africa/Douala|Africa/El_Aaiun|Africa/Freetown|Africa/Gaborone|Africa/Harare|Africa/Johannesburg|Africa/Juba|Africa/Kampala|Africa/Khartoum|Africa/Kigali|Africa/Kinshasa|Africa/Lagos|Africa/Libreville|Africa/Lome|Africa/Luanda|Africa/Lubumbashi|Africa/Lusaka|Africa/Malabo|Africa/Maputo|Africa/Maseru|Africa/Mbabane|Africa/Mogadishu|Africa/Monrovia|Africa/Nairobi|Africa/Ndjamena|Africa/Niamey|Africa/Nouakchott|Africa/Ouagadougou|Africa/Porto-Novo|Africa/Sao_Tome|Africa/Tripoli|Africa/Tunis|Africa/Windhoek|America/Adak|America/Anchorage|America/Anguilla|America/Antigua|America/Araguaina|America/Argentina/Buenos_Aires|America/Argentina/Catamarca|America/Argentina/Cordoba|America/Argentina/Jujuy|America/Argentina/La_Rioja|America/Argentina/Mendoza|America/Argentina/Rio_Gallegos|America/Argentina/Salta|America/Argentina/San_Juan|America/Argentina/San_Luis|America/Argentina/Tucuman|America/Argentina/Ushuaia|America/Aruba|America/Asuncion|America/Atikokan|America/Bahia|America/Bahia_Banderas|America/Barbados|America/Belem|America/Belize|America/Blanc-Sablon|America/Boa_Vista|America/Bogota|America/Boise|America/Cambridge_Bay|America/Campo_Grande|America/Cancun|America/Caracas|America/Cayenne|America/Cayman|America/Chicago|America/Chihuahua|America/Costa_Rica|America/Creston|America/Cuiaba|America/Curacao|America/Danmarkshavn|America/Dawson|America/Dawson_Creek|America/Denver|America/Detroit|America/Dominica|America/Edmonton|America/Eirunepe|America/El_Salvador|America/Fort_Nelson|America/Fortaleza|America/Glace_Bay|America/Godthab|America/Goose_Bay|America/Grand_Turk|America/Grenada|America/Guadeloupe|America/Guatemala|America/Guayaquil|America/Guyana|America/Halifax|America/Havana|America/Hermosillo|America/Indiana/Indianapolis|America/Indiana/Knox|America/Indiana/Marengo|America/Indiana/Petersburg|America/Indiana/Tell_City|America/Indiana/Vevay|America/Indiana/Vincennes|America/Indiana/Winamac|America/Inuvik|America/Iqaluit|America/Jamaica|America/Juneau|America/Kentucky/Louisville|America/Kentucky/Monticello|America/Kralendijk|America/La_Paz|America/Lima|America/Los_Angeles|America/Lower_Princes|America/Maceio|America/Managua|America/Manaus|America/Marigot|America/Martinique|America/Matamoros|America/Mazatlan|America/Menominee|America/Merida|America/Metlakatla|America/Mexico_City|America/Miquelon|America/Moncton|America/Monterrey|America/Montevideo|America/Montserrat|America/Nassau|America/New_York|America/Nipigon|America/Nome|America/Noronha|America/North_Dakota/Beulah|America/North_Dakota/Center|America/North_Dakota/New_Salem|America/Ojinaga|America/Panama|America/Pangnirtung|America/Paramaribo|America/Phoenix|America/Port-au-Prince|America/Port_of_Spain|America/Porto_Velho|America/Puerto_Rico|America/Punta_Arenas|America/Rainy_River|America/Rankin_Inlet|America/Recife|America/Regina|America/Resolute|America/Rio_Branco|America/Santarem|America/Santiago|America/Santo_Domingo|America/Sao_Paulo|America/Scoresbysund|America/Sitka|America/St_Barthelemy|America/St_Johns|America/St_Kitts|America/St_Lucia|America/St_Thomas|America/St_Vincent|America/Swift_Current|America/Tegucigalpa|America/Thule|America/Thunder_Bay|America/Tijuana|America/Toronto|America/Tortola|America/Vancouver|America/Whitehorse|America/Winnipeg|America/Yakutat|America/Yellowknife|Antarctica/Casey|Antarctica/Davis|Antarctica/DumontDUrville|Antarctica/Macquarie|Antarctica/Mawson|Antarctica/McMurdo|Antarctica/Palmer|Antarctica/Rothera|Antarctica/Syowa|Antarctica/Troll|Antarctica/Vostok|Arctic/Longyearbyen|Asia/Aden|Asia/Almaty|Asia/Amman|Asia/Anadyr|Asia/Aqtau|Asia/Aqtobe|Asia/Ashgabat|Asia/Atyrau|Asia/Baghdad|Asia/Bahrain|Asia/Baku|Asia/Bangkok|Asia/Barnaul|Asia/Beirut|Asia/Bishkek|Asia/Brunei|Asia/Chita|Asia/Choibalsan|Asia/Colombo|Asia/Damascus|Asia/Dhaka|Asia/Dili|Asia/Dubai|Asia/Dushanbe|Asia/Famagusta|Asia/Gaza|Asia/Hebron|Asia/Ho_Chi_Minh|Asia/Hong_Kong|Asia/Hovd|Asia/Irkutsk|Asia/Jakarta|Asia/Jayapura|Asia/Jerusalem|Asia/Kabul|Asia/Kamchatka|Asia/Karachi|Asia/Kathmandu|Asia/Khandyga|Asia/Kolkata|Asia/Krasnoyarsk|Asia/Kuala_Lumpur|Asia/Kuching|Asia/Kuwait|Asia/Macau|Asia/Magadan|Asia/Makassar|Asia/Manila|Asia/Muscat|Asia/Nicosia|Asia/Novokuznetsk|Asia/Novosibirsk|Asia/Omsk|Asia/Oral|Asia/Phnom_Penh|Asia/Pontianak|Asia/Pyongyang|Asia/Qatar|Asia/Qyzylorda|Asia/Riyadh|Asia/Sakhalin|Asia/Samarkand|Asia/Seoul|Asia/Shanghai|Asia/Singapore|Asia/Srednekolymsk|Asia/Taipei|Asia/Tashkent|Asia/Tbilisi|Asia/Tehran|Asia/Thimphu|Asia/Tokyo|Asia/Tomsk|Asia/Ulaanbaatar|Asia/Urumqi|Asia/Ust-Nera|Asia/Vientiane|Asia/Vladivostok|Asia/Yakutsk|Asia/Yangon|Asia/Yekaterinburg|Asia/Yerevan|Atlantic/Azores|Atlantic/Bermuda|Atlantic/Canary|Atlantic/Cape_Verde|Atlantic/Faroe|Atlantic/Madeira|Atlantic/Reykjavik|Atlantic/South_Georgia|Atlantic/St_Helena|Atlantic/Stanley|Australia/Adelaide|Australia/Brisbane|Australia/Broken_Hill|Australia/Currie|Australia/Darwin|Australia/Eucla|Australia/Hobart|Australia/Lindeman|Australia/Lord_Howe|Australia/Melbourne|Australia/Perth|Australia/Sydney|Europe/Amsterdam|Europe/Andorra|Europe/Astrakhan|Europe/Athens|Europe/Belgrade|Europe/Berlin|Europe/Bratislava|Europe/Brussels|Europe/Bucharest|Europe/Budapest|Europe/Busingen|Europe/Chisinau|Europe/Copenhagen|Europe/Dublin|Europe/Gibraltar|Europe/Guernsey|Europe/Helsinki|Europe/Isle_of_Man|Europe/Istanbul|Europe/Jersey|Europe/Kaliningrad|Europe/Kiev|Europe/Kirov|Europe/Lisbon|Europe/Ljubljana|Europe/London|Europe/Luxembourg|Europe/Madrid|Europe/Malta|Europe/Mariehamn|Europe/Minsk|Europe/Monaco|Europe/Moscow|Europe/Oslo|Europe/Paris|Europe/Podgorica|Europe/Prague|Europe/Riga|Europe/Rome|Europe/Samara|Europe/San_Marino|Europe/Sarajevo|Europe/Saratov|Europe/Simferopol|Europe/Skopje|Europe/Sofia|Europe/Stockholm|Europe/Tallinn|Europe/Tirane|Europe/Ulyanovsk|Europe/Uzhgorod|Europe/Vaduz|Europe/Vatican|Europe/Vienna|Europe/Vilnius|Europe/Volgograd|Europe/Warsaw|Europe/Zagreb|Europe/Zaporozhye|Europe/Zurich|Indian/Antananarivo|Indian/Chagos|Indian/Christmas|Indian/Cocos|Indian/Comoro|Indian/Kerguelen|Indian/Mahe|Indian/Maldives|Indian/Mauritius|Indian/Mayotte|Indian/Reunion|Pacific/Apia|Pacific/Auckland|Pacific/Bougainville|Pacific/Chatham|Pacific/Chuuk|Pacific/Easter|Pacific/Efate|Pacific/Enderbury|Pacific/Fakaofo|Pacific/Fiji|Pacific/Funafuti|Pacific/Galapagos|Pacific/Gambier|Pacific/Guadalcanal|Pacific/Guam|Pacific/Honolulu|Pacific/Kiritimati|Pacific/Kosrae|Pacific/Kwajalein|Pacific/Majuro|Pacific/Marquesas|Pacific/Midway|Pacific/Nauru|Pacific/Niue|Pacific/Norfolk|Pacific/Noumea|Pacific/Pago_Pago|Pacific/Palau|Pacific/Pitcairn|Pacific/Pohnpei|Pacific/Port_Moresby|Pacific/Rarotonga|Pacific/Saipan|Pacific/Tahiti|Pacific/Tarawa|Pacific/Tongatapu|Pacific/Wake|Pacific/Wallis|UTC
func (t *EmailMessage) CreateScheduledBatchEmailMessage(EventId string, OwnerUserId string, EmailDesignId string, Type string, Targets []string, MessageSendTime string, Timezone string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`emailDesignId`, EmailDesignId)
	queryParameters.Add(`type`, Type)
	for i := range Targets {
		queryParameters.Add(`targets[]`, Targets[i])
	}
	queryParameters.Add(`messageSendTime`, MessageSendTime)
	queryParameters.Add(`timezone`, Timezone)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/CreateScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string MongoQueueId
func (t *EmailMessage) RemoveScheduledBatchEmailMessage(MongoQueueId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, MongoQueueId)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/RemoveScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string MessageContents
// @param string Subject
// @param string FromName
// @param string|null ReplyEmail
// @param bool|null UseEventSpecificLayout true|false
// @param string|null BackgroundColor
// @param string|null DomainMaskEmail
func (t *EmailMessage) SendAMessagePreview(EventId string, OwnerUserId string, MessageContents string, Subject string, FromName string, ReplyEmail *string, UseEventSpecificLayout *bool, BackgroundColor *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`messageContents`, MessageContents)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`fromName`, FromName)
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*UseEventSpecificLayout))
	}
	if BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *BackgroundColor)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessagePreview`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string MessageContents
// @param string Subject
// @param string FromName
// @param string|null ReplyEmail
// @param array|null CcEmails
// @param array|null BccEmails
// @param bool|null UseEventSpecificLayout true|false
// @param string|null BackgroundColor
// @param string|null DomainMaskEmail
func (t *EmailMessage) SendAMessageToAllCheckedInGuests(EventId string, OwnerUserId string, MessageContents string, Subject string, FromName string, ReplyEmail *string, CcEmails *[]string, BccEmails *[]string, UseEventSpecificLayout *bool, BackgroundColor *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`messageContents`, MessageContents)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`fromName`, FromName)
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmails != nil {
		for i := range *CcEmails {
			queryParameters.Add(`ccEmails[]`, (*CcEmails)[i])
		}
	}
	if BccEmails != nil {
		for i := range *BccEmails {
			queryParameters.Add(`bccEmails[]`, (*BccEmails)[i])
		}
	}
	if UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*UseEventSpecificLayout))
	}
	if BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *BackgroundColor)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToAllCheckedInGuests`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string MessageContents
// @param string Subject
// @param string FromName
// @param string|null ReplyEmail
// @param array|null CcEmails
// @param array|null BccEmails
// @param bool|null UseEventSpecificLayout true|false
// @param string|null BackgroundColor
// @param string|null DomainMaskEmail
func (t *EmailMessage) SendAMessageToEventWaitList(EventId string, OwnerUserId string, MessageContents string, Subject string, FromName string, ReplyEmail *string, CcEmails *[]string, BccEmails *[]string, UseEventSpecificLayout *bool, BackgroundColor *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`messageContents`, MessageContents)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`fromName`, FromName)
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmails != nil {
		for i := range *CcEmails {
			queryParameters.Add(`ccEmails[]`, (*CcEmails)[i])
		}
	}
	if BccEmails != nil {
		for i := range *BccEmails {
			queryParameters.Add(`bccEmails[]`, (*BccEmails)[i])
		}
	}
	if UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*UseEventSpecificLayout))
	}
	if BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *BackgroundColor)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToEventWaitList`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string GroupId
// @param string EventId
// @param string OwnerUserId
// @param string MessageContents
// @param string Subject
// @param string FromName
// @param string|null ReplyEmail
// @param array|null CcEmails
// @param array|null BccEmails
// @param bool|null UseEventSpecificLayout true|false
// @param string|null BackgroundColor
// @param string|null DomainMaskEmail
func (t *EmailMessage) SendAMessageToGroup(GroupId string, EventId string, OwnerUserId string, MessageContents string, Subject string, FromName string, ReplyEmail *string, CcEmails *[]string, BccEmails *[]string, UseEventSpecificLayout *bool, BackgroundColor *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`groupId`, GroupId)
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`messageContents`, MessageContents)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`fromName`, FromName)
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmails != nil {
		for i := range *CcEmails {
			queryParameters.Add(`ccEmails[]`, (*CcEmails)[i])
		}
	}
	if BccEmails != nil {
		for i := range *BccEmails {
			queryParameters.Add(`bccEmails[]`, (*BccEmails)[i])
		}
	}
	if UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*UseEventSpecificLayout))
	}
	if BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *BackgroundColor)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToGroup`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param string OwnerUserId
// @param string MessageContents
// @param string Subject
// @param string FromName
// @param string|null ReplyEmail
// @param array|null CcEmails
// @param array|null BccEmails
// @param bool|null UseEventSpecificLayout true|false
// @param string|null BackgroundColor
// @param string|null DomainMaskEmail
func (t *EmailMessage) SendAMessageToNoShows(EventId string, OwnerUserId string, MessageContents string, Subject string, FromName string, ReplyEmail *string, CcEmails *[]string, BccEmails *[]string, UseEventSpecificLayout *bool, BackgroundColor *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`messageContents`, MessageContents)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`fromName`, FromName)
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmails != nil {
		for i := range *CcEmails {
			queryParameters.Add(`ccEmails[]`, (*CcEmails)[i])
		}
	}
	if BccEmails != nil {
		for i := range *BccEmails {
			queryParameters.Add(`bccEmails[]`, (*BccEmails)[i])
		}
	}
	if UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*UseEventSpecificLayout))
	}
	if BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *BackgroundColor)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToNoShows`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
// @param array TicketTypeIds
// @param array InvitationStatusTypes assigned|purchased|confirmed-by-rsvp|declined-by-rsvp|left-behind|not-yet-purchased|registered|unconfirmed|recycled|not-yet-registered|waitlisted
// @param string OwnerUserId
// @param string MessageContents
// @param string Subject
// @param string FromName
// @param string|null ReplyEmail
// @param array|null CcEmails
// @param array|null BccEmails
// @param bool|null UseEventSpecificLayout true|false
// @param string|null BackgroundColor
// @param string|null DomainMaskEmail
func (t *EmailMessage) SendAMessageToTicketTypes(EventId string, TicketTypeIds []string, InvitationStatusTypes []string, OwnerUserId string, MessageContents string, Subject string, FromName string, ReplyEmail *string, CcEmails *[]string, BccEmails *[]string, UseEventSpecificLayout *bool, BackgroundColor *string, DomainMaskEmail *string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)
	for i := range TicketTypeIds {
		queryParameters.Add(`ticketTypeIds[]`, TicketTypeIds[i])
	}
	for i := range InvitationStatusTypes {
		queryParameters.Add(`invitationStatusTypes[]`, InvitationStatusTypes[i])
	}
	queryParameters.Add(`ownerUserId`, OwnerUserId)
	queryParameters.Add(`messageContents`, MessageContents)
	queryParameters.Add(`subject`, Subject)
	queryParameters.Add(`fromName`, FromName)
	if ReplyEmail != nil {
		queryParameters.Add(`replyEmail`, *ReplyEmail)
	}
	if CcEmails != nil {
		for i := range *CcEmails {
			queryParameters.Add(`ccEmails[]`, (*CcEmails)[i])
		}
	}
	if BccEmails != nil {
		for i := range *BccEmails {
			queryParameters.Add(`bccEmails[]`, (*BccEmails)[i])
		}
	}
	if UseEventSpecificLayout != nil {
		queryParameters.Add(`useEventSpecificLayout`, strconv.FormatBool(*UseEventSpecificLayout))
	}
	if BackgroundColor != nil {
		queryParameters.Add(`backgroundColor`, *BackgroundColor)
	}
	if DomainMaskEmail != nil {
		queryParameters.Add(`domainMaskEmail`, *DomainMaskEmail)
	}

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SendAMessageToTicketTypes`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string EventId
func (t *EmailMessage) SimulateEmailMessageStatsForEvent(EventId string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`eventId`, EventId)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/SimulateEmailMessageStatsForEvent`,
		&queryParameters,
		nil,
		nil,
	)
}

// @param string MongoQueueId
// @param string EmailDesignId
// @param array Targets confirmed|purchased|assigned|unconfirmed|attended
// @param string MessageSendTime
// @param string Timezone Africa/Abidjan|Africa/Accra|Africa/Addis_Ababa|Africa/Algiers|Africa/Asmara|Africa/Bamako|Africa/Bangui|Africa/Banjul|Africa/Bissau|Africa/Blantyre|Africa/Brazzaville|Africa/Bujumbura|Africa/Cairo|Africa/Casablanca|Africa/Ceuta|Africa/Conakry|Africa/Dakar|Africa/Dar_es_Salaam|Africa/Djibouti|Africa/Douala|Africa/El_Aaiun|Africa/Freetown|Africa/Gaborone|Africa/Harare|Africa/Johannesburg|Africa/Juba|Africa/Kampala|Africa/Khartoum|Africa/Kigali|Africa/Kinshasa|Africa/Lagos|Africa/Libreville|Africa/Lome|Africa/Luanda|Africa/Lubumbashi|Africa/Lusaka|Africa/Malabo|Africa/Maputo|Africa/Maseru|Africa/Mbabane|Africa/Mogadishu|Africa/Monrovia|Africa/Nairobi|Africa/Ndjamena|Africa/Niamey|Africa/Nouakchott|Africa/Ouagadougou|Africa/Porto-Novo|Africa/Sao_Tome|Africa/Tripoli|Africa/Tunis|Africa/Windhoek|America/Adak|America/Anchorage|America/Anguilla|America/Antigua|America/Araguaina|America/Argentina/Buenos_Aires|America/Argentina/Catamarca|America/Argentina/Cordoba|America/Argentina/Jujuy|America/Argentina/La_Rioja|America/Argentina/Mendoza|America/Argentina/Rio_Gallegos|America/Argentina/Salta|America/Argentina/San_Juan|America/Argentina/San_Luis|America/Argentina/Tucuman|America/Argentina/Ushuaia|America/Aruba|America/Asuncion|America/Atikokan|America/Bahia|America/Bahia_Banderas|America/Barbados|America/Belem|America/Belize|America/Blanc-Sablon|America/Boa_Vista|America/Bogota|America/Boise|America/Cambridge_Bay|America/Campo_Grande|America/Cancun|America/Caracas|America/Cayenne|America/Cayman|America/Chicago|America/Chihuahua|America/Costa_Rica|America/Creston|America/Cuiaba|America/Curacao|America/Danmarkshavn|America/Dawson|America/Dawson_Creek|America/Denver|America/Detroit|America/Dominica|America/Edmonton|America/Eirunepe|America/El_Salvador|America/Fort_Nelson|America/Fortaleza|America/Glace_Bay|America/Godthab|America/Goose_Bay|America/Grand_Turk|America/Grenada|America/Guadeloupe|America/Guatemala|America/Guayaquil|America/Guyana|America/Halifax|America/Havana|America/Hermosillo|America/Indiana/Indianapolis|America/Indiana/Knox|America/Indiana/Marengo|America/Indiana/Petersburg|America/Indiana/Tell_City|America/Indiana/Vevay|America/Indiana/Vincennes|America/Indiana/Winamac|America/Inuvik|America/Iqaluit|America/Jamaica|America/Juneau|America/Kentucky/Louisville|America/Kentucky/Monticello|America/Kralendijk|America/La_Paz|America/Lima|America/Los_Angeles|America/Lower_Princes|America/Maceio|America/Managua|America/Manaus|America/Marigot|America/Martinique|America/Matamoros|America/Mazatlan|America/Menominee|America/Merida|America/Metlakatla|America/Mexico_City|America/Miquelon|America/Moncton|America/Monterrey|America/Montevideo|America/Montserrat|America/Nassau|America/New_York|America/Nipigon|America/Nome|America/Noronha|America/North_Dakota/Beulah|America/North_Dakota/Center|America/North_Dakota/New_Salem|America/Ojinaga|America/Panama|America/Pangnirtung|America/Paramaribo|America/Phoenix|America/Port-au-Prince|America/Port_of_Spain|America/Porto_Velho|America/Puerto_Rico|America/Punta_Arenas|America/Rainy_River|America/Rankin_Inlet|America/Recife|America/Regina|America/Resolute|America/Rio_Branco|America/Santarem|America/Santiago|America/Santo_Domingo|America/Sao_Paulo|America/Scoresbysund|America/Sitka|America/St_Barthelemy|America/St_Johns|America/St_Kitts|America/St_Lucia|America/St_Thomas|America/St_Vincent|America/Swift_Current|America/Tegucigalpa|America/Thule|America/Thunder_Bay|America/Tijuana|America/Toronto|America/Tortola|America/Vancouver|America/Whitehorse|America/Winnipeg|America/Yakutat|America/Yellowknife|Antarctica/Casey|Antarctica/Davis|Antarctica/DumontDUrville|Antarctica/Macquarie|Antarctica/Mawson|Antarctica/McMurdo|Antarctica/Palmer|Antarctica/Rothera|Antarctica/Syowa|Antarctica/Troll|Antarctica/Vostok|Arctic/Longyearbyen|Asia/Aden|Asia/Almaty|Asia/Amman|Asia/Anadyr|Asia/Aqtau|Asia/Aqtobe|Asia/Ashgabat|Asia/Atyrau|Asia/Baghdad|Asia/Bahrain|Asia/Baku|Asia/Bangkok|Asia/Barnaul|Asia/Beirut|Asia/Bishkek|Asia/Brunei|Asia/Chita|Asia/Choibalsan|Asia/Colombo|Asia/Damascus|Asia/Dhaka|Asia/Dili|Asia/Dubai|Asia/Dushanbe|Asia/Famagusta|Asia/Gaza|Asia/Hebron|Asia/Ho_Chi_Minh|Asia/Hong_Kong|Asia/Hovd|Asia/Irkutsk|Asia/Jakarta|Asia/Jayapura|Asia/Jerusalem|Asia/Kabul|Asia/Kamchatka|Asia/Karachi|Asia/Kathmandu|Asia/Khandyga|Asia/Kolkata|Asia/Krasnoyarsk|Asia/Kuala_Lumpur|Asia/Kuching|Asia/Kuwait|Asia/Macau|Asia/Magadan|Asia/Makassar|Asia/Manila|Asia/Muscat|Asia/Nicosia|Asia/Novokuznetsk|Asia/Novosibirsk|Asia/Omsk|Asia/Oral|Asia/Phnom_Penh|Asia/Pontianak|Asia/Pyongyang|Asia/Qatar|Asia/Qyzylorda|Asia/Riyadh|Asia/Sakhalin|Asia/Samarkand|Asia/Seoul|Asia/Shanghai|Asia/Singapore|Asia/Srednekolymsk|Asia/Taipei|Asia/Tashkent|Asia/Tbilisi|Asia/Tehran|Asia/Thimphu|Asia/Tokyo|Asia/Tomsk|Asia/Ulaanbaatar|Asia/Urumqi|Asia/Ust-Nera|Asia/Vientiane|Asia/Vladivostok|Asia/Yakutsk|Asia/Yangon|Asia/Yekaterinburg|Asia/Yerevan|Atlantic/Azores|Atlantic/Bermuda|Atlantic/Canary|Atlantic/Cape_Verde|Atlantic/Faroe|Atlantic/Madeira|Atlantic/Reykjavik|Atlantic/South_Georgia|Atlantic/St_Helena|Atlantic/Stanley|Australia/Adelaide|Australia/Brisbane|Australia/Broken_Hill|Australia/Currie|Australia/Darwin|Australia/Eucla|Australia/Hobart|Australia/Lindeman|Australia/Lord_Howe|Australia/Melbourne|Australia/Perth|Australia/Sydney|Europe/Amsterdam|Europe/Andorra|Europe/Astrakhan|Europe/Athens|Europe/Belgrade|Europe/Berlin|Europe/Bratislava|Europe/Brussels|Europe/Bucharest|Europe/Budapest|Europe/Busingen|Europe/Chisinau|Europe/Copenhagen|Europe/Dublin|Europe/Gibraltar|Europe/Guernsey|Europe/Helsinki|Europe/Isle_of_Man|Europe/Istanbul|Europe/Jersey|Europe/Kaliningrad|Europe/Kiev|Europe/Kirov|Europe/Lisbon|Europe/Ljubljana|Europe/London|Europe/Luxembourg|Europe/Madrid|Europe/Malta|Europe/Mariehamn|Europe/Minsk|Europe/Monaco|Europe/Moscow|Europe/Oslo|Europe/Paris|Europe/Podgorica|Europe/Prague|Europe/Riga|Europe/Rome|Europe/Samara|Europe/San_Marino|Europe/Sarajevo|Europe/Saratov|Europe/Simferopol|Europe/Skopje|Europe/Sofia|Europe/Stockholm|Europe/Tallinn|Europe/Tirane|Europe/Ulyanovsk|Europe/Uzhgorod|Europe/Vaduz|Europe/Vatican|Europe/Vienna|Europe/Vilnius|Europe/Volgograd|Europe/Warsaw|Europe/Zagreb|Europe/Zaporozhye|Europe/Zurich|Indian/Antananarivo|Indian/Chagos|Indian/Christmas|Indian/Cocos|Indian/Comoro|Indian/Kerguelen|Indian/Mahe|Indian/Maldives|Indian/Mauritius|Indian/Mayotte|Indian/Reunion|Pacific/Apia|Pacific/Auckland|Pacific/Bougainville|Pacific/Chatham|Pacific/Chuuk|Pacific/Easter|Pacific/Efate|Pacific/Enderbury|Pacific/Fakaofo|Pacific/Fiji|Pacific/Funafuti|Pacific/Galapagos|Pacific/Gambier|Pacific/Guadalcanal|Pacific/Guam|Pacific/Honolulu|Pacific/Kiritimati|Pacific/Kosrae|Pacific/Kwajalein|Pacific/Majuro|Pacific/Marquesas|Pacific/Midway|Pacific/Nauru|Pacific/Niue|Pacific/Norfolk|Pacific/Noumea|Pacific/Pago_Pago|Pacific/Palau|Pacific/Pitcairn|Pacific/Pohnpei|Pacific/Port_Moresby|Pacific/Rarotonga|Pacific/Saipan|Pacific/Tahiti|Pacific/Tarawa|Pacific/Tongatapu|Pacific/Wake|Pacific/Wallis|UTC
func (t *EmailMessage) UpdateScheduledBatchEmailMessage(MongoQueueId string, EmailDesignId string, Targets []string, MessageSendTime string, Timezone string) (r *http.Response, err error) {
	queryParameters := url.Values{}
	queryParameters.Add(`mongoQueueId`, MongoQueueId)
	queryParameters.Add(`emailDesignId`, EmailDesignId)
	for i := range Targets {
		queryParameters.Add(`targets[]`, Targets[i])
	}
	queryParameters.Add(`messageSendTime`, MessageSendTime)
	queryParameters.Add(`timezone`, Timezone)

	return t.restClient.Post(
		`/v2/EmailMessage/UseCase/UpdateScheduledBatchEmailMessage`,
		&queryParameters,
		nil,
		nil,
	)
}
