package domaintype

type IntegrationStatusMapping struct {
}

func NewIntegrationStatusMapping() *IntegrationStatusMapping {
	return &IntegrationStatusMapping{}
}

type StatusIdType struct {
	Slug                  string
	Name                  string
	Description           string
	IsAssigned            bool
	IsPurchased           bool
	IsConfirmedByRSVP     bool
	IsDeclinedByRSVP      bool
	IsLeftBehind          bool
	IsNotYetPurchased     bool
	IsRegistered          bool
	IsUnconfirmed         bool
	IsRecycled            bool
	IsNotYetRegistered    bool
	IsWaitlisted          bool
	IsAffirmativeResponse bool
	IsUsed                bool
	CheckedIn             bool
}

type StatusMappingType struct {
	Slug                       string
	Name                       string
	Description                string
	IsSalesforceCampaignMember bool
}

func (f *IntegrationStatusMapping) ListStatusIdTypes() []StatusIdType {
	return []StatusIdType{
		{
			Slug:                  `assigned`,
			Name:                  `Assigned`,
			Description:           ``,
			IsAssigned:            true,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: true,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `purchased`,
			Name:                  `Purchased`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           true,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: true,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `confirmed-by-rsvp`,
			Name:                  `Confirmed By RSVP`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     true,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: true,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `declined-by-rsvp`,
			Name:                  `Declined By RSVP`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      true,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: false,
			IsUsed:                false,
			CheckedIn:             false,
		},
		{
			Slug:                  `left-behind`,
			Name:                  `Left Behind`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          true,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: true,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `not-yet-purchased`,
			Name:                  `Not Yet Purchased`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     true,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: false,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `registered`,
			Name:                  `Registered`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          true,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: true,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `unconfirmed`,
			Name:                  `Unconfirmed`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         true,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: false,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `recycled`,
			Name:                  `Recycled`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            true,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: false,
			IsUsed:                false,
			CheckedIn:             false,
		},
		{
			Slug:                  `not-yet-registered`,
			Name:                  `Not Yet Registered`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    true,
			IsWaitlisted:          false,
			IsAffirmativeResponse: false,
			IsUsed:                true,
			CheckedIn:             false,
		},
		{
			Slug:                  `waitlisted`,
			Name:                  `Waitlisted`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          true,
			IsAffirmativeResponse: false,
			IsUsed:                false,
			CheckedIn:             false,
		},
		{
			Slug:                  `checked-in`,
			Name:                  `Checked In`,
			Description:           ``,
			IsAssigned:            false,
			IsPurchased:           false,
			IsConfirmedByRSVP:     false,
			IsDeclinedByRSVP:      false,
			IsLeftBehind:          false,
			IsNotYetPurchased:     false,
			IsRegistered:          false,
			IsUnconfirmed:         false,
			IsRecycled:            false,
			IsNotYetRegistered:    false,
			IsWaitlisted:          false,
			IsAffirmativeResponse: false,
			IsUsed:                true,
			CheckedIn:             true,
		},
	}
}

func (f *IntegrationStatusMapping) ListStatusMappingTypes() []StatusMappingType {
	return []StatusMappingType{
		{
			Slug:                       `salesforce-campaign-member`,
			Name:                       `Salesforce Campaign Member Status Mapping`,
			Description:                `Maps the Event Farm statuses to Salesforce campaign member statuses.`,
			IsSalesforceCampaignMember: true,
		},
	}
}
