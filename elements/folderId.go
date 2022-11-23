package elements

// The FolderId element contains the identifier and change key of a folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folderid
type FolderId struct {
	// Contains a string that identifies a folder in the Exchange store. This attribute is required.
	ID string `xml:"Id,attr"`
	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey string `xml:"ChangeKey,attr"`
}
