package elements

// The FileAsIds element specifies an array of StringAttributedValue elements and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileasids
import "encoding/xml"

type FileAsIds struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (F *FileAsIds) SetForMarshal() {
	F.XMLName.Local = "t:FileAsIds"
}

func (F *FileAsIds) GetSchema() *Schema {
	return &SchemaTypes
}
