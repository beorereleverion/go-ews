package elements

// The DirectoryId element contains the directory ID of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/directoryid
import "encoding/xml"

type DirectoryId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (D *DirectoryId) SetForMarshal() {
	D.XMLName.Local = "t:DirectoryId"
}

func (D *DirectoryId) GetSchema() *Schema {
	return &SchemaTypes
}
