package domaintype

type EmailMessage struct {
}

func NewEmailMessage() *EmailMessage {
	return &EmailMessage{}
}

type EmailArchiveCategoryType struct {
	Slug                    string
	Name                    string
	Description             string
	IsAlert                 bool
	IsInvitation            bool
	IsReceipts              bool
	IsForgotPassword        bool
	IsChangePassword        bool
	IsTransferAuthorization bool
	IsTransferConfirmation  bool
	IsEventMessage          bool
	IsEventMessagePreview   bool
	IsReportNotification    bool
	IsAccount               bool
	IsVerify                bool
	IsSystem                bool
}

type EmailArchiveSubCategoryType struct {
	Slug             string
	Name             string
	Description      string
	IsTicketBlock    bool
	IsGroup          bool
	IsEventCheckedIn bool
	IsTicketType     bool
	IsPreview        bool
	IsNoShow         bool
	IsWaitList       bool
	IsNone           bool
}

type EmailMessageType struct {
	Slug               string
	Name               string
	Description        string
	IsConfirmation     bool
	IsDonation         bool
	IsInvitation       bool
	IsSendAMessage     bool
	IsSendGroupMessage bool
}

func (f *EmailMessage) GetEmailArchiveCategoryType() []EmailArchiveCategoryType {
	return []EmailArchiveCategoryType{
		{
			Slug:                    `alert`,
			Name:                    `Alert`,
			Description:             `Arrival Alert`,
			IsAlert:                 true,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `invitation`,
			Name:                    `Invitation`,
			Description:             `Invitation`,
			IsAlert:                 false,
			IsInvitation:            true,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `receipts`,
			Name:                    `Receipts`,
			Description:             `Receipts`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              true,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `forgotpassword`,
			Name:                    `Forgot Password`,
			Description:             `Forgot Password`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        true,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `changepassword`,
			Name:                    `Change Password`,
			Description:             `Change Password`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        true,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `transferauthorization`,
			Name:                    `Transfer Authorization`,
			Description:             `Transfer Authorization`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: true,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `transferconfirmation`,
			Name:                    `Transfer Confirmation`,
			Description:             `Transfer Authorization`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  true,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `eventmessage`,
			Name:                    `Event Message`,
			Description:             `Event Message`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          true,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `eventmessagepreview`,
			Name:                    `Event Message Preview`,
			Description:             `Event Message Preview`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   true,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `reportnotification`,
			Name:                    `Report Notification`,
			Description:             `Report Notification`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    true,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `account`,
			Name:                    `Account`,
			Description:             `Account`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               true,
			IsVerify:                false,
			IsSystem:                false,
		},
		{
			Slug:                    `verify`,
			Name:                    `verify`,
			Description:             `Verify`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                true,
			IsSystem:                false,
		},
		{
			Slug:                    `system`,
			Name:                    `System`,
			Description:             `System`,
			IsAlert:                 false,
			IsInvitation:            false,
			IsReceipts:              false,
			IsForgotPassword:        false,
			IsChangePassword:        false,
			IsTransferAuthorization: false,
			IsTransferConfirmation:  false,
			IsEventMessage:          false,
			IsEventMessagePreview:   false,
			IsReportNotification:    false,
			IsAccount:               false,
			IsVerify:                false,
			IsSystem:                true,
		},
	}
}

