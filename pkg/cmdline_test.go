package gocodoc

import (
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {

	if err := fileexists("../VERSION"); err != nil {
		t.Error("Error in CmdLine", err.Error())
	}

}

func TestCommandLine(t *testing.T) {

	//We are one directory under the root directory
	os.Args[1] = "--src=../src"
	os.Args[2] = "--dest=../fakelib"

	src, dest, reb := cmdlineprocess()
	if !reb {
		t.Error("Error in CmdLine")
		return
	}
	if err := fileexists(src); err != nil {
		t.Error("Source Does not exist:", err.Error())
	}
	if err := fileexists(dest); err != nil {
		t.Error("Destination Does not exist:", err.Error())
	}

}
