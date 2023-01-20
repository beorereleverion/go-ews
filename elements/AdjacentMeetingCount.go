package elements

// The AdjacentMeetingCount element represents the total number of calendar items that are adjacent to a meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/adjacentmeetingcount
import "encoding/xml"

type AdjacentMeetingCount struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (A *AdjacentMeetingCount) SetForMarshal() {
	A.XMLName.Local = "t:AdjacentMeetingCount"
}

func (A *AdjacentMeetingCount) GetSchema() *Schema {
	return &SchemaTypes
}
