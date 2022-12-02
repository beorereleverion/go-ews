package elements

// The Update element identifies a single item to update in the local client store.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/update-itemsync
type UpdateItemSync struct {
	// The CalendarItem element represents an Exchange calendar item.
	CalendarItem *CalendarItem `xml:"t:CalendarItem"`
	// The Contact element represents a contact item in the Exchange store.
	Contact *Contact `xml:"t:Contact"`
	// The DistributionList element represents a distribution list.
	DistributionList *DistributionList `xml:"t:DistributionList"`
	// The Item element represents a generic item in the Exchange store.
	Item *Item `xml:"t:Item"`
	// The MeetingCancellation element represents a meeting cancellation in the Exchange store.
	MeetingCancellation *MeetingCancellation `xml:"t:MeetingCancellation"`
	// The MeetingMessage element represents a meeting in the Exchange store.
	MeetingMessage *MeetingMessage `xml:"t:MeetingMessage"`
	// The MeetingRequest element represents a meeting request in the Exchange store.
	MeetingRequest *MeetingRequest `xml:"t:MeetingRequest"`
	// The MeetingResponse element represents a meeting response in the Exchange store.
	MeetingResponse *MeetingResponse `xml:"t:MeetingResponse"`
	// The Message element represents a Microsoft Exchange e-mail message.
	Message *Message `xml:"t:Message"`
	// The Task element represents a task in the Exchange store.
	Task *Task `xml:"t:Task"`
}
