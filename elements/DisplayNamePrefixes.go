package elements

// The DisplayNamePrefixes element specifies an array of display name prefixes and the identifiers of their source attributions for the associated persona.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/displaynameprefixes
import "encoding/xml"

type DisplayNamePrefixes struct {
	XMLName xml.Name

	// The StringAttributedValue element specifies an instance in an array of attributes associated with a persona element.
	StringAttributedValue *StringAttributedValue `xml:"StringAttributedValue"`
}

func (D *DisplayNamePrefixes) SetForMarshal() {
	D.XMLName.Local = "t:DisplayNamePrefixes"
}

func (D *DisplayNamePrefixes) GetSchema() *Schema {
	return &SchemaTypes
}
