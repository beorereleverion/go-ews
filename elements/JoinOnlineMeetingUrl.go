package elements

// The JoinOnlineMeetingUrl element specifies the URL to join an online meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/joinonlinemeetingurl
import "encoding/xml"

type JoinOnlineMeetingUrl struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}
