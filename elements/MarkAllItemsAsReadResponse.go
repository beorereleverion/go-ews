package elements

// The MarkAllItemsAsReadResponse element specifies the response to a MarkAllItemsAsRead request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/markallitemsasreadresponse
import "encoding/xml"

type MarkAllItemsAsReadResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (M *MarkAllItemsAsReadResponse) SetForMarshal() {
	M.XMLName.Local = "m:MarkAllItemsAsReadResponse"
}

func (M *MarkAllItemsAsReadResponse) GetSchema() *Schema {
	return &SchemaMessages
}
