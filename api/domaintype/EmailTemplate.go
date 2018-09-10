package domaintype

type EmailTemplate struct {
}

func NewEmailTemplate() *EmailTemplate {
	return &EmailTemplate{}
}

type EmailTemplateType struct {
	Slug               string
	Name               string
	Description        string
	IsFullResponsive   bool
	IsMobileResponsive bool
	IsDefaultInvite    bool
	IsMobileColumns    bool
	IsFullWidthHeader  bool
}

func (f *EmailTemplate) GetEmailTemplateType() []EmailTemplateType {
	return []EmailTemplateType{
		{
			Slug:               `simple-template`,
			Name:               `Simple Template`,
			Description:        `This is simple template`,
			IsFullResponsive:   true,
			IsMobileResponsive: false,
			IsDefaultInvite:    false,
			IsMobileColumns:    false,
			IsFullWidthHeader:  false,
		},
		{
			Slug:               `simple-header`,
			Name:               `Simple Header`,
			Description:        `This is simple header`,
			IsFullResponsive:   false,
			IsMobileResponsive: false,
			IsDefaultInvite:    false,
			IsMobileColumns:    true,
			IsFullWidthHeader:  false,
		},
		{
			Slug:               `simple-template-border`,
			Name:               `Simple Template Border`,
			Description:        `This is simple template border`,
			IsFullResponsive:   false,
			IsMobileResponsive: true,
			IsDefaultInvite:    false,
			IsMobileColumns:    false,
			IsFullWidthHeader:  false,
		},
		{
			Slug:               `default-invite`,
			Name:               `Default Invite`,
			Description:        `This is default invite`,
			IsFullResponsive:   false,
			IsMobileResponsive: false,
			IsDefaultInvite:    true,
			IsMobileColumns:    false,
			IsFullWidthHeader:  false,
		},
		{
			Slug:               `full-width-header`,
			Name:               `Full Width Header`,
			Description:        `This is full width header`,
			IsFullResponsive:   false,
			IsMobileResponsive: false,
			IsDefaultInvite:    false,
			IsMobileColumns:    false,
			IsFullWidthHeader:  true,
		},
	}
}
