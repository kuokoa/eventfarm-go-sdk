/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Event struct {
}

func NewEvent() *Event {
	return &Event{}
}

type CIOMarketingEventCountType struct {
	Slug                   string
	Name                   string
	Description            string
	IsFewerThanFive        bool
	IsMoreThanThirty       bool
	IsBetweenFiveAndThirty bool
}

type CIOMarketingEventType struct {
	Slug                   string
	Name                   string
	Description            string
	IsBrandAwarenessEvents bool
	IsTradeShows           bool
	IsConferences          bool
	IsRoadshows            bool
	IsInternalMeetings     bool
	IsRecruitingEvents     bool
	IsFundraisingEvents    bool
}

type CIOMarketingIndustryType struct {
	Slug                     string
	Name                     string
	Description              string
	IsAgency                 bool
	IsEducation              bool
	IsFinance                bool
	IsSportsAndEntertainment bool
	IsTechnology             bool
	IsOther                  bool
}

type EventDateFilterType struct {
	Slug                   string
	Name                   string
	Description            string
	IsCurrentFuture        bool
	IsPastAll              bool
	IsPast3Months          bool
	IsPast3MonthsAndFuture bool
	IsPast6Months          bool
}

type EventMessageType struct {
	Slug                  string
	Name                  string
	Description           string
	IsDescription         bool
	IsOpening             bool
	IsClosing             bool
	IsConfirmation        bool
	IsDeclination         bool
	IsSoldOut             bool
	IsReveal              bool
	IsDisclaimer          bool
	IsResponseRestriction bool
}

type EventType struct {
	Slug        string
	Name        string
	Description string
	IsEventFarm bool
	IsCio       bool
	IsListly    bool
	IsDnc       bool
	IsRnc       bool
	IsRslc      bool
	IsSundance  bool
}

type MapSourceType struct {
	Slug        string
	Name        string
	Description string
	IsGoogle    bool
	IsYahoo     bool
	IsBing      bool
}

type QuestionContextType struct {
	Slug           string
	Name           string
	Description    string
	IsRegistration bool
	IsLeadCapture  bool
}

type QuestionType struct {
	Slug        string
	Name        string
	Description string
	IsCheckbox  bool
	IsRadio     bool
	IsMulti     bool
	IsText      bool
	IsSelect    bool
	IsDate      bool
	IsWaiver    bool
}

type TrackingScriptType struct {
	Slug           string
	Name           string
	Description    string
	IsRegistration bool
	IsConfirmation bool
}

func (f *Event) ListCIOMarketingEventCountTypes() []CIOMarketingEventCountType {
	return []CIOMarketingEventCountType{
		{
			Slug:                   `fewer-than-5`,
			Name:                   `Fewer than 5`,
			Description:            ``,
			IsFewerThanFive:        true,
			IsMoreThanThirty:       false,
			IsBetweenFiveAndThirty: false,
		},
		{
			Slug:                   `between-5-and-30`,
			Name:                   `Between 5 and 30`,
			Description:            ``,
			IsFewerThanFive:        false,
			IsMoreThanThirty:       false,
			IsBetweenFiveAndThirty: true,
		},
		{
			Slug:                   `more-than-30`,
			Name:                   `More than 30`,
			Description:            ``,
			IsFewerThanFive:        false,
			IsMoreThanThirty:       true,
			IsBetweenFiveAndThirty: false,
		},
	}
}

