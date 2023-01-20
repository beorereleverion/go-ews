package elements

// The GetFolderResponse element defines a response to a GetFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getfolderresponse
import "encoding/xml"

type GetFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetFolderResponse) SetForMarshal() {
	G.XMLName.Local = "m:GetFolderResponse"
}

func (G *GetFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
