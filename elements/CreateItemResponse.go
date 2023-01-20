package elements

// The CreateItemResponse element defines a response to a CreateItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createitemresponse
import "encoding/xml"

type CreateItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CreateItemResponse) SetForMarshal() {
	C.XMLName.Local = "m:CreateItemResponse"
}

func (C *CreateItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
