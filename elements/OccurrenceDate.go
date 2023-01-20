package elements

// The OccurrenceDate element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/occurrencedate
import "encoding/xml"

type OccurrenceDate struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
