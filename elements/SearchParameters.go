package elements

// The SearchParameters element represents the parameters that define a search folder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchparameters
import "encoding/xml"

type SearchParameters struct {
	XMLName xml.Name

	// The BaseFolderIds element represents the collection of folders that will be mined to determine the contents of a search folder.
	BaseFolderIds *BaseFolderIds `xml:"BaseFolderIds"`
	// The Restriction element represents the restriction or query that is used to filter items or folders in FindItem/FindFolder and search folder operations.
	Restriction *Restriction `xml:"Restriction"`
	// Describes how a search folder traverses the folder hierarchy. The options are for either a Deep or a Shallow search.
	Traversal *string `xml:"Traversal,attr"`
}

func (S *SearchParameters) SetForMarshal() {
	S.XMLName.Local = "t:SearchParameters"
}

func (S *SearchParameters) GetSchema() *Schema {
	return &SchemaTypes
}
