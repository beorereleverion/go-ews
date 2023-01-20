package elements

// The FindMessageTrackingReport element specifies criteria for the types of messages to find.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findmessagetrackingreport
import "encoding/xml"

type FindMessageTrackingReport struct {
	XMLName xml.Name

	// The DiagnosticsLevel element represents timing and performance information that will be used to derive the report.
	DiagnosticsLevel *DiagnosticsLevel `xml:"DiagnosticsLevel"`
	// The Domain element represents the domain to search for.
	Domain *DomainMessageTracking `xml:"Domain"`
	// The EndDateTime element specifies the end date and time for a rule or a search.
	EndDateTime *EndDateTime `xml:"EndDateTime"`
	// The FederatedDeliveryMailbox element represents the mailbox to which a cross-premise message was sent.
	FederatedDeliveryMailbox *FederatedDeliveryMailbox `xml:"FederatedDeliveryMailbox"`
	// The MessageId element represents the message identification to search for.
	MessageId *MessageId `xml:"MessageId"`
	// The Properties element contains a list of one or more tracking properties.
	Properties *PropertiesArrayOfTrackingPropertiesType `xml:"Properties"`
	// The PurportedSender element contains contact information for the alleged sender of an e-mail message.
	PurportedSender *PurportedSender `xml:"PurportedSender"`
	// The Recipient element represents the recipient for whom the event occurred.
	Recipient *Recipient `xml:"Recipient"`
	// The Scope element specifies the scope of the message tracking report.
	Scope *ScopeNonEmptyStringType `xml:"Scope"`
	// The Sender element represents the e-mail address for the sender of the message.
	Sender *SenderEmailAddressType `xml:"Sender"`
	// The ServerHint element represents the starting point for tracking a message in a remote site or forest.
	ServerHint *ServerHint `xml:"ServerHint"`
	// The StartDateTime element specifies the start date and time for a rule or a search.
	StartDateTime *StartDateTime `xml:"StartDateTime"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"Subject"`
}

func (F *FindMessageTrackingReport) SetForMarshal() {
	F.XMLName.Local = "m:FindMessageTrackingReport"
}

func (F *FindMessageTrackingReport) GetSchema() *Schema {
	return &SchemaMessages
}
