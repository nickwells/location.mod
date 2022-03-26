package location_test

import (
	"fmt"
	"testing"

	"github.com/nickwells/location.mod/location"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestHasLoop(t *testing.T) {
	l1 := location.New("f1")
	l2 := location.New("f2")
	l3 := location.New("f3")

	chain := location.LocChain{*l1, *l2, *l3}

	testCases := []struct {
		testhelper.ID
		locName     string
		chain       location.LocChain
		expHasLoop  bool
		expLoopDesc string
	}{
		{
			ID:          testhelper.MkID("loop expected"),
			locName:     "f2",
			chain:       chain,
			expHasLoop:  true,
			expLoopDesc: "f1:0 --> f2:0 ==> f3:0 ==> Back to f2",
		},
		{
			ID:      testhelper.MkID("loop not expected"),
			locName: "notInIncludeChain",
			chain:   chain,
		},
	}

	for _, tc := range testCases {
		testDesc := fmt.Sprintf("HasLoop('%s', %s)",
			tc.locName, tc.chain)
		hasLoop, loopDesc := tc.chain.HasLoop(tc.locName)
		if hasLoop != tc.expHasLoop {
			t.Log(tc.IDStr())
			t.Logf("\t : expected: %v\n", tc.expHasLoop)
			t.Logf("\t :      got: %v\n", hasLoop)
			t.Error("\t: " + testDesc)
		}
		if loopDesc != tc.expLoopDesc {
			t.Log(tc.IDStr())
			t.Logf("\t : expected loop description: %v\n", tc.expLoopDesc)
			t.Logf("\t :                       got: %v\n", loopDesc)
			t.Error("\t: " + testDesc)
		}
	}
}
