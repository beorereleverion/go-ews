package elements

// The FreeBusyResponse element contains the free/busy information for a single mailbox user.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyresponse
import "encoding/xml"

type FreeBusyResponse struct {
	XMLName xml.Name

	// The FreeBusyView element contains availability information for a specific user.
	FreeBusyView *FreeBusyView `xml:"FreeBusyView"`
	// The ResponseMessage element provides descriptive information about the response status for a single entity within a request.
	ResponseMessage *ResponseMessage `xml:"ResponseMessage"`
}

func (F *FreeBusyResponse) SetForMarshal() {
	F.XMLName.Local = "m:FreeBusyResponse"
}

func (F *FreeBusyResponse) GetSchema() *Schema {
	return &SchemaMessages
}
