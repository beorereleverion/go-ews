package elements

// The Create element identifies a single folder to create in the local client store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/create-foldersync
import "encoding/xml"

type CreateFolderSync struct {
	XMLName xml.Name

	// The CalendarFolder element represents a folder that primarily contains calendar items.
	CalendarFolder *CalendarFolder `xml:"CalendarFolder"`
	// The ContactsFolder element represents a contacts folder that is contained in a mailbox.
	ContactsFolder *ContactsFolder `xml:"ContactsFolder"`
	// The Folder element defines a folder to create, get, find, synchronize, or update.
	Folder *Folder `xml:"Folder"`
	// The SearchFolder element represents a search folder that is contained in a mailbox.
	SearchFolder *SearchFolder `xml:"SearchFolder"`
	// The TasksFolder element represents a Tasks folder that is contained in a mailbox.
	TasksFolder *TasksFolder `xml:"TasksFolder"`
}

func (C *CreateFolderSync) SetForMarshal() {
	C.XMLName.Local = "t:Create"
}

func (C *CreateFolderSync) GetSchema() *Schema {
	return &SchemaTypes
}
