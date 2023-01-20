package elements

// The GetItemResponse element defines a response to a GetItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getitemresponse
import "encoding/xml"

type GetItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetItemResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetItemResponse"
}

func (G *GetItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
