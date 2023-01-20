package elements

// The AcceptSharingInvitation element is used to accept an invitation that allows access to another user's calendar or contacts data.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/acceptsharinginvitation
import "encoding/xml"

type AcceptSharingInvitation struct {
	XMLName xml.Name

	// The ReferenceItemId element identifies the item to which the response object refers.
	ReferenceItemId *ReferenceItemId `xml:"ReferenceItemId"`
}

func (A *AcceptSharingInvitation) SetForMarshal() {
	A.XMLName.Local = "t:AcceptSharingInvitation"
}

func (A *AcceptSharingInvitation) GetSchema() *Schema {
	return &SchemaTypes
}
