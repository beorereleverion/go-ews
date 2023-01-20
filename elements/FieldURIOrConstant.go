package elements

// The FieldURIOrConstant element represents either a property or a constant value to be used when comparing with another property.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fielduriorconstant
import "encoding/xml"

type FieldURIOrConstant struct {
	XMLName xml.Name

	// The Constant element identifies a constant value in a restriction.
	Constant *Constant `xml:"Constant"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
}

func (F *FieldURIOrConstant) SetForMarshal() {
	F.XMLName.Local = "t:FieldURIOrConstant"
}

func (F *FieldURIOrConstant) GetSchema() *Schema {
	return &SchemaTypes
}
