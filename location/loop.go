package location

import "strings"

// HasLoop checks to see if the name already appears in the LocChain which
// records the chain of sources getting to this point. If a loop is detected
// then this returns true and a description of the path taken to get to the
// loop, otherwise it returns false and an empty string.
func (chain LocChain) HasLoop(name string) (bool, string) {
	var loopDesc strings.Builder

	var loopDetected bool

	incl := " --> "

	for _, l := range chain {
		if l.name == name {
			loopDetected = true
			incl = " ==> "
		}

		loopDesc.WriteString(l.String())
		loopDesc.WriteString(incl)
	}

	if loopDetected {
		loopDesc.WriteString("Back to ")
		loopDesc.WriteString(name)

		return true, loopDesc.String()
	}

	return false, ""
}
