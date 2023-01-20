package elements

// The ExpandDLResponse element defines a response to a request to expand a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/expanddlresponse
import "encoding/xml"

type ExpandDLResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (E *ExpandDLResponse) SetForMarshal() {
	E.XMLName.Local = "t:ExpandDLResponse"
}

func (E *ExpandDLResponse) GetSchema() *Schema {
	return &SchemaTypes
}
