/**
 *  This file was auto generated, please do not edit it directly.
**/

package usecase

import "github.com/eventfarm/go-sdk/rest"

type Factory struct {
	restClient rest.RestClientInterface
}

func NewUseCaseFactory(restClient rest.RestClientInterface) *Factory {
	return &Factory{restClient}
}

func (f *Factory) ActivityLog() *ActivityLog {
	return NewActivityLog(f.restClient)
}

func (f *Factory) Allotment() *Allotment {
	return NewAllotment(f.restClient)
}

func (f *Factory) AppVersion() *AppVersion {
	return NewAppVersion(f.restClient)
}

func (f *Factory) BugReport() *BugReport {
	return NewBugReport(f.restClient)
}

func (f *Factory) Canvas() *Canvas {
	return NewCanvas(f.restClient)
}

func (f *Factory) DomainMask() *DomainMask {
	return NewDomainMask(f.restClient)
}

func (f *Factory) EFx() *EFx {
	return NewEFx(f.restClient)
}

func (f *Factory) EmailDesign() *EmailDesign {
	return NewEmailDesign(f.restClient)
}

func (f *Factory) EmailDesignType() *EmailDesignType {
	return NewEmailDesignType(f.restClient)
}

func (f *Factory) EmailMessage() *EmailMessage {
	return NewEmailMessage(f.restClient)
}

func (f *Factory) EmailNotification() *EmailNotification {
	return NewEmailNotification(f.restClient)
}

func (f *Factory) EmailSample() *EmailSample {
	return NewEmailSample(f.restClient)
}

func (f *Factory) EmailTemplate() *EmailTemplate {
	return NewEmailTemplate(f.restClient)
}

func (f *Factory) Event() *Event {
	return NewEvent(f.restClient)
}

func (f *Factory) EventTheme() *EventTheme {
	return NewEventTheme(f.restClient)
}

func (f *Factory) Feature() *Feature {
	return NewFeature(f.restClient)
}

func (f *Factory) FeatureToggle() *FeatureToggle {
	return NewFeatureToggle(f.restClient)
}

func (f *Factory) Group() *Group {
	return NewGroup(f.restClient)
}

func (f *Factory) Import() *Import {
	return NewImport(f.restClient)
}

func (f *Factory) Integration() *Integration {
	return NewIntegration(f.restClient)
}

func (f *Factory) IntegrationFieldMapping() *IntegrationFieldMapping {
	return NewIntegrationFieldMapping(f.restClient)
}

func (f *Factory) IntegrationStatusMapping() *IntegrationStatusMapping {
	return NewIntegrationStatusMapping(f.restClient)
}

func (f *Factory) Invitation() *Invitation {
	return NewInvitation(f.restClient)
}

func (f *Factory) OAuth() *OAuth {
	return NewOAuth(f.restClient)
}

func (f *Factory) OAuth2() *OAuth2 {
	return NewOAuth2(f.restClient)
}

func (f *Factory) Payment() *Payment {
	return NewPayment(f.restClient)
}

func (f *Factory) Pool() *Pool {
	return NewPool(f.restClient)
}

func (f *Factory) PoolContact() *PoolContact {
	return NewPoolContact(f.restClient)
}

func (f *Factory) PoolContract() *PoolContract {
	return NewPoolContract(f.restClient)
}

func (f *Factory) PoolFeature() *PoolFeature {
	return NewPoolFeature(f.restClient)
}

func (f *Factory) Promotion() *Promotion {
	return NewPromotion(f.restClient)
}

func (f *Factory) Question() *Question {
	return NewQuestion(f.restClient)
}

func (f *Factory) Queue() *Queue {
	return NewQueue(f.restClient)
}

func (f *Factory) Refund() *Refund {
	return NewRefund(f.restClient)
}

func (f *Factory) Region() *Region {
	return NewRegion(f.restClient)
}

func (f *Factory) Report() *Report {
	return NewReport(f.restClient)
}

func (f *Factory) Salesforce() *Salesforce {
	return NewSalesforce(f.restClient)
}

func (f *Factory) SalesforceEventSetting() *SalesforceEventSetting {
	return NewSalesforceEventSetting(f.restClient)
}

func (f *Factory) SalesforcePoolSetting() *SalesforcePoolSetting {
	return NewSalesforcePoolSetting(f.restClient)
}

func (f *Factory) Salutation() *Salutation {
	return NewSalutation(f.restClient)
}

func (f *Factory) SitePage() *SitePage {
	return NewSitePage(f.restClient)
}

func (f *Factory) Stack() *Stack {
	return NewStack(f.restClient)
}

func (f *Factory) TicketBlock() *TicketBlock {
	return NewTicketBlock(f.restClient)
}

func (f *Factory) TicketType() *TicketType {
	return NewTicketType(f.restClient)
}

func (f *Factory) Transaction() *Transaction {
	return NewTransaction(f.restClient)
}

func (f *Factory) Transfer() *Transfer {
	return NewTransfer(f.restClient)
}

func (f *Factory) User() *User {
	return NewUser(f.restClient)
}

func (f *Factory) UserAddress() *UserAddress {
	return NewUserAddress(f.restClient)
}

func (f *Factory) UserAttribute() *UserAttribute {
	return NewUserAttribute(f.restClient)
}

func (f *Factory) UserIdentifier() *UserIdentifier {
	return NewUserIdentifier(f.restClient)
}

func (f *Factory) UserName() *UserName {
	return NewUserName(f.restClient)
}

func (f *Factory) Withdrawal() *Withdrawal {
	return NewWithdrawal(f.restClient)
}
