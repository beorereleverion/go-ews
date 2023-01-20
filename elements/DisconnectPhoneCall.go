package elements

// The DisconnectPhoneCall element represents a request to disconnect a call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disconnectphonecall
import "encoding/xml"

type DisconnectPhoneCall struct {
	XMLName xml.Name

	// The PhoneCallId element specifies the identifier of a phone call. This element is required.
	PhoneCallId *PhoneCallId `xml:"PhoneCallId"`
}

func (D *DisconnectPhoneCall) SetForMarshal() {
	D.XMLName.Local = "m:DisconnectPhoneCall"
}

func (D *DisconnectPhoneCall) GetSchema() *Schema {
	return &SchemaMessages
}
