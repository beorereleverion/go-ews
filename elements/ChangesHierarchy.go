package elements

// The Changes element contains a sequenced array of change types that represent the type of differences between the folders on the client and the folders on the computer that is running Microsoft Exchange Server 2007.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/changes-hierarchy
type ChangesHierarchy struct {
	// The Create element identifies a single folder to create in the local client store.
	Create *CreateFolderSync `xml:"t:Create"`
	// The Delete element identifies a single folder to delete in the local client store.
	Delete *DeleteFolderSync `xml:"t:Delete"`
	// The Update element identifies a single folder to update in the local client store.
	Update *UpdateFolderSync `xml:"t:Update"`
}
