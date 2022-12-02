package elements

// The PhoneNumbers element specifies an array of extracted phone numbers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumbers-arrayofphoneentitiestype
type PhoneNumbersArrayOfPhoneEntitiesType struct {
	// The Phone element specifies a single phone number that results from a phone number entity extraction.
	Phone *PhonePhoneEntityType `xml:"t:Phone"`
}
