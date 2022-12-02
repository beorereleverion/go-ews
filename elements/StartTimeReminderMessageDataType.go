package elements

// The StartTime (ReminderMessageDataType) element specifies the starting time of the item that the reminder is for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttime-remindermessagedatatype
import "time"

type StartTimeReminderMessageDataType time.Time
