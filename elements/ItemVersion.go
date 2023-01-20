package elements

// The ItemVersion element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemversion
import "encoding/xml"

type ItemVersion struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
