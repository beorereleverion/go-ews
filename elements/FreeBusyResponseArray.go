package elements

// The FreeBusyResponseArray element contains the requested users' availability information and the response status.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/freebusyresponsearray
import "encoding/xml"

type FreeBusyResponseArray struct {
	XMLName xml.Name

	// The FreeBusyResponse element contains the free/busy information for a single mailbox user.
	FreeBusyResponse *FreeBusyResponse `xml:"FreeBusyResponse"`
}

func (F *FreeBusyResponseArray) SetForMarshal() {
	F.XMLName.Local = "m:FreeBusyResponseArray"
}

func (F *FreeBusyResponseArray) GetSchema() *Schema {
	return &SchemaMessages
}
