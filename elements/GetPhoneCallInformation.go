package elements

// The GetPhoneCallInformation element specifies a request to get telephone call information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getphonecallinformation
import "encoding/xml"

type GetPhoneCallInformation struct {
	XMLName xml.Name

	// The PhoneCallId element specifies the identifier of a phone call. This element is required.
	PhoneCallId *PhoneCallId `xml:"PhoneCallId"`
}

func (G *GetPhoneCallInformation) SetForMarshal() {
	G.XMLName.Local = "m:GetPhoneCallInformation"
}

func (G *GetPhoneCallInformation) GetSchema() *Schema {
	return &SchemaMessages
}
