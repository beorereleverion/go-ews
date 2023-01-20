package elements

// The ExtendedAttribute element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedattribute
import "encoding/xml"

type ExtendedAttribute struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
