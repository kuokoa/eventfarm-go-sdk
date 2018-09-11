package domaintype

type Factory struct {
}

func NewDomainTypeFactory() *Factory {
	return &Factory{}
}

func (f *Factory) GetAppVersionDomainModel() *AppVersion {
	return NewAppVersion()
}

func (f *Factory) GetEmailDesignDomainModel() *EmailDesign {
	return NewEmailDesign()
}

func (f *Factory) GetEmailMessageDomainModel() *EmailMessage {
	return NewEmailMessage()
}

func (f *Factory) GetEmailTemplateDomainModel() *EmailTemplate {
	return NewEmailTemplate()
}

func (f *Factory) GetEventDomainModel() *Event {
	return NewEvent()
}

func (f *Factory) GetImportDomainModel() *Import {
	return NewImport()
}

func (f *Factory) GetIntegrationDomainModel() *Integration {
	return NewIntegration()
}

func (f *Factory) GetIntegrationFieldMappingDomainModel() *IntegrationFieldMapping {
	return NewIntegrationFieldMapping()
}

func (f *Factory) GetIntegrationStatusMappingDomainModel() *IntegrationStatusMapping {
	return NewIntegrationStatusMapping()
}

func (f *Factory) GetInvitationDomainModel() *Invitation {
	return NewInvitation()
}

func (f *Factory) GetPoolDomainModel() *Pool {
	return NewPool()
}

func (f *Factory) GetPromotionDomainModel() *Promotion {
	return NewPromotion()
}

func (f *Factory) GetQueueDomainModel() *Queue {
	return NewQueue()
}

func (f *Factory) GetRegionDomainModel() *Region {
	return NewRegion()
}

func (f *Factory) GetReportDomainModel() *Report {
	return NewReport()
}

func (f *Factory) GetSalesforceDomainModel() *Salesforce {
	return NewSalesforce()
}

func (f *Factory) GetSalutationDomainModel() *Salutation {
	return NewSalutation()
}

func (f *Factory) GetSitePageDomainModel() *SitePage {
	return NewSitePage()
}

func (f *Factory) GetStackDomainModel() *Stack {
	return NewStack()
}

func (f *Factory) GetUserDomainModel() *User {
	return NewUser()
}
