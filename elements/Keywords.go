package elements

// The Keywords element specifies keywords for a FindMailboxStatisticsByKeywords operation search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/keywords-ex15websvcsotherref
import "encoding/xml"

type Keywords struct {
	XMLName xml.Name

	// The String element represents a string that is used by items, contacts, tasks, and conversations.
	String *String `xml:"String"`
}
