package elements

// The StartWallClock element specifies the start time of a meeting in the time zone of the location in which the meeting takes place.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startwallclock
import "time"

type StartWallClock time.Time
