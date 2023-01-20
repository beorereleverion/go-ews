package elements

// The GroupingAction (ItemType) element is intended for internal use only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupingaction-itemtype
import "encoding/xml"

type GroupingActionItemType struct {
	XMLName xml.Name
	TEXT    interface{} `xml:",chardata"`
}
