package elements

// The ForwardAsAttachmentToRecipients element indicates the e-mail addresses to which messages are to be forwarded as attachments.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/forwardasattachmenttorecipients
import "encoding/xml"

type ForwardAsAttachmentToRecipients struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (F *ForwardAsAttachmentToRecipients) SetForMarshal() {
	F.XMLName.Local = "m:ForwardAsAttachmentToRecipients"
}

func (F *ForwardAsAttachmentToRecipients) GetSchema() *Schema {
	return &SchemaMessages
}
