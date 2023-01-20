package elements

// The FileAsId element specifies the FileAs identifier.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileasid
import "encoding/xml"

type FileAsId struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FileAsId) SetForMarshal() {
	F.XMLName.Local = "t:FileAsId"
}

func (F *FileAsId) GetSchema() *Schema {
	return &SchemaTypes
}
