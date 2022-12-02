package elements

// The DisconnectPhoneCall element represents a request to disconnect a call.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/disconnectphonecall
type DisconnectPhoneCall struct {
	// The PhoneCallId element specifies the identifier of a phone call. This element is required.
	PhoneCallId *PhoneCallId `xml:"m:PhoneCallId"`
}
