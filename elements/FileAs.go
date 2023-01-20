package elements

// The FileAs element represents how a contact or distribution list is filed in the Contacts folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileas
import "encoding/xml"

type FileAs struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FileAs) SetForMarshal() {
	F.XMLName.Local = "t:FileAs"
}

func (F *FileAs) GetSchema() *Schema {
	return &SchemaTypes
}
