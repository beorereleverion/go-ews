package elements

// The IsClutter element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isclutter
import "encoding/xml"

type IsClutter struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
