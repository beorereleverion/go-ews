package elements

// The IsMember element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/ismember
import "encoding/xml"

type IsMember struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
