package elements

// The WorkFaxes element specifies an array of work fax numbers and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/workfaxes
type WorkFaxes struct {
	// The PhoneNumberAttributedValue element specifies an instance of an array of phone numbers and their associated attributions.
	PhoneNumberAttributedValue *PhoneNumberAttributedValue `xml:"t:PhoneNumberAttributedValue"`
}
