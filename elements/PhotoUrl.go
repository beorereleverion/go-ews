package elements

// The PhotoUrl element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/photourl
import "encoding/xml"

type PhotoUrl struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
