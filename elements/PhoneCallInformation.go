package elements

// The PhoneCallInformation element specifies the state information for a phone call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonecallinformation
import "encoding/xml"

type PhoneCallInformation struct {
	XMLName xml.Name

	// The ConnectionFailureCause element specifies the reason for a disconnection from a telephone call.
	ConnectionFailureCause *ConnectionFailureCause `xml:"ConnectionFailureCause"`
	// The PhoneCallState element specifies the current state for a phone call.
	PhoneCallState *PhoneCallState `xml:"PhoneCallState"`
	// The SIPResponseCode element specifies the SIP response code.
	SIPResponseCode *SIPResponseCode `xml:"SIPResponseCode"`
	// The SIPResponseText element specifies the SIP response text.
	SIPResponseText *SIPResponseText `xml:"SIPResponseText"`
}

func (P *PhoneCallInformation) SetForMarshal() {
	P.XMLName.Local = "m:PhoneCallInformation"
}

func (P *PhoneCallInformation) GetSchema() *Schema {
	return &SchemaMessages
}
