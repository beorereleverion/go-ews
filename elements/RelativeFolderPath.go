package elements

// The RelativeFolderPath element contains an array of folders that indicate the relative folder path of the folder path to be created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/relativefolderpath
import "encoding/xml"

type RelativeFolderPath struct {
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

func (R *RelativeFolderPath) SetForMarshal() {
	R.XMLName.Local = "m:RelativeFolderPath"
}

func (R *RelativeFolderPath) GetSchema() *Schema {
	return &SchemaMessages
}
