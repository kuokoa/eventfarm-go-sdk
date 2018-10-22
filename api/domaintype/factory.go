/**
 *  This file was auto generated, please do not edit it directly.
**/

package domaintype

type Factory struct {
}

func NewDomainTypeFactory() *Factory {
	return &Factory{}
}

func (f *Factory) AppVersion() *AppVersion {
	return NewAppVersion()
}

func (f *Factory) EmailDesign() *EmailDesign {
	return NewEmailDesign()
}

func (f *Factory) EmailMessage() *EmailMessage {
	return NewEmailMessage()
}

func (f *Factory) EmailTemplate() *EmailTemplate {
	return NewEmailTemplate()
}

func (f *Factory) Event() *Event {
	return NewEvent()
}

func (f *Factory) Import() *Import {
	return NewImport()
}

func (f *Factory) Integration() *Integration {
	return NewIntegration()
}

func (f *Factory) IntegrationFieldMapping() *IntegrationFieldMapping {
	return NewIntegrationFieldMapping()
}

func (f *Factory) IntegrationStatusMapping() *IntegrationStatusMapping {
	return NewIntegrationStatusMapping()
}

func (f *Factory) Invitation() *Invitation {
	return NewInvitation()
}

func (f *Factory) Pool() *Pool {
	return NewPool()
}

func (f *Factory) Promotion() *Promotion {
	return NewPromotion()
}

func (f *Factory) Queue() *Queue {
	return NewQueue()
}

func (f *Factory) Region() *Region {
	return NewRegion()
}

func (f *Factory) Report() *Report {
	return NewReport()
}

func (f *Factory) Salesforce() *Salesforce {
	return NewSalesforce()
}

func (f *Factory) Salutation() *Salutation {
	return NewSalutation()
}

func (f *Factory) SitePage() *SitePage {
	return NewSitePage()
}

func (f *Factory) Stack() *Stack {
	return NewStack()
}

func (f *Factory) User() *User {
	return NewUser()
}
