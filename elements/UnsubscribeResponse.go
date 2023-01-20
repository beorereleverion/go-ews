package elements

// The UnsubscribeResponse element defines a response to an Unsubscribe request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/unsubscriberesponse
import "encoding/xml"

type UnsubscribeResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (U *UnsubscribeResponse) SetForMarshal() {
	U.XMLName.Local = "m:UnsubscribeResponse"
}

func (U *UnsubscribeResponse) GetSchema() *Schema {
	return &SchemaMessages
}
