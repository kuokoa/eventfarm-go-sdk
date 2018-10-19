package domaintype

type Stack struct {
}

func NewStack() *Stack {
	return &Stack{}
}

type StackMethodType struct {
	Slug                   string
	Name                   string
	Description            string
	IsPublicRegistration   bool
	IsPublicPurchase       bool
	IsInviteToRegister     bool
	IsInviteToPurchase     bool
	IsInviteToRSVP         bool
	IsInviteToRegisterFCFS bool
	IsInviteToPurchaseFCFS bool
	IsInviteToRSVPFCFS     bool
	IsInvitation           bool
	IsFirstComeFirstServe  bool
	IsAnyInviteToPurchase  bool
	IsAnyInviteToRegister  bool
	IsAnyInviteToRSVP      bool
}

func (f *Stack) ListStackMethodTypes() []StackMethodType {
	return []StackMethodType{
		{
			Slug:                   `public-registration`,
			Name:                   `Public Registration`,
			Description:            ``,
			IsPublicRegistration:   true,
			IsPublicPurchase:       false,
			IsInviteToRegister:     false,
			IsInviteToPurchase:     false,
			IsInviteToRSVP:         false,
			IsInviteToRegisterFCFS: false,
			IsInviteToPurchaseFCFS: false,
			IsInviteToRSVPFCFS:     false,
			IsInvitation:           false,
			IsFirstComeFirstServe:  false,
			IsAnyInviteToPurchase:  false,
			IsAnyInviteToRegister:  false,
			IsAnyInviteToRSVP:      false,
		},
		{
			Slug:                   `public-purchase`,
			Name:                   `Public Purchase`,
			Description:            ``,
			IsPublicRegistration:   false,
			IsPublicPurchase:       true,
			IsInviteToRegister:     false,
			IsInviteToPurchase:     false,
			IsInviteToRSVP:         false,
			IsInviteToRegisterFCFS: false,
			IsInviteToPurchaseFCFS: false,
			IsInviteToRSVPFCFS:     false,
			IsInvitation:           false,
			IsFirstComeFirstServe:  false,
			IsAnyInviteToPurchase:  false,
			IsAnyInviteToRegister:  false,
			IsAnyInviteToRSVP:      false,
		},
		{
			Slug:                   `invite-to-register`,
			Name:                   `Invite to Register`,
			Description:            ``,
			IsPublicRegistration:   false,
			IsPublicPurchase:       false,
			IsInviteToRegister:     true,
			IsInviteToPurchase:     false,
			IsInviteToRSVP:         false,
			IsInviteToRegisterFCFS: false,
			IsInviteToPurchaseFCFS: false,
			IsInviteToRSVPFCFS:     false,
			IsInvitation:           true,
			IsFirstComeFirstServe:  false,
			IsAnyInviteToPurchase:  false,
			IsAnyInviteToRegister:  true,
			IsAnyInviteToRSVP:      false,
		},
		{
			Slug:                   `invite-to-purchase`,
			Name:                   `Invite to Purchase`,
			Description:            ``,
			IsPublicRegistration:   false,
			IsPublicPurchase:       false,
			IsInviteToRegister:     false,
			IsInviteToPurchase:     true,
			IsInviteToRSVP:         false,
			IsInviteToRegisterFCFS: false,
			IsInviteToPurchaseFCFS: false,
			IsInviteToRSVPFCFS:     false,
			IsInvitation:           true,
			IsFirstComeFirstServe:  false,
			IsAnyInviteToPurchase:  true,
			IsAnyInviteToRegister:  false,
			IsAnyInviteToRSVP:      false,
		},
		{
			Slug:                   `invite-to-rsvp`,
			Name:                   `Invite to RSVP`,
			Description:            ``,
			IsPublicRegistration:   false,
			IsPublicPurchase:       false,
			IsInviteToRegister:     false,
			IsInviteToPurchase:     false,
			IsInviteToRSVP:         true,
			IsInviteToRegisterFCFS: false,
			IsInviteToPurchaseFCFS: false,
			IsInviteToRSVPFCFS:     false,
			IsInvitation:           true,
			IsFirstComeFirstServe:  false,
			IsAnyInviteToPurchase:  false,
			IsAnyInviteToRegister:  false,
			IsAnyInviteToRSVP:      true,
		},
		{
			Slug:                   `invite-to-register-fcfs`,
			Name:                   `Invite to Register (FCFS)`,
			Description:            ``,
			IsPublicRegistration:   false,
			IsPublicPurchase:       false,
			IsInviteToRegister:     false,
			IsInviteToPurchase:     false,
			IsInviteToRSVP:         false,
			IsInviteToRegisterFCFS: true,
			IsInviteToPurchaseFCFS: false,
			IsInviteToRSVPFCFS:     false,
			IsInvitation:           true,
			IsFirstComeFirstServe:  true,
			IsAnyInviteToPurchase:  false,
			IsAnyInviteToRegister:  true,
			IsAnyInviteToRSVP:      false,
		},
		{
			Slug:                   `invite-to-purchase-fcfs`,
			Name:                   `Invite to Purchase (FCFS)`,
			Description:            ``,
			IsPublicRegistration:   false,
			IsPublicPurchase:       false,
			IsInviteToRegister:     false,
			IsInviteToPurchase:     false,
			IsInviteToRSVP:         false,
			IsInviteToRegisterFCFS: false,
			IsInviteToPurchaseFCFS: true,
			IsInviteToRSVPFCFS:     false,
			IsInvitation:           true,
			IsFirstComeFirstServe:  true,
			IsAnyInviteToPurchase:  true,
			IsAnyInviteToRegister:  false,
			IsAnyInviteToRSVP:      false,
		},
		{
			Slug:                   `invite-to-rsvp-fcfs`,
			Name:                   `Invite to RSVP (FCFS)`,
			Description:            ``,
			IsPublicRegistration:   false,
			IsPublicPurchase:       false,
			IsInviteToRegister:     false,
			IsInviteToPurchase:     false,
			IsInviteToRSVP:         false,
			IsInviteToRegisterFCFS: false,
			IsInviteToPurchaseFCFS: false,
			IsInviteToRSVPFCFS:     true,
			IsInvitation:           true,
			IsFirstComeFirstServe:  true,
			IsAnyInviteToPurchase:  false,
			IsAnyInviteToRegister:  false,
			IsAnyInviteToRSVP:      true,
		},
	}
}
