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
			expStr: " (at :0)",
		},
		{
			ID: testhelper.MkID("no Loc"),
			e: Err{
				Msg: "no Loc",
			},
			expStr: "no Loc (at :0)",
		},
		{
			ID: testhelper.MkID("with Loc"),
			e: Err{
				Msg: "Msg",
				Loc: *loc,
			},
			expStr: "Msg (at test:0)",
		},
	}

	for _, tc := range testCases {
		if s := tc.e.Error(); s != tc.expStr {
			t.Log(tc.IDStr())
			t.Errorf("\t: e.Error() should have returned '%s' not '%s'",
				tc.expStr, s)
		}
	}

}
