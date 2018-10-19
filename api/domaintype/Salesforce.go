package domaintype

type Salesforce struct {
}

func NewSalesforce() *Salesforce {
	return &Salesforce{}
}

type CampaignMemberType struct {
	Slug        string
	Name        string
	Description string
	IsContact   bool
	IsLead      bool
}

type NewContactRuleType struct {
	Slug            string
	Name            string
	Description     string
	IsDoNothing     bool
	IsCreateContact bool
	IsCreateLead    bool
}

func (f *Salesforce) ListCampaignMemberTypes() []CampaignMemberType {
	return []CampaignMemberType{
		{
			Slug:        `contact`,
			Name:        `Contact`,
			Description: ``,
			IsContact:   true,
			IsLead:      false,
		},
		{
			Slug:        `lead`,
			Name:        `Lead`,
			Description: ``,
			IsContact:   false,
			IsLead:      true,
		},
	}
}

func (f *Salesforce) ListNewContactRuleTypes() []NewContactRuleType {
	return []NewContactRuleType{
		{
			Slug:            `do-nothing`,
			Name:            `Do Nothing`,
			Description:     `Do nothing. They will remain only in Event Farm.`,
			IsDoNothing:     true,
			IsCreateContact: false,
			IsCreateLead:    false,
		},
		{
			Slug:            `create-contact`,
			Name:            `Create Contact`,
			Description:     `Create a new Contact in Salesforce.`,
			IsDoNothing:     false,
			IsCreateContact: true,
			IsCreateLead:    false,
		},
		{
			Slug:            `create-lead`,
			Name:            `Create Lead`,
			Description:     `Create a new Lead in Salesforce.`,
			IsDoNothing:     false,
			IsCreateContact: false,
			IsCreateLead:    true,
		},
	}
}
