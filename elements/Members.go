package elements

// The Members element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/members
import "encoding/xml"

type Members struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