func (f *Event) ListCIOMarketingEventTypes() []CIOMarketingEventType {
	return []CIOMarketingEventType{
		{
			Slug:                   `brand-awareness-events`,
			Name:                   `Brand Awareness Events`,
			Description:            ``,
			IsBrandAwarenessEvents: true,
			IsTradeShows:           false,
			IsConferences:          false,
			IsRoadshows:            false,
			IsInternalMeetings:     false,
			IsRecruitingEvents:     false,
			IsFundraisingEvents:    false,
		},
		{
			Slug:                   `trade-shows`,
			Name:                   `Trade Shows`,
			Description:            ``,
			IsBrandAwarenessEvents: false,
			IsTradeShows:           true,
			IsConferences:          false,
			IsRoadshows:            false,
			IsInternalMeetings:     false,
			IsRecruitingEvents:     false,
			IsFundraisingEvents:    false,
		},
		{
			Slug:                   `conferences`,
			Name:                   `Conferences`,
			Description:            ``,
			IsBrandAwarenessEvents: false,
			IsTradeShows:           false,
			IsConferences:          true,
			IsRoadshows:            false,
			IsInternalMeetings:     false,
			IsRecruitingEvents:     false,
			IsFundraisingEvents:    false,
		},
		{
			Slug:                   `roadshows`,
			Name:                   `Roadshows`,
			Description:            ``,
			IsBrandAwarenessEvents: false,
			IsTradeShows:           false,
			IsConferences:          false,
			IsRoadshows:            true,
			IsInternalMeetings:     false,
			IsRecruitingEvents:     false,
			IsFundraisingEvents:    false,
		},
		{
			Slug:                   `internal-meetings`,
			Name:                   `Internal Meetings`,
			Description:            ``,
			IsBrandAwarenessEvents: false,
			IsTradeShows:           false,
			IsConferences:          false,
			IsRoadshows:            false,
			IsInternalMeetings:     true,
			IsRecruitingEvents:     false,
			IsFundraisingEvents:    false,
		},
		{
			Slug:                   `recruiting-events`,
			Name:                   `Recruiting Events`,
			Description:            ``,
			IsBrandAwarenessEvents: false,
			IsTradeShows:           false,
			IsConferences:          false,
			IsRoadshows:            false,
			IsInternalMeetings:     false,
			IsRecruitingEvents:     true,
			IsFundraisingEvents:    false,
		},
		{
			Slug:                   `fundraising-events`,
			Name:                   `Fundraising Events`,
			Description:            ``,
			IsBrandAwarenessEvents: false,
			IsTradeShows:           false,
			IsConferences:          false,
			IsRoadshows:            false,
			IsInternalMeetings:     false,
			IsRecruitingEvents:     false,
			IsFundraisingEvents:    true,
		},
	}
}

func (f *Event) ListCIOMarketingIndustryTypes() []CIOMarketingIndustryType {
	return []CIOMarketingIndustryType{
		{
			Slug:                     `consumer-goods`,
			Name:                     `Consumer Goods`,
			Description:              ``,
			IsAgency:                 false,
			IsEducation:              false,
			IsFinance:                false,
			IsSportsAndEntertainment: false,
			IsTechnology:             false,
			IsOther:                  false,
		},
		{
			Slug:                     `agency`,
			Name:                     `Agency`,
			Description:              ``,
			IsAgency:                 true,
			IsEducation:              false,
			IsFinance:                false,
			IsSportsAndEntertainment: false,
			IsTechnology:             false,
			IsOther:                  false,
		},
		{
			Slug:                     `education`,
			Name:                     `Education`,
			Description:              ``,
			IsAgency:                 false,
			IsEducation:              true,
			IsFinance:                false,
			IsSportsAndEntertainment: false,
			IsTechnology:             false,
			IsOther:                  false,
		},
		{
			Slug:                     `finance`,
			Name:                     `Finance`,
			Description:              ``,
			IsAgency:                 false,
			IsEducation:              false,
			IsFinance:                true,
			IsSportsAndEntertainment: false,
			IsTechnology:             false,
			IsOther:                  false,
		},
		{
			Slug:                     `sports-and-entertainment`,
			Name:                     `Sports and Entertainment`,
			Description:              ``,
			IsAgency:                 false,
			IsEducation:              false,
			IsFinance:                false,
			IsSportsAndEntertainment: true,
			IsTechnology:             false,
			IsOther:                  false,
		},
		{
			Slug:                     `technology`,
			Name:                     `Technology`,
			Description:              ``,
			IsAgency:                 false,
			IsEducation:              false,
			IsFinance:                false,
			IsSportsAndEntertainment: false,
			IsTechnology:             true,
			IsOther:                  false,
		},
		{
			Slug:                     `other`,
			Name:                     `Other`,
			Description:              ``,
			IsAgency:                 false,
			IsEducation:              false,
			IsFinance:                false,
			IsSportsAndEntertainment: false,
			IsTechnology:             false,
			IsOther:                  true,
		},
	}
}

func (f *Event) ListEventDateFilterTypes() []EventDateFilterType {
	return []EventDateFilterType{
		{
			Slug:                   `current-future`,
			Name:                   `Current &amp; Future Events`,
			Description:            `This will show you all your future and current events.`,
			IsCurrentFuture:        true,
			IsPastAll:              false,
			IsPast3Months:          false,
			IsPast3MonthsAndFuture: false,
			IsPast6Months:          false,
		},
		{
			Slug:                   `past-all`,
			Name:                   `All Past Events`,
			Description:            ``,
			IsCurrentFuture:        false,
			IsPastAll:              true,
			IsPast3Months:          false,
			IsPast3MonthsAndFuture: false,
			IsPast6Months:          false,
		},
		{
			Slug:                   `past-3-months`,
			Name:                   `Past 3 Months`,
			Description:            ``,
			IsCurrentFuture:        false,
			IsPastAll:              false,
			IsPast3Months:          true,
			IsPast3MonthsAndFuture: false,
			IsPast6Months:          false,
		},
		{
			Slug:                   `past-3-months-and-future`,
			Name:                   `Past 3 Months and Future Events`,
			Description:            `This will show you past 3 months events and all future events`,
			IsCurrentFuture:        false,
			IsPastAll:              false,
			IsPast3Months:          false,
			IsPast3MonthsAndFuture: true,
			IsPast6Months:          false,
		},
		{
			Slug:                   `past-6-months`,
			Name:                   `Past 6 Months`,
			Description:            ``,
			IsCurrentFuture:        false,
			IsPastAll:              false,
			IsPast3Months:          false,
			IsPast3MonthsAndFuture: false,
			IsPast6Months:          true,
		},
	}
}

