package elements

// The PredictedActionReason element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/predictedactionreason
import "encoding/xml"

type PredictedActionReason struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
