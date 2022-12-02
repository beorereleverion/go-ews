package elements

// The GetPhoneCallInformation element specifies a request to get telephone call information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getphonecallinformation
type GetPhoneCallInformation struct {
	// The PhoneCallId element specifies the identifier of a phone call. This element is required.
	PhoneCallId *PhoneCallId `xml:"m:PhoneCallId"`
}
