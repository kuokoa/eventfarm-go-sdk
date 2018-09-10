package domaintype

type Import struct {
}

func NewImport() *Import {
	return &Import{}
}

type ImportColumnHeaderType struct {
	Slug              string
	Name              string
	Description       string
	IsAddress1        bool
	IsAddress2        bool
	IsArrivalAlert    bool
	IsCheckInNotes    bool
	IsCity            bool
	IsCompany         bool
	IsCountry         bool
	IsEmail           bool
	IsFirstName       bool
	IsIgnored         bool
	IsInvitationNotes bool
	IsLastName        bool
	IsOther           bool
	IsPosition        bool
	IsPostalCode      bool
	IsQuantity        bool
	IsState           bool
	IsTelephone       bool
	IsTicketStatus    bool
	IsTicketType      bool
	IsTitle           bool
}

func (f *Import) GetImportColumnHeaderType() []ImportColumnHeaderType {
	return []ImportColumnHeaderType{
		{
			Slug:              `address1`,
			Name:              `Address 1`,
			Description:       ``,
			IsAddress1:        true,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `address2`,
			Name:              `Address 2`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        true,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `arrivalAlert`,
			Name:              `Arrival Alert`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    true,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `checkInNotes`,
			Name:              `Check In Notes`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    true,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `city`,
			Name:              `City`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            true,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `company`,
			Name:              `Company`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         true,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `country`,
			Name:              `Country`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         true,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `email`,
			Name:              `Email`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           true,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `firstName`,
			Name:              `First Name`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       true,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `ignored`,
			Name:              `Ignored`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         true,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `invitationNotes`,
			Name:              `Invitation Notes`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: true,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `lastName`,
			Name:              `Last Name`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        true,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `other`,
			Name:              `Other`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           true,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `position`,
			Name:              `Position`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        true,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `postalCode`,
			Name:              `Postal Code`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      true,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `quantity`,
			Name:              `Quantity`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        true,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `state`,
			Name:              `State`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           true,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `telephone`,
			Name:              `Telephone`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       true,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `ticketStatus`,
			Name:              `Ticket Status`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    true,
			IsTicketType:      false,
			IsTitle:           false,
		},
		{
			Slug:              `ticketType`,
			Name:              `Access Type`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      true,
			IsTitle:           false,
		},
		{
			Slug:              `title`,
			Name:              `Title`,
			Description:       ``,
			IsAddress1:        false,
			IsAddress2:        false,
			IsArrivalAlert:    false,
			IsCheckInNotes:    false,
			IsCity:            false,
			IsCompany:         false,
			IsCountry:         false,
			IsEmail:           false,
			IsFirstName:       false,
			IsIgnored:         false,
			IsInvitationNotes: false,
			IsLastName:        false,
			IsOther:           false,
			IsPosition:        false,
			IsPostalCode:      false,
			IsQuantity:        false,
			IsState:           false,
			IsTelephone:       false,
			IsTicketStatus:    false,
			IsTicketType:      false,
			IsTitle:           true,
		},
	}
}
