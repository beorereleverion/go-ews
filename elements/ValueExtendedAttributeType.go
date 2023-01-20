package elements

// The Value (ExendedAttributeType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/value-extendedattributetype
import "encoding/xml"

type ValueExtendedAttributeType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
