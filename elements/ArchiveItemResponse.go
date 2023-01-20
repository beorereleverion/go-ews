package elements

// The ArchiveItemResponse element specifies the response to an ArchiveItem request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/archiveitemresponse
import "encoding/xml"

type ArchiveItemResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (A *ArchiveItemResponse) SetForMarshal() {
	A.XMLName.Local = "m:ArchiveItemResponse"
}

func (A *ArchiveItemResponse) GetSchema() *Schema {
	return &SchemaMessages
}
