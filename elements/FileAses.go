package elements

// The FileAses element specifies an array of StringAttributedValue elements and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fileases
import "encoding/xml"

type FileAses struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (F *FileAses) SetForMarshal() {
	F.XMLName.Local = "t:FileAses"
}

func (F *FileAses) GetSchema() *Schema {
	return &SchemaTypes
}
