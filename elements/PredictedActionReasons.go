package elements

// The PredictedActionReasons element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/predictedactionreasons
import "encoding/xml"

type PredictedActionReasons struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
