package elements

// The DeleteItemResponse element defines a response to a single DeleteItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/deleteitemresponse
import "encoding/xml"

type DeleteItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (D *DeleteItemResponse) SetForMarshal() {
	D.XMLName.Local = "m:DeleteItemResponse"
}

func (D *DeleteItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
