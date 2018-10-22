/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type IntegrationFieldMapping struct {
}

func NewIntegrationFieldMapping() *IntegrationFieldMapping {
	return &IntegrationFieldMapping{}
}

type FieldMappingType struct {
	Slug                      string
	Name                      string
	Description               string
	IsSalesforceImportLead    bool
	IsSalesforceImportContact bool
	IsSalesforceExportLead    bool
	IsSalesforceExportContact bool
	IsMarketoFieldMapping     bool
}

type FieldType struct {
	Slug             string
	Name             string
	Description      string
	IsQuestion       bool
	IsUserAttribute  bool
	IsUserIdentifier bool
	IsUserName       bool
}

type UpdateRuleType struct {
	Slug        string
	Name        string
	Description string
	IsNever     bool
	IsDifferent bool
	IsBlank     bool
}

func (f *IntegrationFieldMapping) ListFieldMappingTypes() []FieldMappingType {
	return []FieldMappingType{
		{
			Slug:                      `salesforce-import-lead`,
			Name:                      `Salesforce Import Lead`,
			Description:               ``,
			IsSalesforceImportLead:    true,
			IsSalesforceImportContact: false,
			IsSalesforceExportLead:    false,
			IsSalesforceExportContact: false,
			IsMarketoFieldMapping:     false,
		},
		{
			Slug:                      `salesforce-import-contact`,
			Name:                      `Salesforce Import Contact`,
			Description:               ``,
			IsSalesforceImportLead:    false,
			IsSalesforceImportContact: false,
			IsSalesforceExportLead:    false,
			IsSalesforceExportContact: false,
			IsMarketoFieldMapping:     false,
		},
		{
			Slug:                      `salesforce-export-lead`,
			Name:                      `Salesforce Export Lead`,
			Description:               ``,
			IsSalesforceImportLead:    false,
			IsSalesforceImportContact: false,
			IsSalesforceExportLead:    true,
			IsSalesforceExportContact: false,
			IsMarketoFieldMapping:     false,
		},
		{
			Slug:                      `salesforce-export-contact`,
			Name:                      `Salesforce Export Contact`,
			Description:               ``,
			IsSalesforceImportLead:    false,
			IsSalesforceImportContact: true,
			IsSalesforceExportLead:    false,
			IsSalesforceExportContact: true,
			IsMarketoFieldMapping:     false,
		},
		{
			Slug:                      `marketo-field-mapping`,
			Name:                      `Marketo Field Mapping`,
			Description:               ``,
			IsSalesforceImportLead:    false,
			IsSalesforceImportContact: false,
			IsSalesforceExportLead:    false,
			IsSalesforceExportContact: false,
			IsMarketoFieldMapping:     true,
		},
	}
}

func (f *IntegrationFieldMapping) ListFieldTypes() []FieldType {
	return []FieldType{
		{
			Slug:             `question`,
			Name:             `Question`,
			Description:      ``,
			IsQuestion:       true,
			IsUserAttribute:  false,
			IsUserIdentifier: false,
			IsUserName:       false,
		},
		{
			Slug:             `user-attribute`,
			Name:             `User Attribute`,
			Description:      ``,
			IsQuestion:       false,
			IsUserAttribute:  true,
			IsUserIdentifier: false,
			IsUserName:       false,
		},
		{
			Slug:             `user-name`,
			Name:             `User Name`,
			Description:      ``,
			IsQuestion:       false,
			IsUserAttribute:  false,
			IsUserIdentifier: false,
			IsUserName:       true,
		},
		{
			Slug:             `user-identifier`,
			Name:             `User Identifier`,
			Description:      ``,
			IsQuestion:       false,
			IsUserAttribute:  false,
			IsUserIdentifier: true,
			IsUserName:       false,
		},
	}
}

func (f *IntegrationFieldMapping) ListUpdateRuleTypes() []UpdateRuleType {
	return []UpdateRuleType{
		{
			Slug:        `never`,
			Name:        `Never`,
			Description: ``,
			IsNever:     true,
			IsDifferent: false,
			IsBlank:     false,
		},
		{
			Slug:        `different`,
			Name:        `Only If Different`,
			Description: ``,
			IsNever:     false,
			IsDifferent: true,
			IsBlank:     false,
		},
		{
			Slug:        `blank`,
			Name:        `Only If Blank`,
			Description: ``,
			IsNever:     false,
			IsDifferent: false,
			IsBlank:     true,
		},
	}
}
