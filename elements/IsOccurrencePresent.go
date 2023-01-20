package elements

// The IsOccurrencePresent element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isoccurrencepresent
import "encoding/xml"

type IsOccurrencePresent struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
