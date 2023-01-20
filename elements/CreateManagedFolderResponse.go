package elements

// The CreateManagedFolderResponse element defines a response to a CreateManagedFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createmanagedfolderresponse
import "encoding/xml"

type CreateManagedFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CreateManagedFolderResponse) SetForMarshal() {
	C.XMLName.Local = "m:CreateManagedFolderResponse"
}

func (C *CreateManagedFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
