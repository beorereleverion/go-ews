package elements

// The UserMailbox element identifies a user mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/usermailbox
import "encoding/xml"

type UserMailbox struct {
	XMLName xml.Name

	// The text value of the Id attribute is the identifier of the mailbox.
	Id *string `xml:"Id,attr"`
	// The text value of the IsArchive attribute indicates whether the mailbox is an archive mailbox. A text value of true for the IsArchive attribute indicates that the mailbox is an archive mailbox. A value of false for the IsArchive attribute indicates that the mailbox is a primary mailbox.
	IsArchive *string `xml:"IsArchive,attr"`
}

func (U *UserMailbox) SetForMarshal() {
	U.XMLName.Local = "t:UserMailbox"
}

func (U *UserMailbox) GetSchema() *Schema {
	return &SchemaTypes
}
