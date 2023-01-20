package elements

// The Association element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/association
import "encoding/xml"

type Association struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
