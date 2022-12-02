package elements

// The PhoneCallId element specifies the identifier of a phone call. This element is required.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonecallid
type PhoneCallId struct {
	// Identifies the phone call to disconnect. This attribute is required.
	Id *string `xml:"Id,attr"`
}
