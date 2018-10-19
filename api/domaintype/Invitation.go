package domaintype

type Invitation struct {
}

func NewInvitation() *Invitation {
	return &Invitation{}
}

type InvitationActionType struct {
	Slug           string
	Name           string
	Description    string
	IsNone         bool
	IsClicked      bool
	IsReceived     bool
	IsProcessed    bool
	IsSpamReport   bool
	IsBounce       bool
	IsDeferred     bool
	IsDropped      bool
	IsOpen         bool
	IsDelivered    bool
	IsRs           bool
	IsBlocked      bool
	IsUnsubscribed bool
	IsUnknown      bool
}

type InvitationCreationType struct {
	Slug                 string
	Name                 string
	Description          string
	IsConfirmedNoEmail   bool
	IsUnconfirmedNoEmail bool
	IsSendEmail          bool
	ShouldSendEmail      bool
}

type InvitationStatusType struct {
	Slug                                 string
	Name                                 string
	Description                          string
	IsAssigned                           bool
	IsPurchased                          bool
	IsConfirmedByRSVP                    bool
	IsDeclinedByRSVP                     bool
	IsLeftBehind                         bool
	IsNotYetPurchased                    bool
	IsRegistered                         bool
	IsUnconfirmed                        bool
	IsRecycled                           bool
	IsNotYetRegistered                   bool
	IsWaitlisted                         bool
	IsAffirmativeResponse                bool
	IsUsed                               bool
	IsSendToConfirmedOrRegisteredTypeIds bool
	IsSendToPurchasedTypeIds             bool
	IsSendToAssignedTypeIds              bool
	IsSendToUnconfirmedTypeIds           bool
}

type InviteSourceType struct {
	Slug                string
	Name                string
	Description         string
	IsDirectInvite      bool
	IsApiAddition       bool
	IsDistribution      bool
	IsEventInvite       bool
	IsGroupInvite       bool
	IsImport            bool
	IsImportMarketo     bool
	IsImportSalesforce  bool
	IsIosLeaveBehind    bool
	IsLeaveBehind       bool
	IsMobileLeaveBehind bool
	IsMobilePurchase    bool
	IsOutsidePurchase   bool
	IsPublicInterface   bool
	IsTransferred       bool
}

type WebhookType struct {
	Slug             string
	Name             string
	Description      string
	IsAffirmative    bool
	IsAssigned       bool
	IsCheckin        bool
	IsConfirmed      bool
	IsCreated        bool
	IsLeaveBehind    bool
	IsPurchased      bool
	IsRegistered     bool
	IsRemoved        bool
	IsUpdated        bool
	IsAffirmativeSMS bool
	IsUpdatedSMS     bool
	IsCreatedSMS     bool
}

func (f *Invitation) ListInvitationActionTypes() []InvitationActionType {
	return []InvitationActionType{
		{
			Slug:           `none`,
			Name:           `None`,
			Description:    ``,
			IsNone:         true,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `clicked`,
			Name:           `Clicked`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      true,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `received`,
			Name:           `Received`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     true,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `processed`,
			Name:           `Processed`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    true,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `spam-report`,
			Name:           `Spam Report`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   true,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `bounce`,
			Name:           `Bounce`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       true,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `deferred`,
			Name:           `Deferred`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     true,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `dropped`,
			Name:           `Dropped`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      true,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `open`,
			Name:           `Open`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         true,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `delivered`,
			Name:           `Delivered`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    true,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `rs`,
			Name:           `Rs`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           true,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `blocked`,
			Name:           `Blocked`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      true,
			IsUnsubscribed: false,
			IsUnknown:      false,
		},
		{
			Slug:           `unsubscribed`,
			Name:           `Unsubscribed`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: true,
			IsUnknown:      false,
		},
		{
			Slug:           `unknown`,
			Name:           `Unknown`,
			Description:    ``,
			IsNone:         false,
			IsClicked:      false,
			IsReceived:     false,
			IsProcessed:    false,
			IsSpamReport:   false,
			IsBounce:       false,
			IsDeferred:     false,
			IsDropped:      false,
			IsOpen:         false,
			IsDelivered:    false,
			IsRs:           false,
			IsBlocked:      false,
			IsUnsubscribed: false,
			IsUnknown:      true,
		},
	}
}

func (f *Invitation) ListInvitationCreationTypes() []InvitationCreationType {
	return []InvitationCreationType{
		{
			Slug:                 `unconfirmed-no-email`,
			Name:                 `Unconfirmed No Email`,
			Description:          `Add as Unconfirmed, email will not be sent.`,
			IsConfirmedNoEmail:   false,
			IsUnconfirmedNoEmail: true,
			IsSendEmail:          false,
			ShouldSendEmail:      false,
		},
		{
			Slug:                 `confirmed-no-email`,
			Name:                 `Confirmed No Email`,
			Description:          `Add as Confirmed, email will not be sent.`,
			IsConfirmedNoEmail:   true,
			IsUnconfirmedNoEmail: false,
			IsSendEmail:          false,
			ShouldSendEmail:      false,
		},
		{
			Slug:                 `send-email`,
			Name:                 `Send Invitation`,
			Description:          `Send Invitation by Email.`,
			IsConfirmedNoEmail:   false,
			IsUnconfirmedNoEmail: false,
			IsSendEmail:          true,
			ShouldSendEmail:      true,
		},
	}
}

