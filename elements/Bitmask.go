package elements

// The Bitmask element represents a hexadecimal or decimal mask to be used during an Excludes restriction operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/bitmask
import "encoding/xml"

type Bitmask struct {
	XMLName xml.Name

	// Represents a decimal or hexadecimal bitmask. The value is represented by the following regular expression:((0x|0X)[0-9A-Fa-f]*)|([0-9]*).The following are examples of hexadecimal values for this attribute:- 0x12AF- 0X334AEThe following are examples of decimal values for this attribute:- 10- 255- 4562
	Value *string `xml:"Value,attr"`
}

func (B *Bitmask) SetForMarshal() {
	B.XMLName.Local = "t:Bitmask"
}

func (B *Bitmask) GetSchema() *Schema {
	return &SchemaTypes
}
