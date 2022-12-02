package elements

// The SetFolderField element represents an update that sets the value for a single property on a folder in an UpdateFolder operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setfolderfield
type SetFolderField struct {
	// The CalendarFolder element represents a folder that primarily contains calendar items.
	CalendarFolder *CalendarFolder `xml:"t:CalendarFolder"`
	// The ContactsFolder element represents a contacts folder that is contained in a mailbox.
	ContactsFolder *ContactsFolder `xml:"t:ContactsFolder"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"t:ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"t:FieldURI"`
	// The Folder element defines a folder to create, get, find, synchronize, or update.
	Folder *Folder `xml:"t:Folder"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"t:IndexedFieldURI"`
	// The SearchFolder element represents a search folder that is contained in a mailbox.
	SearchFolder *SearchFolder `xml:"t:SearchFolder"`
	// The TasksFolder element represents a Tasks folder that is contained in a mailbox.
	TasksFolder *TasksFolder `xml:"t:TasksFolder"`
}
