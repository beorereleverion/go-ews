package elements

// The EntityExtractionResult element specifies the EntityExtractionResult property of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/entityextractionresult
type EntityExtractionResult struct {
	// The Addresses element specifies an array of AddressEntity elements.
	Addresses *AddressesArrayOfAddressEntitiesType `xml:"t:Addresses"`
	// The Contacts element specifies an array of contacts.
	Contacts *ContactsArrayOfContactsType `xml:"t:Contacts"`
	// The EmailAddresses element specifies an array of email address entities.
	EmailAddresses *EmailAddressesArrayOfEmailAddressEntitiesType `xml:"t:EmailAddresses"`
	// The MeetingSuggestions element specifies an array of MeetingSuggestion elements that contain entity extraction results.
	MeetingSuggestions *MeetingSuggestions `xml:"t:MeetingSuggestions"`
	// The PhoneNumbers element specifies an array of extracted phone numbers.
	PhoneNumbers *PhoneNumbersArrayOfPhoneEntitiesType `xml:"t:PhoneNumbers"`
	// The TaskSuggestions element specifies an array of task suggestions extracted from an item.
	TaskSuggestions *TaskSuggestions `xml:"t:TaskSuggestions"`
	// The Urls element specifies an array of URLs that are the result of entity extraction from an item in the mailbox.
	Urls *UrlsArrayOfUrlEntitiesType `xml:"t:Urls"`
}
