package elements

// The Excludes element performs a bitwise mask of the specified property and a supplied value.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/excludes
import "encoding/xml"

type Excludes struct {
	XMLName xml.Name

	// The Bitmask element represents a hexadecimal or decimal mask to be used during an Excludes restriction operation.
	Bitmask *Bitmask `xml:"Bitmask"`
	// The Excludes element performs a bitwise mask of the specified property and a supplied value.
	Excludes *Excludes `xml:"Excludes"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
}

func (E *Excludes) SetForMarshal() {
	E.XMLName.Local = "t:Excludes"
}

func (E *Excludes) GetSchema() *Schema {
	return &SchemaTypes
}
