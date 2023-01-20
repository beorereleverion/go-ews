package elements

// The FileAsMapping element defines how to construct what is displayed for a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileasmapping
import "encoding/xml"

type FileAsMapping struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FileAsMapping) SetForMarshal() {
	F.XMLName.Local = "t:FileAsMapping"
}

func (F *FileAsMapping) GetSchema() *Schema {
	return &SchemaTypes
}