func (f *EmailMessage) GetEmailArchiveSubCategoryType() []EmailArchiveSubCategoryType {
	return []EmailArchiveSubCategoryType{
		{
			Slug:             `ticketblock`,
			Name:             `Ticket Block`,
			Description:      ``,
			IsTicketBlock:    true,
			IsGroup:          false,
			IsEventCheckedIn: false,
			IsTicketType:     false,
			IsPreview:        false,
			IsNoShow:         false,
			IsWaitList:       false,
			IsNone:           false,
		},
		{
			Slug:             `group`,
			Name:             `Group`,
			Description:      ``,
			IsTicketBlock:    false,
			IsGroup:          true,
			IsEventCheckedIn: false,
			IsTicketType:     false,
			IsPreview:        false,
			IsNoShow:         false,
			IsWaitList:       false,
			IsNone:           false,
		},
		{
			Slug:             `eventcheckedin`,
			Name:             `Event Checked-In`,
			Description:      ``,
			IsTicketBlock:    false,
			IsGroup:          false,
			IsEventCheckedIn: true,
			IsTicketType:     false,
			IsPreview:        false,
			IsNoShow:         false,
			IsWaitList:       false,
			IsNone:           false,
		},
		{
			Slug:             `tickettype`,
			Name:             `Ticket Type`,
			Description:      ``,
			IsTicketBlock:    false,
			IsGroup:          false,
			IsEventCheckedIn: false,
			IsTicketType:     true,
			IsPreview:        false,
			IsNoShow:         false,
			IsWaitList:       false,
			IsNone:           false,
		},
		{
			Slug:             `preview`,
			Name:             `Preview`,
			Description:      ``,
			IsTicketBlock:    false,
			IsGroup:          false,
			IsEventCheckedIn: false,
			IsTicketType:     false,
			IsPreview:        true,
			IsNoShow:         false,
			IsWaitList:       false,
			IsNone:           false,
		},
		{
			Slug:             `noshow`,
			Name:             `No-Show`,
			Description:      ``,
			IsTicketBlock:    false,
			IsGroup:          false,
			IsEventCheckedIn: false,
			IsTicketType:     false,
			IsPreview:        false,
			IsNoShow:         true,
			IsWaitList:       false,
			IsNone:           false,
		},
		{
			Slug:             `waitlist`,
			Name:             `Wait List`,
			Description:      ``,
			IsTicketBlock:    false,
			IsGroup:          false,
			IsEventCheckedIn: false,
			IsTicketType:     false,
			IsPreview:        false,
			IsNoShow:         false,
			IsWaitList:       true,
			IsNone:           false,
		},
		{
			Slug:             `None`,
			Name:             `None`,
			Description:      ``,
			IsTicketBlock:    false,
			IsGroup:          false,
			IsEventCheckedIn: false,
			IsTicketType:     false,
			IsPreview:        false,
			IsNoShow:         false,
			IsWaitList:       false,
			IsNone:           true,
		},
	}
}

func (f *EmailMessage) GetEmailMessageType() []EmailMessageType {
	return []EmailMessageType{
		{
			Slug:               `confirmation`,
			Name:               `Confirmation`,
			Description:        `Confirmation Message Type`,
			IsConfirmation:     true,
			IsDonation:         false,
			IsInvitation:       false,
			IsSendAMessage:     false,
			IsSendGroupMessage: false,
		},
		{
			Slug:               `donation`,
			Name:               `Donation`,
			Description:        `Donation Message Type`,
			IsConfirmation:     false,
			IsDonation:         true,
			IsInvitation:       false,
			IsSendAMessage:     false,
			IsSendGroupMessage: false,
		},
		{
			Slug:               `invitation`,
			Name:               `Invitation`,
			Description:        `Invitation Message Type`,
			IsConfirmation:     false,
			IsDonation:         false,
			IsInvitation:       true,
			IsSendAMessage:     false,
			IsSendGroupMessage: false,
		},
		{
			Slug:               `sendamessage`,
			Name:               `Send A Message`,
			Description:        `Send A Message Type`,
			IsConfirmation:     false,
			IsDonation:         false,
			IsInvitation:       false,
			IsSendAMessage:     true,
			IsSendGroupMessage: false,
		},
		{
			Slug:               `sendgroupmessage`,
			Name:               `Send Group Message`,
			Description:        `Send Group Message Type`,
			IsConfirmation:     false,
			IsDonation:         false,
			IsInvitation:       false,
			IsSendAMessage:     false,
			IsSendGroupMessage: true,
		},
	}
}
