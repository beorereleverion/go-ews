package elements

// The JoinDate element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/joindate
import "encoding/xml"

type JoinDate struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
