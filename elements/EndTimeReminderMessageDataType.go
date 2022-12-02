package elements

// The EndTime element represents the end of the time span to query for reminders.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/endtime-remindermessagedatatype
import "time"

type EndTimeReminderMessageDataType time.Time
