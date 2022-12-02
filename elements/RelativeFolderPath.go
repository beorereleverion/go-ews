package elements

// The RelativeFolderPath element contains an array of folders that indicate the relative folder path of the folder path to be created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/relativefolderpath
type RelativeFolderPath struct {
	// The CalendarFolder element represents a folder that primarily contains calendar items.
	CalendarFolder *CalendarFolder `xml:"t:CalendarFolder"`
	// The ContactsFolder element represents a contacts folder that is contained in a mailbox.
	ContactsFolder *ContactsFolder `xml:"t:ContactsFolder"`
	// The Folder element defines a folder to create, get, find, synchronize, or update.
	Folder *Folder `xml:"t:Folder"`
	// The SearchFolder element represents a search folder that is contained in a mailbox.
	SearchFolder *SearchFolder `xml:"t:SearchFolder"`
	// The TasksFolder element represents a Tasks folder that is contained in a mailbox.
	TasksFolder *TasksFolder `xml:"t:TasksFolder"`
}
