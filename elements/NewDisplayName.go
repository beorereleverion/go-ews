package elements

// The NewDisplayName element contains the updated display name of an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/newdisplayname
import "encoding/xml"

type NewDisplayName struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (N *NewDisplayName) SetForMarshal() {
	N.XMLName.Local = "m:NewDisplayName"
}

func (N *NewDisplayName) GetSchema() *Schema {
	return &SchemaMessages
}
