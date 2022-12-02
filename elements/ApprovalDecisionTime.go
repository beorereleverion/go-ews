package elements

// The ApprovalDecisionTime element specifies the time at which the approval decision was made.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/approvaldecisiontime
import "time"

type ApprovalDecisionTime time.Time
