package elements

// The Url element represents the location of the client Web service for push notifications.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/url-ex15websvcsotherref
import "encoding/xml"

type Url struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
