package elements

// The SendItemResponse element defines a response to a SendItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/senditemresponse
import "encoding/xml"

type SendItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (S *SendItemResponse) SetForMarshal() {
	S.XMLName.Local = "m:SendItemResponse"
}

func (S *SendItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
