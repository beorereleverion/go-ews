package elements

// The Argument element specifies arguments to the action.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/argument
import "encoding/xml"

type Argument struct {
	XMLName xml.Name

	// A non-empty string value that represents the value of an argument to the action part of a protection rule. This attribute is required.
	Value *string `xml:"Value,attr"`
}

func (A *Argument) SetForMarshal() {
	A.XMLName.Local = "t:Argument"
}

func (A *Argument) GetSchema() *Schema {
	return &SchemaTypes
}
