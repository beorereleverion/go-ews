package elements

// The PhoneCallInformation element specifies the state information for a phone call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonecallinformation
type PhoneCallInformation struct {
	// The ConnectionFailureCause element specifies the reason for a disconnection from a telephone call.
	ConnectionFailureCause *ConnectionFailureCause `xml:"t:ConnectionFailureCause"`
	// The PhoneCallState element specifies the current state for a phone call.
	PhoneCallState *PhoneCallState `xml:"t:PhoneCallState"`
	// The SIPResponseCode element specifies the SIP response code.
	SIPResponseCode *SIPResponseCode `xml:"t:SIPResponseCode"`
	// The SIPResponseText element specifies the SIP response text.
	SIPResponseText *SIPResponseText `xml:"t:SIPResponseText"`
}
