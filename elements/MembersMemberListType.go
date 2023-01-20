package elements

// The Members element provides the list of members for a distribution list.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/members-memberlisttype
import "encoding/xml"

type MembersMemberListType struct {
	XMLName xml.Name

	// The Member element represents a member of a distribution list.
	Member *Member `xml:"Member"`
}

func (M *MembersMemberListType) SetForMarshal() {
	M.XMLName.Local = "t:Members"
}

func (M *MembersMemberListType) GetSchema() *Schema {
	return &SchemaTypes
}
