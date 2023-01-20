package elements

// The AlternateId element describes an identifier to convert in a request and the results of a converted identifier in the response.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/alternateid
import "encoding/xml"

type AlternateId struct {
	XMLName xml.Name

	// Describes the source format in a ConvertId operation request and describes the destination format in a ConvertId operation response. The destination format is described by the DestinationFormat attribute of the ConvertId element in the request. This attribute is of type IdFormatType.
	Format *string `xml:"Format,attr"`
	// Describes the source identifier in a ConvertId operation request and describes the destination identifier in a ConvertId operation response.
	Id *string `xml:"Id,attr"`
	// Indicates whether the identifier represents an archived item or folder. A value of true indicates that the identifier represents an archived item or folder. This attribute is optional.
	IsArchive *string `xml:"IsArchive,attr"`
	// Describes the mailbox primary Simple Mail Transfer Protocol (SMTP) address that contains the identifiers to translate.
	Mailbox *string `xml:"Mailbox,attr"`
}

const (
	// Describes identifiers that are produced by Exchange Web Services in the initial release version of Exchange 2007.
	AlternateIdEwsLegacyId = `EwsLegacyId`
	// Describes identifiers that are produced by Exchange Web Services starting with Exchange 2007 SP1.
	AlternateIdEwsId = `EwsId`
	// Describes MAPI identifiers, as in the PR_ENTRYID property.
	AlternateIdEntryId = `EntryId`
	// Describes a hexadecimal-encoded representation of the PR_ENTRYID property. This is the format of availability calendar event identifiers.
	AlternateIdHexEntryId = `HexEntryId`
	// Describes Exchange store identifiers.
	AlternateIdStoreId = `StoreId`
	// Describes an Outlook Web App identifier.
	AlternateIdOwaId = `OwaId`
)

func (A *AlternateId) SetForMarshal() {
	A.XMLName.Local = "m:AlternateId"
}

func (A *AlternateId) GetSchema() *Schema {
	return &SchemaMessages
}
