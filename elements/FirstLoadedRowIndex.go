package elements

// The FirstLoadedRowIndex element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/firstloadedrowindex
import "encoding/xml"

type FirstLoadedRowIndex struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
