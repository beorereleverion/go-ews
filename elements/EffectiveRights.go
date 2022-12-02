package elements

// The EffectiveRights element contains the client's rights based on the permission settings for the item or folder. This element is read-only.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/effectiverights
type EffectiveRights struct {
	// The CreateAssociated element indicates whether a client can create an associated contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CreateAssociated *CreateAssociated `xml:"t:CreateAssociated"`
	// The CreateContents element indicates whether a client can create a contents table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CreateContents *CreateContents `xml:"t:CreateContents"`
	// The CreateHierarchy element indicates whether a client can create a hierarchy table. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	CreateHierarchy *CreateHierarchy `xml:"t:CreateHierarchy"`
	// The Delete element indicates whether a client can delete a folder or item.
	Delete *Delete `xml:"t:Delete"`
	// The Modify element indicates whether a client can modify a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	Modify *Modify `xml:"t:Modify"`
	// The Read element indicates whether a client can read a folder or item. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	Read *Read `xml:"t:Read"`
	// The ViewPrivateItems element indicates whether a delegate user or client application has permission to view private items in the principal's mailbox.
	ViewPrivateItems *ViewPrivateItems `xml:"t:ViewPrivateItems"`
}
