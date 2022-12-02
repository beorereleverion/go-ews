package elements

// The PhoneNumbers element specifies an array of phone numbers.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumbers-arrayofphonestype
type PhoneNumbersArrayOfPhonesType struct {
	// The Phone element specifies a single phone number that results from a contact entity extraction.
	Phone *Phone `xml:"t:Phone"`
}
