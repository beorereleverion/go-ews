package elements

// The CreateFolderResponse element defines a response to a CreateFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createfolderresponse
import "encoding/xml"

type CreateFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CreateFolderResponse) SetForMarshal() {
	C.XMLName.Local = "m:CreateFolderResponse"
}

func (C *CreateFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
