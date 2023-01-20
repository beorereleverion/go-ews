package elements

// The CompleteName element represents the complete name of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/completename
import "encoding/xml"

type CompleteName struct {
	XMLName xml.Name

	// The FirstName element represents the first name of a contact.
	FirstName *FirstName `xml:"FirstName"`
	// The FullName element represents the full name of a contact.
	FullName *FullName `xml:"FullName"`
	// The Initials element represents the initials of a contact.
	Initials *Initials `xml:"Initials"`
	// The LastName element represents the last name of a contact.
	LastName *LastName `xml:"LastName"`
	// The MiddleName element represents the middle name of a contact.
	MiddleName *MiddleName `xml:"MiddleName"`
	// The Nickname element represents the nickname of a contact.
	Nickname *Nickname `xml:"Nickname"`
	// The Suffix element represents a suffix to a contact's name.
	Suffix *Suffix `xml:"Suffix"`
	// The Title element represents the title of a contact.
	Title *Title `xml:"Title"`
	// The YomiFirstName element represents the name that is used in Japan for the searchable or phonetic spelling for a Japanese first name.
	YomiFirstName *YomiFirstName `xml:"YomiFirstName"`
	// The YomiLastName element represents the name that is used in Japan for the searchable or phonetic spelling for a Japanese last name.
	YomiLastName *YomiLastName `xml:"YomiLastName"`
}

func (C *CompleteName) SetForMarshal() {
	C.XMLName.Local = "t:CompleteName"
}

func (C *CompleteName) GetSchema() *Schema {
	return &SchemaTypes
}
