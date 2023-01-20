package elements

// The SetItemField element represents an update to a single property of an item in an UpdateItem operation.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/setitemfield
import "encoding/xml"

type SetItemField struct {
	XMLName xml.Name

	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"CalendarItem"`
	// The Contact element represents a contact item in the Exchange store.
	Contact *Contact `xml:"Contact"`
	// The DistributionList element represents a distribution list.
	DistributionList *DistributionList `xml:"DistributionList"`
	// The ExtendedFieldURI element identifies an extended MAPI property.
	ExtendedFieldURI *ExtendedFieldURI `xml:"ExtendedFieldURI"`
	// The FieldURI element identifies frequently referenced properties by URI.
	FieldURI *FieldURI `xml:"FieldURI"`
	// The IndexedFieldURI element identifies individual members of a dictionary.
	IndexedFieldURI *IndexedFieldURI `xml:"IndexedFieldURI"`
	// The Item element represents a generic item in the Exchange store.
	Item *Item `xml:"Item"`
	// The MeetingCancellation element represents a meeting cancellation in the Exchange store.
	MeetingCancellation *MeetingCancellation `xml:"MeetingCancellation"`
	// The MeetingMessage element represents a meeting in the Exchange store.
	MeetingMessage *MeetingMessage `xml:"MeetingMessage"`
	// The MeetingRequest element represents a meeting request in the Exchange store.
	MeetingRequest *MeetingRequest `xml:"MeetingRequest"`
	// The MeetingResponse element represents a meeting response in the Exchange store.
	MeetingResponse *MeetingResponse `xml:"MeetingResponse"`
	// The Message element represents a Microsoft Exchange e-mail message.
	Message *Message `xml:"Message"`
	// The Task element represents a task in the Exchange store.
	Task *Task `xml:"Task"`
}

func (S *SetItemField) SetForMarshal() {
	S.XMLName.Local = "t:SetItemField"
}

func (S *SetItemField) GetSchema() *Schema {
	return &SchemaTypes
}
