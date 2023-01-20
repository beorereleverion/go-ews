package elements

// The PhoneCallId element specifies the identifier of a phone call. This element is required.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonecallid
import "encoding/xml"

type PhoneCallId struct {
	XMLName xml.Name

	// Identifies the phone call to disconnect. This attribute is required.
	Id *string `xml:"Id,attr"`
}

func (P *PhoneCallId) SetForMarshal() {
	P.XMLName.Local = "m:PhoneCallId"
}

func (P *PhoneCallId) GetSchema() *Schema {
	return &SchemaMessages
}
