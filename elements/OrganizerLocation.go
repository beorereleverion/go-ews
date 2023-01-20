package elements

// The OrganizerLocation element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/organizerlocation
import "encoding/xml"

type OrganizerLocation struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
