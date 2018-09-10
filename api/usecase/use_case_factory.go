package usecase

import "gosdk"

type Factory struct {
	restClient gosdk.RestClientInterface
}

func NewUseCaseFactory(restClient gosdk.RestClientInterface) *Factory {
	return &Factory{restClient}
}

func (f *Factory) GetAllotmentDomainModel() *Allotment {
	return NewAllotment(f.restClient)
}

func (f *Factory) GetAppVersionDomainModel() *AppVersion {
	return NewAppVersion(f.restClient)
}

func (f *Factory) GetBugReportDomainModel() *BugReport {
	return NewBugReport(f.restClient)
}

func (f *Factory) GetCanvasDomainModel() *Canvas {
	return NewCanvas(f.restClient)
}

func (f *Factory) GetDomainMaskDomainModel() *DomainMask {
	return NewDomainMask(f.restClient)
}

func (f *Factory) GetEmailDesignDomainModel() *EmailDesign {
	return NewEmailDesign(f.restClient)
}

func (f *Factory) GetEmailDesignTypeDomainModel() *EmailDesignType {
	return NewEmailDesignType(f.restClient)
}

func (f *Factory) GetEmailMessageDomainModel() *EmailMessage {
	return NewEmailMessage(f.restClient)
}

func (f *Factory) GetEmailNotificationDomainModel() *EmailNotification {
	return NewEmailNotification(f.restClient)
}

func (f *Factory) GetEmailSampleDomainModel() *EmailSample {
	return NewEmailSample(f.restClient)
}

func (f *Factory) GetEmailTemplateDomainModel() *EmailTemplate {
	return NewEmailTemplate(f.restClient)
}

func (f *Factory) GetEventDomainModel() *Event {
	return NewEvent(f.restClient)
}

func (f *Factory) GetFeatureDomainModel() *Feature {
	return NewFeature(f.restClient)
}

func (f *Factory) GetFeatureToggleDomainModel() *FeatureToggle {
	return NewFeatureToggle(f.restClient)
}

func (f *Factory) GetGroupDomainModel() *Group {
	return NewGroup(f.restClient)
}

func (f *Factory) GetImportDomainModel() *Import {
	return NewImport(f.restClient)
}

func (f *Factory) GetIntegrationDomainModel() *Integration {
	return NewIntegration(f.restClient)
}

func (f *Factory) GetIntegrationFieldMappingDomainModel() *IntegrationFieldMapping {
	return NewIntegrationFieldMapping(f.restClient)
}

func (f *Factory) GetIntegrationStatusMappingDomainModel() *IntegrationStatusMapping {
	return NewIntegrationStatusMapping(f.restClient)
}

func (f *Factory) GetInteractionDomainModel() *Interaction {
	return NewInteraction(f.restClient)
}

func (f *Factory) GetInvitationDomainModel() *Invitation {
	return NewInvitation(f.restClient)
}

func (f *Factory) GetOAuthDomainModel() *OAuth {
	return NewOAuth(f.restClient)
}

func (f *Factory) GetPoolDomainModel() *Pool {
	return NewPool(f.restClient)
}

func (f *Factory) GetPoolContactDomainModel() *PoolContact {
	return NewPoolContact(f.restClient)
}

func (f *Factory) GetPoolContractDomainModel() *PoolContract {
	return NewPoolContract(f.restClient)
}

func (f *Factory) GetPoolFeatureDomainModel() *PoolFeature {
	return NewPoolFeature(f.restClient)
}

func (f *Factory) GetPromotionDomainModel() *Promotion {
	return NewPromotion(f.restClient)
}

func (f *Factory) GetQueueDomainModel() *Queue {
	return NewQueue(f.restClient)
}

func (f *Factory) GetRegionDomainModel() *Region {
	return NewRegion(f.restClient)
}

func (f *Factory) GetReportDomainModel() *Report {
	return NewReport(f.restClient)
}

func (f *Factory) GetSalesforceDomainModel() *Salesforce {
	return NewSalesforce(f.restClient)
}

func (f *Factory) GetSalesforceEventSettingDomainModel() *SalesforceEventSetting {
	return NewSalesforceEventSetting(f.restClient)
}

func (f *Factory) GetSalesforcePoolSettingDomainModel() *SalesforcePoolSetting {
	return NewSalesforcePoolSetting(f.restClient)
}

func (f *Factory) GetSalutationDomainModel() *Salutation {
	return NewSalutation(f.restClient)
}

func (f *Factory) GetSitePageDomainModel() *SitePage {
	return NewSitePage(f.restClient)
}

func (f *Factory) GetStackDomainModel() *Stack {
	return NewStack(f.restClient)
}

func (f *Factory) GetTicketBlockDomainModel() *TicketBlock {
	return NewTicketBlock(f.restClient)
}

func (f *Factory) GetTicketTypeDomainModel() *TicketType {
	return NewTicketType(f.restClient)
}

func (f *Factory) GetUserDomainModel() *User {
	return NewUser(f.restClient)
}

func (f *Factory) GetUserAddressDomainModel() *UserAddress {
	return NewUserAddress(f.restClient)
}

func (f *Factory) GetUserAttributeDomainModel() *UserAttribute {
	return NewUserAttribute(f.restClient)
}

func (f *Factory) GetUserIdentifierDomainModel() *UserIdentifier {
	return NewUserIdentifier(f.restClient)
}

func (f *Factory) GetUserNameDomainModel() *UserName {
	return NewUserName(f.restClient)
}
