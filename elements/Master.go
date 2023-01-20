package elements

// The Master element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/master
import "encoding/xml"

type Master struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
