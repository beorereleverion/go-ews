package elements

// The LargeAudienceCap element specifies the maximum number of recipients for an email message.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/largeaudiencecap
import "encoding/xml"

type LargeAudienceCap struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}
