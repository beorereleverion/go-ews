package elements

// The ItemCount element specifies the total number of items in a search result.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemcount
import "encoding/xml"

type ItemCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}
