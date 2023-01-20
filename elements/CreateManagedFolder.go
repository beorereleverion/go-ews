package elements

// The CreateManagedFolder element defines a request to add managed custom folders to a mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createmanagedfolder
import "encoding/xml"

type CreateManagedFolder struct {
	XMLName xml.Name

	// The FolderNames element contains an array of named managed folders to add to a mailbox.
	FolderNames *FolderNames `xml:"FolderNames"`
	// The Mailbox element identifies a mail-enabled Active Directory object.
	Mailbox *Mailbox `xml:"Mailbox"`
}

func (C *CreateManagedFolder) SetForMarshal() {
	C.XMLName.Local = "m:CreateManagedFolder"
}

func (C *CreateManagedFolder) GetSchema() *Schema {
	return &SchemaMessages
}
