package elements

// The GroupingAction element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupingaction
import "encoding/xml"

type GroupingAction struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
