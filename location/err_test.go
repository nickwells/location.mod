package location

import (
	"testing"

	"github.com/nickwells/testhelper.mod/testhelper"
)

func TestError(t *testing.T) {
	loc := New("test")
	testCases := []struct {
		testhelper.ID
		e      Err
		expStr string
	}{
		{
			ID:     testhelper.MkID("nil"),
			expStr: "\nAt: :0",
		},
		{
			ID: testhelper.MkID("no Loc"),
			e: Err{
				Msg: "no Loc",
			},
			expStr: "no Loc\nAt: :0",
		},
		{
			ID: testhelper.MkID("with Loc"),
			e: Err{
				Msg: "Msg",
				Loc: *loc,
			},
			expStr: "Msg\nAt: test:0",
		},
	}

	for _, tc := range testCases {
		testhelper.DiffString(t, tc.IDStr(), "Error()", tc.e.Error(), tc.expStr)
	}

}
