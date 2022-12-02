package elements

// The SourceIds element contains the source identifiers to convert.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/sourceids
type SourceIds struct {
	// The AlternateId element describes an identifier to convert in a request and the results of a converted identifier in the response.
	AlternateId *AlternateId `xml:"m:AlternateId"`
	// The AlternatePublicFolderId element describes a public folder identifier to convert to another identifier format. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	AlternatePublicFolderId *AlternatePublicFolderId `xml:"t:AlternatePublicFolderId"`
	// The AlternatePublicFolderItemId element describes a public folder item identifier to convert to another identifier format. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
	AlternatePublicFolderItemId *AlternatePublicFolderItemId `xml:"t:AlternatePublicFolderItemId"`
}
