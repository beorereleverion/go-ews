package elements

// The AddNewImContactToGroup element defines a request to add a new instant messaging contact to an instant messaging group.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/addnewimcontacttogroup
type AddNewImContactToGroup struct {
	// The DisplayName element contains the display name of a new instant messaging group contact or the display name of a new instant messaging group.
	DisplayName *DisplayNameNonEmptyStringType `xml:"t:DisplayName"`
	// The GroupId element uniquely identifies a group.
	GroupId *GroupId `xml:"m:GroupId"`
	// The ImAddress element contains the instant messaging address of a new contact that will be added to an instant messaging group.
	ImAddress *ImAddressNonEmptyStringType `xml:"m:ImAddress"`
}
