package elements

// The Contact element specifies a contact in the Unified Contact Store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/contact-contacttype
type ContactContactType struct {
	// The Addresses element specifies an array of Address elements.
	Addresses *AddressesArrayOfAddressesType `xml:"t:Addresses"`
	// The BusinessName element specifies the name of a business.
	BusinessName *BusinessName `xml:"t:BusinessName"`
	// The ContactString element specifies the display name of a contact.
	ContactString *ContactString `xml:"t:ContactString"`
	// The EmailAddresses element specifies an array of extracted email addresses.
	EmailAddresses *EmailAddressesArrayOfExtractedEmailAddresses `xml:"t:EmailAddresses"`
	// The PersonName element specifies the name of an individual found by means of entity extraction.
	PersonName *PersonName `xml:"t:PersonName"`
	// The PhoneNumbers element represents a collection of telephone numbers for a contact.
	PhoneNumbers *PhoneNumbers `xml:"t:PhoneNumbers"`
	// The Urls element specifies an array of URLs for a persona.
	Urls *Urls `xml:"t:Urls"`
}
