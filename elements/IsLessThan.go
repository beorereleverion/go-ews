package elements

// The IsLessThan element represents a search expression that compares a property with either a constant value or another property and returns true if the first property is less than the second.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/islessthan
import "encoding/xml"

type IsLessThan struct {
	XMLName xml.Name

	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
	FieldURIOrConstant *FieldURIOrConstant `xml:"FieldURIOrConstant"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
}

func (I *IsLessThan) SetForMarshal() {
	I.XMLName.Local = "t:IsLessThan"
}

func (I *IsLessThan) GetSchema() *Schema {
	return &SchemaTypes
}
