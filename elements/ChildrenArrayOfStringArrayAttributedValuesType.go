package elements

// The Children element specifies an array of child names and identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/children-arrayofstringarrayattributedvaluestype
import "encoding/xml"

type ChildrenArrayOfStringArrayAttributedValuesType struct {
	XMLName xml.Name

	// The StringArrayAttributedValue element specifies an instance of an array of string data for a persona element.
	StringArrayAttributedValue *StringArrayAttributedValue `xml:"StringArrayAttributedValue"`
}

func (C *ChildrenArrayOfStringArrayAttributedValuesType) SetForMarshal() {
	C.XMLName.Local = "t:Children"
}

func (C *ChildrenArrayOfStringArrayAttributedValuesType) GetSchema() *Schema {
	return &SchemaTypes
}
