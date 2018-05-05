// Interface for detecting change in storing
// files & documents in tdoc
package storage

import "github.com/dsalin/tdoc/core"

// Detects all changed files according to the latest index used
type ChangeTracker interface {
	Track(path string, index *core.Index) []string
}
