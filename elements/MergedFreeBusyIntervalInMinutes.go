package elements

// The MergedFreeBusyIntervalInMinutes element represents the time difference between two successive slots in the FreeBusyMerged view.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/mergedfreebusyintervalinminutes
import "encoding/xml"

type MergedFreeBusyIntervalInMinutes struct {
	XMLName xml.Name
	TEXT    int64 `xml:",chardata"`
}

func (M *MergedFreeBusyIntervalInMinutes) SetForMarshal() {
	M.XMLName.Local = "t:MergedFreeBusyIntervalInMinutes"
}

func (M *MergedFreeBusyIntervalInMinutes) GetSchema() *Schema {
	return &SchemaTypes
}
