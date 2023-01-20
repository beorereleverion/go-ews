package elements

// The AddNewImContactToGroup element defines a request to add a new instant messaging contact to an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addnewimcontacttogroup
import "encoding/xml"

type AddNewImContactToGroup struct {
	XMLName xml.Name

	// The DisplayName element contains the display name of a new instant messaging group contact or the display name of a new instant messaging group.
	DisplayName *DisplayNameNonEmptyStringType `xml:"DisplayName"`
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"GroupId"`
	// The ImAddress element contains the instant messaging address of a new contact that will be added to an instant messaging group.
	ImAddress *ImAddressNonEmptyStringType `xml:"ImAddress"`
}

func (A *AddNewImContactToGroup) SetForMarshal() {
	A.XMLName.Local = "m:AddNewImContactToGroup"
}

func (A *AddNewImContactToGroup) GetSchema() *Schema {
	return &SchemaMessages
}
