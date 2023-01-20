package elements

// The NextPredictedAction element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nextpredictedaction
import "encoding/xml"

type NextPredictedAction struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
