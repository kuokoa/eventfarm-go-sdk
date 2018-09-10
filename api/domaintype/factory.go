package domaintype

type Factory struct {
}

func NewDomainTypeFactory() *Factory {
	return &Factory{}
}

func (f *Factory) GetAllotmentDomainModel() *Allotment {
	return NewAllotment()
}

func (f *Factory) GetAppVersionDomainModel() *AppVersion {
	return NewAppVersion()
}

func (f *Factory) GetBugReportDomainModel() *BugReport {
	return NewBugReport()
}

func (f *Factory) GetCanvasDomainModel() *Canvas {
	return NewCanvas()
}

func (f *Factory) GetDomainMaskDomainModel() *DomainMask {
	return NewDomainMask()
}

func (f *Factory) GetEmailDesignDomainModel() *EmailDesign {
	return NewEmailDesign()
}

func (f *Factory) GetEmailDesignTypeDomainModel() *EmailDesignType {
	return NewEmailDesignType()
}

func (f *Factory) GetEmailMessageDomainModel() *EmailMessage {
	return NewEmailMessage()
}

func (f *Factory) GetEmailNotificationDomainModel() *EmailNotification {
	return NewEmailNotification()
}

func (f *Factory) GetEmailSampleDomainModel() *EmailSample {
	return NewEmailSample()
}

func (f *Factory) GetEmailTemplateDomainModel() *EmailTemplate {
	return NewEmailTemplate()
}

func (f *Factory) GetEventDomainModel() *Event {
	return NewEvent()
}

func (f *Factory) GetFeatureDomainModel() *Feature {
	return NewFeature()
}

func (f *Factory) GetFeatureToggleDomainModel() *FeatureToggle {
	return NewFeatureToggle()
}

func (f *Factory) GetGroupDomainModel() *Group {
	return NewGroup()
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

func (f *Factory) GetInteractionDomainModel() *Interaction {
	return NewInteraction()
}

func (f *Factory) GetInvitationDomainModel() *Invitation {
	return NewInvitation()
}

func (f *Factory) GetOAuthDomainModel() *OAuth {
	return NewOAuth()
}

func (f *Factory) GetPoolDomainModel() *Pool {
	return NewPool()
}

func (f *Factory) GetPoolContactDomainModel() *PoolContact {
	return NewPoolContact()
}

func (f *Factory) GetPoolContractDomainModel() *PoolContract {
	return NewPoolContract()
}

func (f *Factory) GetPoolFeatureDomainModel() *PoolFeature {
	return NewPoolFeature()
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

func (f *Factory) GetSalesforceEventSettingDomainModel() *SalesforceEventSetting {
	return NewSalesforceEventSetting()
}

func (f *Factory) GetSalesforcePoolSettingDomainModel() *SalesforcePoolSetting {
	return NewSalesforcePoolSetting()
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

func (f *Factory) GetTicketBlockDomainModel() *TicketBlock {
	return NewTicketBlock()
}

func (f *Factory) GetTicketTypeDomainModel() *TicketType {
	return NewTicketType()
}

func (f *Factory) GetUserDomainModel() *User {
	return NewUser()
}

func (f *Factory) GetUserAddressDomainModel() *UserAddress {
	return NewUserAddress()
}

func (f *Factory) GetUserAttributeDomainModel() *UserAttribute {
	return NewUserAttribute()
}

func (f *Factory) GetUserIdentifierDomainModel() *UserIdentifier {
	return NewUserIdentifier()
}

func (f *Factory) GetUserNameDomainModel() *UserName {
	return NewUserName()
}
