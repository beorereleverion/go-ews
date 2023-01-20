package elements

// The MoveItemResponse element defines a response to a MoveItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/moveitemresponse
import "encoding/xml"

type MoveItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (M *MoveItemResponse) SetForMarshal() {
	M.XMLName.Local = "m:MoveItemResponse"
}

func (M *MoveItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
