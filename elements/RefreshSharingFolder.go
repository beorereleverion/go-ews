package elements

// The RefreshSharingFolder element defines a request to refresh the specified local folder. It is the base element for the RefreshSharingFolder operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/refreshsharingfolder
type RefreshSharingFolder struct {
	// The SharingFolderId element represents the identifier of the local folder in a sharing relationship.
	SharingFolderId *SharingFolderId `xml:"m:SharingFolderId"`
}
