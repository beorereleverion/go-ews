package elements

// The Groups element represents an array of instant messaging (IM) groups.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/groups-arrayofimgrouptype
import "encoding/xml"

type GroupsArrayOfImGroupType struct {
	XMLName xml.Name

	// The ImGroup element represents an instant messaging group.
	ImGroup *ImGroup `xml:"ImGroup"`
}
