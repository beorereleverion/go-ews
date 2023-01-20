package elements

// The Folders element contains an array of folders that are used in folder operations.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/folders-ex15websvcsotherref
import "encoding/xml"

type Folders struct {
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

func (F *Folders) SetForMarshal() {
	F.XMLName.Local = "t:Folders"
}

func (F *Folders) GetSchema() *Schema {
	return &SchemaTypes
}
