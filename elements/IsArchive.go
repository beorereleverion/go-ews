package elements

// The IsArchive element specifies a Boolean value that indicates whether the mailbox is an archive mailbox.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/isarchive
type IsArchive bool

const (
	// false
	IsArchivefalse IsArchive = false
	// true
	IsArchivetrue IsArchive = true
)
