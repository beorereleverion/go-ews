package elements

// The Notes element contains supplementary contact information.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/notes-contact
import "encoding/xml"

type NotesContact struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NotesContact) SetForMarshal() {
	N.XMLName.Local = "t:Notes"
}

func (N *NotesContact) GetSchema() *Schema {
	return &SchemaTypes
}
