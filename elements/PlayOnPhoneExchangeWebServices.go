package elements

// The PlayOnPhone element represents a request to read an item on a phone.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/playonphone-exchange-web-services
type PlayOnPhoneExchangeWebServices struct {
	// The DialString element represents the dial string of the telephone number that is called to play an item by telephone. This element is required.
	DialString *DialStringExchangeWebServices `xml:"m:DialString"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
}
