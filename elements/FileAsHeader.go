package elements

// The FileAsHeader specifies the header for the File As option.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileasheader
import "encoding/xml"

type FileAsHeader struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (F *FileAsHeader) SetForMarshal() {
	F.XMLName.Local = "t:FileAsHeader"
}

func (F *FileAsHeader) GetSchema() *Schema {
	return &SchemaTypes
}
