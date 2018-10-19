package domaintype

type Factory struct {
}

func NewDomainTypeFactory() *Factory {
	return &Factory{}
}

func (f *Factory) Allotment() *Allotment {
	return NewAllotment()
}

func (f *Factory) AppVersion() *AppVersion {
	return NewAppVersion()
}

func (f *Factory) BugReport() *BugReport {
	return NewBugReport()
}

func (f *Factory) Canvas() *Canvas {
	return NewCanvas()
}

func (f *Factory) DomainMask() *DomainMask {
	return NewDomainMask()
}

func (f *Factory) EmailDesign() *EmailDesign {
	return NewEmailDesign()
}

func (f *Factory) EmailDesignType() *EmailDesignType {
	return NewEmailDesignType()
}

func (f *Factory) EmailMessage() *EmailMessage {
	return NewEmailMessage()
}

func (f *Factory) EmailNotification() *EmailNotification {
	return NewEmailNotification()
}

func (f *Factory) EmailSample() *EmailSample {
	return NewEmailSample()
}

func (f *Factory) EmailTemplate() *EmailTemplate {
	return NewEmailTemplate()
}

func (f *Factory) Event() *Event {
	return NewEvent()
}

func (f *Factory) Feature() *Feature {
	return NewFeature()
}

func (f *Factory) FeatureToggle() *FeatureToggle {
	return NewFeatureToggle()
}

func (f *Factory) Group() *Group {
	return NewGroup()
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

func (f *Factory) Interaction() *Interaction {
	return NewInteraction()
}

func (f *Factory) Invitation() *Invitation {
	return NewInvitation()
}

func (f *Factory) OAuth() *OAuth {
	return NewOAuth()
}

func (f *Factory) OAuth2() *OAuth2 {
	return NewOAuth2()
}

func (f *Factory) Pool() *Pool {
	return NewPool()
}

func (f *Factory) PoolContact() *PoolContact {
	return NewPoolContact()
}

func (f *Factory) PoolContract() *PoolContract {
	return NewPoolContract()
}

func (f *Factory) PoolFeature() *PoolFeature {
	return NewPoolFeature()
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

func (f *Factory) SalesforceEventSetting() *SalesforceEventSetting {
	return NewSalesforceEventSetting()
}

func (f *Factory) SalesforcePoolSetting() *SalesforcePoolSetting {
	return NewSalesforcePoolSetting()
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

func (f *Factory) TicketBlock() *TicketBlock {
	return NewTicketBlock()
}

func (f *Factory) TicketType() *TicketType {
	return NewTicketType()
}

func (f *Factory) User() *User {
	return NewUser()
}

func (f *Factory) UserAddress() *UserAddress {
	return NewUserAddress()
}

func (f *Factory) UserAttribute() *UserAttribute {
	return NewUserAttribute()
}

func (f *Factory) UserIdentifier() *UserIdentifier {
	return NewUserIdentifier()
}

func (f *Factory) UserName() *UserName {
	return NewUserName()
}
