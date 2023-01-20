package elements

// The IsPin element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ispin
import "encoding/xml"

type IsPin struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
