package elements

// The UpdateItemResponse element defines a response to an UpdateItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/updateitemresponse
import "encoding/xml"

type UpdateItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (U *UpdateItemResponse) SetForMarshal() {
	U.XMLName.Local = "m:UpdateItemResponse"
}

func (U *UpdateItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
