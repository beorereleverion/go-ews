package elements

// The SearchDumpster element specifies whether to search in the Exchange Dumpster.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/searchdumpster
type SearchDumpster string

const (
	// false
	SearchDumpsterfalse SearchDumpster = `false`
	// true
	SearchDumpstertrue SearchDumpster = `true`
)
