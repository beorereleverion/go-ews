package elements

// The NextPredictedAction (ItemType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/nextpredictedaction-itemtype
import "encoding/xml"

type NextPredictedActionItemType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
