package elements

// The EntityExtractionResult element specifies the EntityExtractionResult property of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/entityextractionresult
import "encoding/xml"

type EntityExtractionResult struct {
	XMLName xml.Name

	// The Addresses element specifies an array of AddressEntity elements.
	Addresses *AddressesArrayOfAddressEntitiesType `xml:"Addresses"`
	// The Contacts element specifies an array of contacts.
	Contacts *ContactsArrayOfContactsType `xml:"Contacts"`
	// The EmailAddresses element specifies an array of email address entities.
	EmailAddresses *EmailAddressesArrayOfEmailAddressEntitiesType `xml:"EmailAddresses"`
	// The MeetingSuggestions element specifies an array of MeetingSuggestion elements that contain entity extraction results.
	MeetingSuggestions *MeetingSuggestions `xml:"MeetingSuggestions"`
	// The PhoneNumbers element specifies an array of extracted phone numbers.
	PhoneNumbers *PhoneNumbersArrayOfPhoneEntitiesType `xml:"PhoneNumbers"`
	// The TaskSuggestions element specifies an array of task suggestions extracted from an item.
	TaskSuggestions *TaskSuggestions `xml:"TaskSuggestions"`
	// The Urls element specifies an array of URLs that are the result of entity extraction from an item in the mailbox.
	Urls *UrlsArrayOfUrlEntitiesType `xml:"Urls"`
}

func (E *EntityExtractionResult) SetForMarshal() {
	E.XMLName.Local = "t:EntityExtractionResult"
}

func (E *EntityExtractionResult) GetSchema() *Schema {
	return &SchemaTypes
}
