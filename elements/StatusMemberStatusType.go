package elements

// The Status element provides information about the status of a distribution list member on the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/status-memberstatustype
type StatusMemberStatusType string

const (
	// Demoted
	StatusMemberStatusTypeDemoted StatusMemberStatusType = `Demoted`
	// Normal
	StatusMemberStatusTypeNormal StatusMemberStatusType = `Normal`
	// Unrecognized
	StatusMemberStatusTypeUnrecognized StatusMemberStatusType = `Unrecognized`
)
