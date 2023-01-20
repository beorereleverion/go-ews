package elements

// The Keyword element specifies a single keyword.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/keyword
import "encoding/xml"

type Keyword struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
