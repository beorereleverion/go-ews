package elements

// The IndexedFieldURI element identifies individual members of a dictionary.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/indexedfielduri
type IndexedFieldURI struct {
	// Identifies the dictionary that contains the member to return. This attribute is required.
	FieldURI *IndexedFieldAttrURIFieldURI `xml:"FieldURI,attr"`
	// Identifies the member of the dictionary to return. This attribute is required.
	FieldIndex *string `xml:"FieldIndex,attr"`
}

type IndexedFieldAttrURIFieldURI string

const (
	// Represents the message header of an item.
	IndexedFieldAttrURIFieldURIitemInternetMessageHeader IndexedFieldAttrURIFieldURI = `item:InternetMessageHeader`
	// Represents the instant messaging address of a contact.
	IndexedFieldAttrURIFieldURIcontactsImAddress IndexedFieldAttrURIFieldURI = `contacts:ImAddress`
	// Represents the street address of a contact.
	IndexedFieldAttrURIFieldURIcontactsPhysicalAddressStreet IndexedFieldAttrURIFieldURI = `contacts:PhysicalAddress:Street`
	// Represents the city of a contact.
	IndexedFieldAttrURIFieldURIcontactsPhysicalAddressCity IndexedFieldAttrURIFieldURI = `contacts:PhysicalAddress:City`
	// Represents the state of a contact.
	IndexedFieldAttrURIFieldURIcontactsPhysicalAddressState IndexedFieldAttrURIFieldURI = `contacts:PhysicalAddress:State`
	// Represents the country/region of a contact.
	IndexedFieldAttrURIFieldURIcontactsPhysicalAddressCountry IndexedFieldAttrURIFieldURI = `contacts:PhysicalAddress:Country`
	// Represents the postal code of a contact.
	IndexedFieldAttrURIFieldURIcontactsPhysicalAddressPostalCode IndexedFieldAttrURIFieldURI = `contacts:PhysicalAddress:PostalCode`
	// Represents the phone number of a contact.
	IndexedFieldAttrURIFieldURIcontactsPhoneNumber IndexedFieldAttrURIFieldURI = `contacts:PhoneNumber`
	// Represents the e-mail address of a contact.
	IndexedFieldAttrURIFieldURIcontactsEmailAddress IndexedFieldAttrURIFieldURI = `contacts:EmailAddress`
	// Represents a member of a distribution list.
	IndexedFieldAttrURIFieldURIdistributionlistMembersMember IndexedFieldAttrURIFieldURI = `distributionlist:Members:Member`
)
