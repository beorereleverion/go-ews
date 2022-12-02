package elements

// The Reminder element specifies a reminder for a task or a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/reminder
type Reminder struct {
	// The EndDate element specifies the end date of the item the reminder is for.
	EndDate *EndDateReminderType `xml:"t:EndDate"`
	// The ItemId element contains the unique identifier and change key of an item in the Exchange store.
	ItemId *ItemId `xml:"t:ItemId"`
	// The Location element represents the location of a meeting, appointment, or persona.
	Location *Location `xml:"t:Location"`
	// The RecurringMasterItemId (ItemIdType) element identifies a recurrence master item by identifying the identifiers of one of its related occurrence items.
	RecurringMasterItemId *RecurringMasterItemIdItemIdType `xml:"t:RecurringMasterItemId"`
	// The ReminderGroup element specifies whether the reminder is for a calendar item or a task.
	ReminderGroup *ReminderGroup `xml:"t:ReminderGroup"`
	// The ReminderTime element specifies the time for the reminder to occur.
	ReminderTime *ReminderTime `xml:"t:ReminderTime"`
	// The StartDate element represents the start date of an item.
	StartDate *StartDate `xml:"t:StartDate"`
	// The Subject element represents the subject property of Exchange store items. The subject is limited to 255 characters.
	Subject *Subject `xml:"t:Subject"`
	// The UID element uniquely identifies a calendar item.
	UID *UID `xml:"t:UID"`
}
