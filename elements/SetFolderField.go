package elements

// The SetFolderField element represents an update that sets the value for a single property on a folder in an UpdateFolder operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setfolderfield
import "encoding/xml"

type SetFolderField struct {
	XMLName xml.Name

	// The CalendarFolder element represents a folder that primarily contains calendar items.
	CalendarFolder *CalendarFolder `xml:"CalendarFolder"`
	// The ContactsFolder element represents a contacts folder that is contained in a mailbox.
	ContactsFolder *ContactsFolder `xml:"ContactsFolder"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The Folder element defines a folder to create, get, find, synchronize, or update.
	Folder *Folder `xml:"Folder"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
	// The SearchFolder element represents a search folder that is contained in a mailbox.
	SearchFolder *SearchFolder `xml:"SearchFolder"`
	// The TasksFolder element represents a Tasks folder that is contained in a mailbox.
	TasksFolder *TasksFolder `xml:"TasksFolder"`
}

func (S *SetFolderField) SetForMarshal() {
	S.XMLName.Local = "t:SetFolderField"
}

func (S *SetFolderField) GetSchema() *Schema {
	return &SchemaTypes
}
