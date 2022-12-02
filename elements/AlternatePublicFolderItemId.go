package elements

// The AlternatePublicFolderItemId element describes a public folder item identifier to convert to another identifier format. This element was introduced in Microsoft Exchange Server 2007 Service Pack 1 (SP1).
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/alternatepublicfolderitemid
type AlternatePublicFolderItemId struct {
	// Identifies the public folder that contains the public folder item. This attribute is required.
	FolderId *string `xml:"FolderId,attr"`
	// Identifies the format that describes the public folder item identifier to convert. This attribute is required.
	Format *string `xml:"Format,attr"`
	// Identifier the public folder item to convert. This attribute is required.
	ItemId *string `xml:"ItemId,attr"`
}

const (
	// Describes identifiers that are produced by Exchange Web Services in the initial release version of Exchange 2007.
	AlternatePublicFolderItemIdEwsLegacyId = `EwsLegacyId`
	// Describes identifiers that are produced by Exchange Web Services starting with Exchange 2007 SP1.
	AlternatePublicFolderItemIdEwsId = `EwsId`
	// Describes MAPI identifiers, as in the PR_ENTRYID property.
	AlternatePublicFolderItemIdEntryId = `EntryId`
	// Describes a hexadecimal-encoded representation of the PR_ENTRYID property. This is the format of availability calendar event identifiers.
	AlternatePublicFolderItemIdHexEntryId = `HexEntryId`
	// Describes Exchange store identifiers.
	AlternatePublicFolderItemIdStoreId = `StoreId`
	// Describes an Outlook Web Access identifier.
	AlternatePublicFolderItemIdOwaId = `OwaId`
)
