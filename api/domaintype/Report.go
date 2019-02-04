/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Report struct {
}

func NewReport() *Report {
	return &Report{}
}

type ReportFormatType struct {
	Slug        string
	Name        string
	Description string
	IsCsv       bool
	IsHtml      bool
	IsPdf       bool
	IsZip       bool
}

type ReportStatusType struct {
	Slug        string
	Name        string
	Description string
	IsNew       bool
	IsPending   bool
	IsDone      bool
	IsError     bool
}

type ReportType struct {
	Slug                  string
	Name                  string
	Description           string
	IsInvitation          bool
	IsConfirmed           bool
	IsUnconfirmed         bool
	IsPurchased           bool
	IsWaitlist            bool
	IsCheckin             bool
	IsTransfer            bool
	IsTicketblockUsers    bool
	IsEmailDeliverability bool
	IsSentEmails          bool
	IsWaiver              bool
	IsActivityLog         bool
	IsAdminEvents         bool
}

func (f *Report) ListReportFormatTypes() []ReportFormatType {
	return []ReportFormatType{
		{
			Slug:        `csv`,
			Name:        `csv`,
			Description: `csv`,
			IsCsv:       true,
			IsHtml:      false,
			IsPdf:       false,
			IsZip:       false,
		},
		{
			Slug:        `html`,
			Name:        `html`,
			Description: `html`,
			IsCsv:       false,
			IsHtml:      true,
			IsPdf:       false,
			IsZip:       false,
		},
		{
			Slug:        `pdf`,
			Name:        `pdf`,
			Description: `pdf`,
			IsCsv:       false,
			IsHtml:      false,
			IsPdf:       true,
			IsZip:       false,
		},
		{
			Slug:        `zip`,
			Name:        `zip`,
			Description: `zip`,
			IsCsv:       false,
			IsHtml:      false,
			IsPdf:       false,
			IsZip:       true,
		},
	}
}

func (f *Report) ListReportStatusTypes() []ReportStatusType {
	return []ReportStatusType{
		{
			Slug:        `new`,
			Name:        `New`,
			Description: ``,
			IsNew:       true,
			IsPending:   false,
			IsDone:      false,
			IsError:     false,
		},
		{
			Slug:        `pending`,
			Name:        `Pending`,
			Description: ``,
			IsNew:       false,
			IsPending:   true,
			IsDone:      false,
			IsError:     false,
		},
		{
			Slug:        `done`,
			Name:        `Done`,
			Description: ``,
			IsNew:       false,
			IsPending:   false,
			IsDone:      true,
			IsError:     false,
		},
		{
			Slug:        `error`,
			Name:        `Error`,
			Description: ``,
			IsNew:       false,
			IsPending:   false,
			IsDone:      false,
			IsError:     true,
		},
	}
}

func (f *Report) ListReportTypes() []ReportType {
	return []ReportType{
		{
			Slug:                  `invitation`,
			Name:                  `Invitation`,
			Description:           `Guest Summary Report`,
			IsInvitation:          true,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `confirmed`,
			Name:                  `Confirmed`,
			Description:           `Confirmed Guest List`,
			IsInvitation:          false,
			IsConfirmed:           true,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `unconfirmed`,
			Name:                  `Unconfirmed`,
			Description:           `Unconfirmed Guest List`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         true,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `purchase`,
			Name:                  `Purchased`,
			Description:           `Ticket Purchase List`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           true,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `waitlist`,
			Name:                  `Waitlist`,
			Description:           `Event Wait List`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            true,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `checkin`,
			Name:                  `Check In`,
			Description:           `Check In Report`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             true,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `transfer`,
			Name:                  `Transfer`,
			Description:           `Event Transfer List`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            true,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `ticketblock`,
			Name:                  `Ticketblock Users`,
			Description:           `Ticket Block User List`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    true,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `email`,
			Name:                  `Email Deliverability`,
			Description:           `Email Deliverability Report`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: true,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `sent`,
			Name:                  `Sent Emails`,
			Description:           `Event Sent Email Report`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          true,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `waiver`,
			Name:                  `Waiver`,
			Description:           `Event Waiver Agreements`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              true,
			IsActivityLog:         false,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `activitylog`,
			Name:                  `Activity Log`,
			Description:           `Log of Events Activities`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         true,
			IsAdminEvents:         false,
		},
		{
			Slug:                  `adminevents`,
			Name:                  `Admin Events`,
			Description:           `Admin Event List`,
			IsInvitation:          false,
			IsConfirmed:           false,
			IsUnconfirmed:         false,
			IsPurchased:           false,
			IsWaitlist:            false,
			IsCheckin:             false,
			IsTransfer:            false,
			IsTicketblockUsers:    false,
			IsEmailDeliverability: false,
			IsSentEmails:          false,
			IsWaiver:              false,
			IsActivityLog:         false,
			IsAdminEvents:         true,
		},
	}
}
