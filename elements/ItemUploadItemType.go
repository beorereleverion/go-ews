package elements

// The Item element represents a single item to upload into a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/item-uploaditemtype
type ItemUploadItemType struct {
	// The Data element contains the data of a single exported item or an item to upload into a mailbox.
	Data *Database64Binary `xml:"m:Data"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The ParentFolderId element represents the identifier of the parent folder that contains the item or folder.
	ParentFolderId *ParentFolderId `xml:"t:ParentFolderId"`
	// Specifies the action for uploading an item into a mailbox. This attribute is required.
	CreateAction *string `xml:"CreateAction,attr"`
	// Specifies whether the uploaded item is a folder associated item. This attribute is a Boolean value. A value of true indicates that the item is a folder associated item. This attribute is optional.
	IsAssociated *string `xml:"IsAssociated,attr"`
}
