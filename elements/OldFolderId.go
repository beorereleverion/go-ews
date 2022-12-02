package elements

// The OldFolderId element contains the original identifier of a folder that was moved or copied.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/oldfolderid
type OldFolderId struct {
	// Contains a string that identifies a version of a folder that is identified by the Id attribute. This attribute is optional. Use this attribute to make sure that the correct version of a folder is used.
	ChangeKey *string `xml:"ChangeKey,attr"`
	// Contains a string that identifies a folder in the Exchange store. This attribute is required.
	Id *string `xml:"Id,attr"`
}
