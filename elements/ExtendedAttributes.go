package elements

// The ExtendedAttributes element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/extendedattributes
import "encoding/xml"

type ExtendedAttributes struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