func (f *Invitation) ListInvitationStatusTypes() []InvitationStatusType {
	return []InvitationStatusType{
		{
			Slug:                                 `assigned`,
			Name:                                 `Assigned`,
			Description:                          ``,
			IsAssigned:                           true,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                true,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: true,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              true,
			IsSendToUnconfirmedTypeIds:           false,
		},
		{
			Slug:                                 `purchased`,
			Name:                                 `Purchased`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          true,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                true,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             true,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           false,
		},
		{
			Slug:                                 `confirmed-by-rsvp`,
			Name:                                 `Confirmed By RSVP`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    true,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                true,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: true,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           false,
		},
		{
			Slug:                                 `declined-by-rsvp`,
			Name:                                 `Declined By RSVP`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     true,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                false,
			IsUsed:                               false,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           false,
		},
		{
			Slug:                                 `left-behind`,
			Name:                                 `Left Behind`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         true,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                true,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           false,
		},
		{
			Slug:                                 `not-yet-purchased`,
			Name:                                 `Not Yet Purchased`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    true,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                false,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           true,
		},
		{
			Slug:                                 `registered`,
			Name:                                 `Registered`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         true,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                true,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: true,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           false,
		},
		{
			Slug:                                 `unconfirmed`,
			Name:                                 `Unconfirmed`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        true,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                false,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           true,
		},
		{
			Slug:                                 `recycled`,
			Name:                                 `Recycled`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           true,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                false,
			IsUsed:                               false,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           false,
		},
		{
			Slug:                                 `not-yet-registered`,
			Name:                                 `Not Yet Registered`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   true,
			IsWaitlisted:                         false,
			IsAffirmativeResponse:                false,
			IsUsed:                               true,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           true,
		},
		{
			Slug:                                 `waitlisted`,
			Name:                                 `Waitlisted`,
			Description:                          ``,
			IsAssigned:                           false,
			IsPurchased:                          false,
			IsConfirmedByRSVP:                    false,
			IsDeclinedByRSVP:                     false,
			IsLeftBehind:                         false,
			IsNotYetPurchased:                    false,
			IsRegistered:                         false,
			IsUnconfirmed:                        false,
			IsRecycled:                           false,
			IsNotYetRegistered:                   false,
			IsWaitlisted:                         true,
			IsAffirmativeResponse:                false,
			IsUsed:                               false,
			IsSendToConfirmedOrRegisteredTypeIds: false,
			IsSendToPurchasedTypeIds:             false,
			IsSendToAssignedTypeIds:              false,
			IsSendToUnconfirmedTypeIds:           false,
		},
	}
}

func (f *Invitation) ListInviteSourceTypes() []InviteSourceType {
	return []InviteSourceType{
		{
			Slug:                `direct-invite`,
			Name:                `Direct Invite`,
			Description:         ``,
			IsDirectInvite:      true,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `api-addition`,
			Name:                `Api Addition`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       true,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `distribution`,
			Name:                `Distribution`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      true,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `event-invite`,
			Name:                `Event Invite`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       true,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `group-invite`,
			Name:                `Group Invite`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       true,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `import`,
			Name:                `Import`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            true,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `ios-leave-behind`,
			Name:                `iOS Leave Behind`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    true,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `leave-behind`,
			Name:                `Leave Behind`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       true,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `mobile-leave-behind`,
			Name:                `Mobile Leave Behind`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: true,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `mobile-purchase`,
			Name:                `Mobile Purchase`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    true,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `outside-purchase`,
			Name:                `Outside Purchase`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   true,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `public-interface`,
			Name:                `Public Interface`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   true,
			IsTransferred:       false,
		},
		{
			Slug:                `transferred`,
			Name:                `Transferred`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       true,
		},
		{
			Slug:                `import-salesforce`,
			Name:                `Import from Salesforce`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     false,
			IsImportSalesforce:  true,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
		{
			Slug:                `import-marketo`,
			Name:                `Import from Marketo`,
			Description:         ``,
			IsDirectInvite:      false,
			IsApiAddition:       false,
			IsDistribution:      false,
			IsEventInvite:       false,
			IsGroupInvite:       false,
			IsImport:            false,
			IsImportMarketo:     true,
			IsImportSalesforce:  false,
			IsIosLeaveBehind:    false,
			IsLeaveBehind:       false,
			IsMobileLeaveBehind: false,
			IsMobilePurchase:    false,
			IsOutsidePurchase:   false,
			IsPublicInterface:   false,
			IsTransferred:       false,
		},
	}
}

func (f *Invitation) ListWebhookTypes() []WebhookType {
	return []WebhookType{
		{
			Slug:             `affirmative`,
			Name:             `Affirmative`,
			Description:      ``,
			IsAffirmative:    true,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `assigned`,
			Name:             `Assigned`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       true,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `checkin`,
			Name:             `Check-In`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        true,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `confirmed`,
			Name:             `Confirmed`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      true,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `created`,
			Name:             `Created`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        true,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `leave-behind`,
			Name:             `Leave Behind`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    true,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `purchased`,
			Name:             `Purchased`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      true,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `registered`,
			Name:             `Registered`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     true,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `removed`,
			Name:             `Updated`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        true,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `updated`,
			Name:             `Updated`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        true,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `affirmative-sms`,
			Name:             `Affirmative SMS`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: true,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `updated-sms`,
			Name:             `Updated SMS`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     true,
			IsCreatedSMS:     false,
		},
		{
			Slug:             `created-sms`,
			Name:             `Created SMS`,
			Description:      ``,
			IsAffirmative:    false,
			IsAssigned:       false,
			IsCheckin:        false,
			IsConfirmed:      false,
			IsCreated:        false,
			IsLeaveBehind:    false,
			IsPurchased:      false,
			IsRegistered:     false,
			IsRemoved:        false,
			IsUpdated:        false,
			IsAffirmativeSMS: false,
			IsUpdatedSMS:     false,
			IsCreatedSMS:     true,
		},
	}
}
