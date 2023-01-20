package elements

// The SetImGroup element represents a request to change the display name of an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setimgroup
import "encoding/xml"

type SetImGroup struct {
	XMLName xml.Name

	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"GroupId"`
	// The NewDisplayName element contains the updated display name of an instant messaging group.
	NewDisplayName *NewDisplayName `xml:"NewDisplayName"`
}

func (S *SetImGroup) SetForMarshal() {
	S.XMLName.Local = "m:SetImGroup"
}

func (S *SetImGroup) GetSchema() *Schema {
	return &SchemaMessages
}
