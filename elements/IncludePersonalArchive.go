package elements

// The IncludePersonalArchive element specifies whether to include the personal archive in the search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/includepersonalarchive
type IncludePersonalArchive bool

const (
	// false
	IncludePersonalArchivefalse IncludePersonalArchive = false
	// true
	IncludePersonalArchivetrue IncludePersonalArchive = true
)
