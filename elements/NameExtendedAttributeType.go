package elements

// The Name (ExtendedAttributeType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/name-extendedattributetype
import "encoding/xml"

type NameExtendedAttributeType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
