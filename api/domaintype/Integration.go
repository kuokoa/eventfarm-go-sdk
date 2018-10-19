package domaintype

type Integration struct {
}

func NewIntegration() *Integration {
	return &Integration{}
}

type IntegrationInvitationCreationType struct {
	Slug                 string
	Name                 string
	Description          string
	IsConfirmedNoEmail   bool
	IsUnconfirmedNoEmail bool
}

type IntegrationSettingType struct {
	Slug         string
	Name         string
	Description  string
	IsSalesforce bool
	IsMarketo    bool
}

type UserAttributeInfoFieldType struct {
	Slug        string
	Name        string
	Description string
}

func (f *Integration) ListIntegrationInvitationCreationTypes() []IntegrationInvitationCreationType {
	return []IntegrationInvitationCreationType{
		{
			Slug:                 `unconfirmed-no-email`,
			Name:                 `Unconfirmed No Email`,
			Description:          `Add as Unconfirmed, email will not be sent.`,
			IsConfirmedNoEmail:   false,
			IsUnconfirmedNoEmail: true,
		},
		{
			Slug:                 `confirmed-no-email`,
			Name:                 `Confirmed No Email`,
			Description:          `Add as Confirmed, email will not be sent.`,
			IsConfirmedNoEmail:   true,
			IsUnconfirmedNoEmail: false,
		},
	}
}

func (f *Integration) ListIntegrationSettingTypes() []IntegrationSettingType {
	return []IntegrationSettingType{
		{
			Slug:         `salesforce`,
			Name:         `Salesforce`,
			Description:  ``,
			IsSalesforce: true,
			IsMarketo:    false,
		},
		{
			Slug:         `marketo`,
			Name:         `Marketo`,
			Description:  ``,
			IsSalesforce: false,
			IsMarketo:    true,
		},
	}
}

func (f *Integration) ListUserAttributeInfoFieldTypes() []UserAttributeInfoFieldType {
	return []UserAttributeInfoFieldType{
		{
			Slug:        `company`,
			Name:        `Company`,
			Description: ``,
		},
		{
			Slug:        `position`,
			Name:        `Position`,
			Description: ``,
		},
		{
			Slug:        `title`,
			Name:        `Title`,
			Description: ``,
		},
		{
			Slug:        `telephone`,
			Name:        `Telephone`,
			Description: ``,
		},
	}
}
