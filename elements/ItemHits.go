package elements

// The ItemHits element identifies how many times a keyword was found.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/itemhits
import "encoding/xml"

type ItemHits struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}
