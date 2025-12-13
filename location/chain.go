package location

import "strings"

// LocChain represents a chain of locations. This might be used to record a
// string of included files.
type LocChain []L

// NewChain returns a new (empty but non-nil) instance of a LocChain
func NewChain() LocChain {
	return make(LocChain, 0)
}

// String provides a string representing a LocChain
func (lc LocChain) String() string {
	var chainDesc strings.Builder

	sep := ""
	for _, l := range lc {
		chainDesc.WriteString(sep)
		chainDesc.WriteString(l.String())

		sep = " --> "
	}

	return chainDesc.String()
}
