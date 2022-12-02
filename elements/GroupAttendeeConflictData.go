package elements

// The GroupAttendeeConflictData element contains aggregate conflict information about the number of users who are available, the number of users who have conflicts, and the number of users who do not have availability information in a distribution list for a suggested meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groupattendeeconflictdata
type GroupAttendeeConflictData struct {
	// The NumberOfMembers element represents the number of users, resources, and rooms in a distribution list.
	NumberOfMembers *NumberOfMembers `xml:"t:NumberOfMembers"`
	// The NumberOfMembersAvailable element represents the number of distribution list members who are available for a suggested meeting time. This element represents members for whom the status is Free.
	NumberOfMembersAvailable *NumberOfMembersAvailable `xml:"t:NumberOfMembersAvailable"`
	// The NumberOfMembersWithConflict element represents the number of distribution list members who have a conflict with a suggested meeting time. This element represents members who have a status of Busy, OOF, or Tentative.
	NumberOfMembersWithConflict *NumberOfMembersWithConflict `xml:"t:NumberOfMembersWithConflict"`
	// The NumberOfMembersWithNoData element represents the number of distribution list members who do not have published free/busy data to compare to a suggested meeting time. This element represents members of a distribution list that is too large or members who have No Data status.
	NumberOfMembersWithNoData *NumberOfMembersWithNoData `xml:"t:NumberOfMembersWithNoData"`
}
