package elements

// The Identity element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/identity
import "encoding/xml"

type Identity struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
