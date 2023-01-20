package elements

// The ResolveNames element defines a request to resolve ambiguous names.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/resolvenames
import "encoding/xml"

type ResolveNames struct {
	XMLName xml.Name

	// The ParentFolderIds element identifies folders for the FindItem and FindFolder operations to search.
	ParentFolderIds *ParentFolderIds `xml:"ParentFolderIds"`
	// The UnresolvedEntry element contains the name of a contact or distribution list to resolve.
	UnresolvedEntry *UnresolvedEntry `xml:"UnresolvedEntry"`
	// Identifies the property set returned for contacts. This attribute was introduced in Exchange Server 2010 Service Pack 2 (SP2).
	ContactDataShape *string `xml:"ContactDataShape,attr"`
	// Describes whether the full contact details for public contacts for a resolved name are returned in the response. This attribute is required for public contacts. This value does not affect private contacts and private distribution lists, for which ItemId is always returned.
	ReturnFullContactData *string `xml:"ReturnFullContactData,attr"`
	// Identifies the order and scope for a ResolveNames search.
	SearchScope *string `xml:"SearchScope,attr"`
}

const (
	// The contact item identifier property is returned.
	ResolveNamesIdOnly = `IdOnly`
	// The Default set of contact item properties is returned. For more information, see Response shapes in EWS.
	ResolveNamesDefault = `Default`
	// The AllProperties set of contact item properties are returned. For more information, see Response shapes in EWS.
	ResolveNamesAllProperties = `AllProperties`
	// Full contact details for public contacts are returned.
	ResolveNamesTrue = `True`
	// Full contact details for public contacts are not returned.
	ResolveNamesFalse = `False`
	// Only the Active Directory directory service is searched.
	ResolveNamesActiveDirectory = `ActiveDirectory`
	// Active Directory is searched first, and then the contact folders that are specified in the ParentFolderIds property are searched.
	ResolveNamesActiveDirectoryContacts = `ActiveDirectoryContacts`
	// Only the contact folders that are identified by the ParentFolderIds property are searched.
	ResolveNamesContacts = `Contacts`
	// Contact folders that are identified by the ParentFolderIds property are searched first and then Active Directory is searched.
	ResolveNamesContactsActiveDirectory = `ContactsActiveDirectory`
)

func (R *ResolveNames) SetForMarshal() {
	R.XMLName.Local = "m:ResolveNames"
}

func (R *ResolveNames) GetSchema() *Schema {
	return &SchemaMessages
}
