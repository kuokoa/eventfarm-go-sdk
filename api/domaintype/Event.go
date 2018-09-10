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

func (f *Event) GetCIOMarketingEventCountType() []CIOMarketingEventCountType {
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

func (f *Event) GetCIOMarketingEventType() []CIOMarketingEventType {
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

func (f *Event) GetCIOMarketingIndustryType() []CIOMarketingIndustryType {
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

func (f *Event) GetEventDateFilterType() []EventDateFilterType {
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

func (f *Event) GetEventType() []EventType {
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

func (f *Event) GetMapSourceType() []MapSourceType {
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

func (f *Event) GetQuestionContextType() []QuestionContextType {
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

func (f *Event) GetQuestionType() []QuestionType {
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
