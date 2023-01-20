package elements

// The CopyFolderResponse element defines a response to a CopyFolder request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/copyfolderresponse
import "encoding/xml"

type CopyFolderResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (C *CopyFolderResponse) SetForMarshal() {
	C.XMLName.Local = "m:CopyFolderResponse"
}

func (C *CopyFolderResponse) GetSchema() *Schema {
	return &SchemaMessages
}
