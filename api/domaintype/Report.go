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
			IsSentEmails:          true,
			IsWaiver:              false,
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
			IsSentEmails:          true,
			IsWaiver:              false,
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
			IsAdminEvents:         true,
		},
	}
}
