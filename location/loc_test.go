package location_test

import (
	"testing"

	"github.com/nickwells/location.mod/location"
)

// checkStr confirms that the string is as expected and prints an error if
// not
func checkStr(t *testing.T, name, expectedStr, s string) {
	t.Helper()
	if s != expectedStr {
		t.Log(name)
		t.Log("\t: string-form should be: '" + expectedStr + "'")
		t.Log("\t: string-form       was: '" + s + "'")
		t.Error("\t: unexpected string-form of the location")
	}
}

func TestLocString(t *testing.T) {
	l := location.New("test1")
	checkStr(t, "initially", "test1:0", l.String())

	l.Incr()
	checkStr(t, "after Incr", "test1:1", l.String())

	l.SetContent("content")
	checkStr(t, "after SetContent", "test1:1: content", l.String())

	l.Incr() // Incr should clear the content
	checkStr(t, "after second Incr", "test1:2", l.String())

	l.SetNote("note")
	checkStr(t, "after SetNote", "[ note ]: test1:2", l.String())

	l.Incr() // Incr should not affect the notes
	checkStr(t, "after third Incr", "[ note ]: test1:3", l.String())
}

// checkSrcIdx confirms that the string is as expected and prints an error if
// not
func checkSrcIdx(t *testing.T, name, expSrc, src string, expIdx, idx int64) {
	t.Helper()
	if src != expSrc {
		t.Log(name)
		t.Log("\t: source should be: '" + expSrc + "'")
		t.Log("\t: source       was: '" + src + "'")
		t.Error("\t: unexpected Source() of the location")
	}
	if idx != expIdx {
		t.Log(name)
		t.Logf("\t: idx should be: %d", expIdx)
		t.Logf("\t: idx       was: %d", idx)
		t.Error("\t: unexpected Idx() of the location")
	}
}

func TestLoc(t *testing.T) {
	src := "test1"
	l := location.New(src)
	checkSrcIdx(t, "Initially", src, l.Source(), 0, l.Idx())

	l.Incr()
	checkSrcIdx(t, "After first Incr", src, l.Source(), 1, l.Idx())

	l.Incr()
	checkSrcIdx(t, "After second Incr", src, l.Source(), 2, l.Idx())
}

// checkErr confirms that the string is as expected and prints an error if
// not
func checkErr(t *testing.T, name string, expLoc location.L, expMsg string, err location.Err) {
	t.Helper()
	if err.Msg != expMsg {
		t.Log(name)
		t.Log("\t: message should be: '" + expMsg + "'")
		t.Log("\t: message       was: '" + err.Msg + "'")
		t.Error("\t: unexpected Msg part of the error")
	}
	if err.Loc != expLoc {
		t.Log(name)
		t.Logf("\t: location should be: %s", expLoc)
		t.Logf("\t: location       was: %s", err.Loc)
		t.Error("\t: unexpected Loc part of the error")
	}
}

func TestLocError(t *testing.T) {
	l := location.New("test")
	e1 := l.Error("msg 1")
	checkErr(t, "Error", *l, "msg 1", e1)
	e2 := l.Errorf("msg %d", 2)
	checkErr(t, "Errorf", *l, "msg 2", e2)
}
