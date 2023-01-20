package elements

// The Owners element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/owners
import "encoding/xml"

type Owners struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
