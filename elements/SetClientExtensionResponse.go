package elements

// The SetClientExtensionResponse element contains the response to a SetClientExtension request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setclientextensionresponse
import "encoding/xml"

type SetClientExtensionResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (S *SetClientExtensionResponse) SetForMarshal() {
	S.XMLName.Local = "m:SetClientExtensionResponse"
}

func (S *SetClientExtensionResponse) GetSchema() *Schema {
	return &SchemaMessages
}
