package location_test

import (
	"testing"

	"github.com/nickwells/location.mod/location"
	"github.com/nickwells/testhelper.mod/testhelper"
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
	checkStr(t, "after SetNote", "[note]: test1:2", l.String())

	l.Incr() // Incr should not affect the notes
	checkStr(t, "after third Incr", "[note]: test1:3", l.String())
}

// checkLocContents confirms that the string is as expected and prints an
// error if not
func checkLocContents(t *testing.T, loc *location.L,
	testID, src, note, content string, hasContent bool, idx int64) {
	t.Helper()

	testhelper.DiffString(t, "loc: "+testID, "source", loc.Source(), src)
	testhelper.DiffInt64(t, "loc: "+testID, "idx", loc.Idx(), idx)
	testhelper.DiffString(t, "loc: "+testID, "note", loc.Note(), note)
	c, hc := loc.Content()
	testhelper.DiffBool(t, "loc: "+testID, "has content", hc, hasContent)
	testhelper.DiffString(t, "loc: "+testID, "content", c, content)
}

func TestLoc(t *testing.T) {
	src := "test1"
	l := location.New(src)
	checkLocContents(t, l, "Initially", src, "", "", false, 0)

	l.Incr()
	checkLocContents(t, l, "1st Incr", src, "", "", false, 1)

	l.Incr()
	checkLocContents(t, l, "2nd Incr", src, "", "", false, 2)

	l.SetNote("note")
	checkLocContents(t, l, "After note set", src, "note", "", false, 2)

	l.SetContent("content")
	checkLocContents(t, l, "After content set", src, "note", "content", true, 2)

	l.Incr()
	checkLocContents(t, l, "3rd Incr", src, "note", "", false, 3)
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
