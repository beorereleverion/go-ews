package elements

// The ImItemList element contains a list of instant messaging groups and instant messaging contacts.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/imitemlist
type ImItemList struct {
	// The Groups element represents an array of instant messaging (IM) groups.
	Groups *GroupsArrayOfImGroupType `xml:"Groups"`
	// The Personas element specifies an array of personas returned from the GetImItems and GetImItemList operations.
	Personas *Personas `xml:"t:Personas"`
}
