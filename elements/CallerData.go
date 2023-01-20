package elements

// The CallerData element in intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/callerdata
import "encoding/xml"

type CallerData struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
