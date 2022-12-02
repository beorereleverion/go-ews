package elements

// The PhoneNumbers element represents a collection of telephone numbers for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/phonenumbers
type PhoneNumbers struct {
	// The Entry element represents a telephone number for a contact.
	Entry *EntryPhoneNumber `xml:"t:Entry"`
}
