package elements

// The Constant element identifies a constant value in a restriction.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/constant
import "encoding/xml"

type Constant struct {
	XMLName xml.Name

	// Specifies the value to compare in the restriction.
	Value *string `xml:"Value,attr"`
}

func (C *Constant) SetForMarshal() {
	C.XMLName.Local = "t:Constant"
}

func (C *Constant) GetSchema() *Schema {
	return &SchemaTypes
}
