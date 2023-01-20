package elements

// The SendSMSAlertToRecipients element indicates the mobile phone numbers to which a Short Message Service (SMS) alert is to be sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sendsmsalerttorecipients
import "encoding/xml"

type SendSMSAlertToRecipients struct {
	XMLName xml.Name

	// The Address element represents a fully resolved e-mail address.
	Address *AddressEmailAddressType `xml:"Address"`
}

func (S *SendSMSAlertToRecipients) SetForMarshal() {
	S.XMLName.Local = "m:SendSMSAlertToRecipients"
}

func (S *SendSMSAlertToRecipients) GetSchema() *Schema {
	return &SchemaMessages
}