func (f *Event) ListEventMessageTypes() []EventMessageType {
	return []EventMessageType{
		{
			Slug:                  `introduction`,
			Name:                  `Registration Introduction`,
			Description:           `The Introduction Message is displayed to the user at the beginning of the purchase/registration/RSVP process.`,
			IsDescription:         true,
			IsOpening:             false,
			IsClosing:             false,
			IsConfirmation:        false,
			IsDeclination:         false,
			IsSoldOut:             false,
			IsReveal:              false,
			IsDisclaimer:          false,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `before-opening`,
			Name:                  `Prior To Opening`,
			Description:           `The Prior to Opening Message is displayed as a placeholder for the registration or purchase form before the event goes live and tickets become available.`,
			IsDescription:         false,
			IsOpening:             true,
			IsClosing:             false,
			IsConfirmation:        false,
			IsDeclination:         false,
			IsSoldOut:             false,
			IsReveal:              false,
			IsDisclaimer:          false,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `after-closing`,
			Name:                  `After Closing`,
			Description:           `The After Closing Message is displayed as a placeholder for the registration or purchase form once registration has closed for an event.`,
			IsDescription:         false,
			IsOpening:             false,
			IsClosing:             true,
			IsConfirmation:        false,
			IsDeclination:         false,
			IsSoldOut:             false,
			IsReveal:              false,
			IsDisclaimer:          false,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `confirmation`,
			Name:                  `Confirmation`,
			Description:           `This message will display on the confirmation screen after a successful purchase, registration or RSVP.`,
			IsDescription:         false,
			IsOpening:             false,
			IsClosing:             false,
			IsConfirmation:        true,
			IsDeclination:         false,
			IsSoldOut:             false,
			IsReveal:              false,
			IsDisclaimer:          false,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `declination`,
			Name:                  `Declination`,
			Description:           `This message will display on the confirmation screen after a guest declines their RSVP.`,
			IsDescription:         false,
			IsOpening:             false,
			IsClosing:             false,
			IsConfirmation:        false,
			IsDeclination:         true,
			IsSoldOut:             false,
			IsReveal:              false,
			IsDisclaimer:          false,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `sold-out`,
			Name:                  `Sold-Out`,
			Description:           `The Sold-Out Message is displayed when the specific ticket stack or ticket type is no longer available (either sold out or at capacity).`,
			IsDescription:         false,
			IsOpening:             false,
			IsClosing:             false,
			IsConfirmation:        false,
			IsDeclination:         false,
			IsSoldOut:             true,
			IsReveal:              false,
			IsDisclaimer:          false,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `reveal`,
			Name:                  `Reveal`,
			Description:           `If Invitation Reveal is enabled, the &quot;Reveal Message&quot; will display when a guest begins the registration process.`,
			IsDescription:         false,
			IsOpening:             false,
			IsClosing:             false,
			IsConfirmation:        false,
			IsDeclination:         false,
			IsSoldOut:             false,
			IsReveal:              true,
			IsDisclaimer:          false,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `disclaimer`,
			Name:                  `Disclaimer`,
			Description:           `If you have created a Waiver question, the Disclaimer Message will display as part of the purchase, registration, and RSVP process. Guest must check the disclaimer box to complete registration.`,
			IsDescription:         false,
			IsOpening:             false,
			IsClosing:             false,
			IsConfirmation:        false,
			IsDeclination:         false,
			IsSoldOut:             false,
			IsReveal:              false,
			IsDisclaimer:          true,
			IsResponseRestriction: false,
		},
		{
			Slug:                  `response-restriction`,
			Name:                  `Response Restriction`,
			Description:           `If the Event Setting “Guest can change response” is “NO”, this message will display for any registered guest who clicks the invitation link again.`,
			IsDescription:         false,
			IsOpening:             false,
			IsClosing:             false,
			IsConfirmation:        false,
			IsDeclination:         false,
			IsSoldOut:             false,
			IsReveal:              false,
			IsDisclaimer:          false,
			IsResponseRestriction: true,
		},
	}
}

