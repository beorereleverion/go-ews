package elements

// The SyncFolderItemsResponse element defines a response to a SyncFolderItems request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/syncfolderitemsresponse
import "encoding/xml"

type SyncFolderItemsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (S *SyncFolderItemsResponse) SetForMarshal() {
	S.XMLName.Local = "m:SyncFolderItemsResponse"
}

func (S *SyncFolderItemsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
