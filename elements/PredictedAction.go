package elements

// The PredictedAction element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/predictedaction
import "encoding/xml"

type PredictedAction struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