func (f *Event) ListEventTypes() []EventType {
	return []EventType{
		{
			Slug:        `eventfarm`,
			Name:        `Event Farm`,
			Description: ``,
			IsEventFarm: true,
			IsCio:       false,
			IsListly:    false,
			IsDnc:       false,
			IsRnc:       false,
			IsRslc:      false,
			IsSundance:  false,
		},
		{
			Slug:        `cio`,
			Name:        `Check-in-Only`,
			Description: ``,
			IsEventFarm: false,
			IsCio:       true,
			IsListly:    false,
			IsDnc:       false,
			IsRnc:       false,
			IsRslc:      false,
			IsSundance:  false,
		},
		{
			Slug:        `listly`,
			Name:        `Listly`,
			Description: ``,
			IsEventFarm: false,
			IsCio:       false,
			IsListly:    true,
			IsDnc:       false,
			IsRnc:       false,
			IsRslc:      false,
			IsSundance:  false,
		},
		{
			Slug:        `dnc`,
			Name:        `DNC`,
			Description: ``,
			IsEventFarm: false,
			IsCio:       false,
			IsListly:    false,
			IsDnc:       true,
			IsRnc:       false,
			IsRslc:      false,
			IsSundance:  false,
		},
		{
			Slug:        `rnc`,
			Name:        `RNC`,
			Description: ``,
			IsEventFarm: false,
			IsCio:       false,
			IsListly:    false,
			IsDnc:       false,
			IsRnc:       true,
			IsRslc:      false,
			IsSundance:  false,
		},
		{
			Slug:        `rslc`,
			Name:        `RSLC`,
			Description: ``,
			IsEventFarm: false,
			IsCio:       false,
			IsListly:    false,
			IsDnc:       false,
			IsRnc:       false,
			IsRslc:      true,
			IsSundance:  false,
		},
		{
			Slug:        `sundance`,
			Name:        `Sundance`,
			Description: ``,
			IsEventFarm: false,
			IsCio:       false,
			IsListly:    false,
			IsDnc:       false,
			IsRnc:       false,
			IsRslc:      false,
			IsSundance:  true,
		},
	}
}

func (f *Event) ListMapSourceTypes() []MapSourceType {
	return []MapSourceType{
		{
			Slug:        `google`,
			Name:        `Google`,
			Description: `Google Maps`,
			IsGoogle:    true,
			IsYahoo:     false,
			IsBing:      false,
		},
		{
			Slug:        `yahoo`,
			Name:        `Yahoo`,
			Description: `Yahoo Maps`,
			IsGoogle:    false,
			IsYahoo:     true,
			IsBing:      false,
		},
		{
			Slug:        `bing`,
			Name:        `Bing`,
			Description: `Bing Maps`,
			IsGoogle:    false,
			IsYahoo:     false,
			IsBing:      true,
		},
	}
}

func (f *Event) ListQuestionContextTypes() []QuestionContextType {
	return []QuestionContextType{
		{
			Slug:           `registration`,
			Name:           `Registration`,
			Description:    ``,
			IsRegistration: true,
			IsLeadCapture:  false,
		},
		{
			Slug:           `lead`,
			Name:           `Lead Capture`,
			Description:    ``,
			IsRegistration: false,
			IsLeadCapture:  true,
		},
	}
}

func (f *Event) ListQuestionTypes() []QuestionType {
	return []QuestionType{
		{
			Slug:        `checkbox`,
			Name:        `Checkbox`,
			Description: ``,
			IsCheckbox:  true,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
		},
		{
			Slug:        `radio`,
			Name:        `Radio`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     true,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
		},
		{
			Slug:        `multi`,
			Name:        `Multi`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     true,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
		},
		{
			Slug:        `text`,
			Name:        `Text`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      true,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    false,
		},
		{
			Slug:        `select`,
			Name:        `Select`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    true,
			IsDate:      false,
			IsWaiver:    false,
		},
		{
			Slug:        `date`,
			Name:        `Date`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      true,
			IsWaiver:    false,
		},
		{
			Slug:        `waiver`,
			Name:        `Waiver`,
			Description: ``,
			IsCheckbox:  false,
			IsRadio:     false,
			IsMulti:     false,
			IsText:      false,
			IsSelect:    false,
			IsDate:      false,
			IsWaiver:    true,
		},
	}
}

func (f *Event) ListTrackingScriptTypes() []TrackingScriptType {
	return []TrackingScriptType{
		{
			Slug:           `registration`,
			Name:           `Registration Tracking Script`,
			Description:    ``,
			IsRegistration: true,
			IsConfirmation: false,
		},
		{
			Slug:           `confirmation`,
			Name:           `Confirmation Tracking Script`,
			Description:    ``,
			IsRegistration: false,
			IsConfirmation: true,
		},
	}
}
